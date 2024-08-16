package cmd

// This file contains all the daemon-related commands when starting `cedana daemon ...`

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/cedana/cedana/pkg/api"
	"github.com/cedana/cedana/pkg/api/services"
	"github.com/cedana/cedana/pkg/api/services/task"
	"github.com/cedana/cedana/pkg/jobservice"
	"github.com/cedana/cedana/pkg/utils"
	"github.com/maragudk/goqite"
	"github.com/maragudk/goqite/jobs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start daemon for cedana client. Must be run as root, needed for all other cedana functionality.",
}

var cudaVersions = map[string]string{
	"11.8": "cuda11_8",
	"12.1": "cuda12_1",
	"12.2": "cuda12_2",
	"12.4": "cuda12_4",
}

var startDaemonCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the rpc server. To run as a daemon, use the provided script (systemd) or use systemd/sysv/upstart.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		logger := ctx.Value("logger").(*zerolog.Logger)

		if os.Getuid() != 0 {
			return fmt.Errorf("daemon must be run as root")
		}

		_, err := utils.InitOtel(cmd.Context(), cmd.Parent().Version)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to initialize otel")
			return err
		}
		logger.Info().Msg("otel initialized")

		if viper.GetBool("profiling_enabled") {
			go startProfiler()
		}
		gpuEnabled, _ := cmd.Flags().GetBool(gpuEnabledFlag)
		// defaults to 11_8, this continues if --cuda is not specified
		cudaVersion, _ := cmd.Flags().GetString(cudaVersionFlag)
		if _, ok := cudaVersions[cudaVersion]; !ok {
			err = fmt.Errorf("invalid cuda version %s, must be one of %v", cudaVersion, cudaVersions)
			logger.Error().Err(err).Msg("invalid cuda version")
			return err
		}

		logger.Info().Msgf("starting daemon version %s", rootCmd.Version)

		grpcPort, _ := cmd.Flags().GetUint64(gprcPortFlag)

		err = api.StartServer(ctx, &api.ServeOpts{
			GPUEnabled:  gpuEnabled,
			CUDAVersion: cudaVersions[cudaVersion],
			GrpcPort:    grpcPort,
		})
		if err != nil {
			logger.Error().Err(err).Msgf("failed to start grpc service, stopping daemon")
			return err
		}

		// sqlite queue
		db, err := sql.Open("sqlite3", ":memory:?_journal=WAL&_timeout=5000&_fk=true")
		if err != nil {
			logger.Error().Err(err).Msgf("failed to open sqlite db")
		}
		db.SetMaxOpenConns(1)
		db.SetMaxIdleConns(1)

		if err := goqite.Setup(context.Background(), db); err != nil {
			logger.Error().Err(err).Msgf("failed to setup context")
		}
		q := goqite.New(goqite.NewOpts{DB: db, Name: "jobs"})
		r := jobs.NewRunner(jobs.NewRunnerOpts{
			Limit:        1,
			PollInterval: 10 * time.Millisecond,
			Queue:        q,
		})

		cts, err := services.NewClient()
		if err != nil {
			logger.Error().Err(err).Msgf("failed to create grpc services client")
			return err
		}

		jobservice.RegisterJobs(r, cts)

		// start job service
		jobServicePort, _ := cmd.Flags().GetUint64(jobServicePortFlag)
		echoServer, err := jobservice.StartService(q, jobServicePort)
		if err != nil {
			logger.Error().Err(err).Msgf("failed to start job service, stopping daemon")
			return err
		}
		echoServer.Start(fmt.Sprintf("localhost:%d", jobServicePort))

		// handle signal cancellation
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
		defer cancel()

		// start the runner
		r.Start(ctx)

		return nil
	},
}

var checkDaemonCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if daemon is running and healthy",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		logger := ctx.Value("logger").(*zerolog.Logger)

		logger.Debug().Msg("Running daemon status check")

		cts, err := services.NewClient()
		if err != nil {
			logger.Error().Err(err).Msg("error creating client")
			return err
		}
		defer cts.Close()

		// regular health check
		healthy, err := cts.HealthCheck(cmd.Context())
		if err != nil {
			logger.Error().Err(err).Msg("health check failed")
			return err
		}
		logger.Info().Msgf("health check returned: %v", healthy)

		// Detailed health check. Need to grab uid and gid to start
		// controller properly and with the right perms.
		var uid int32
		var gid int32
		var groups []int32 = []int32{}

		uid = int32(os.Getuid())
		gid = int32(os.Getgid())
		groups_int, err := os.Getgroups()
		if err != nil {
			logger.Error().Err(err).Msg("error getting user groups")
			return err
		}
		for _, g := range groups_int {
			groups = append(groups, int32(g))
		}

		req := &task.DetailedHealthCheckRequest{
			UID:    uid,
			GID:    gid,
			Groups: groups,
		}

		resp, err := cts.DetailedHealthCheck(cmd.Context(), req)
		if err != nil {
			logger.Error().Err(err).Msg("health check failed")
			return err
		}

		logger.Info().Msgf("health check output: %v", resp)

		return nil
	},
}

// Used for debugging and profiling only!
func startProfiler() {
	utils.StartPprofServer()
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.AddCommand(startDaemonCmd)
	daemonCmd.AddCommand(checkDaemonCmd)
	startDaemonCmd.Flags().BoolP(gpuEnabledFlag, "g", false, "start daemon with GPU support")
	startDaemonCmd.Flags().String(cudaVersionFlag, "11.8", "cuda version to use")
	startDaemonCmd.Flags().Uint64P(gprcPortFlag, "p", 8080, "port for grpc server")
	startDaemonCmd.Flags().Uint64P(jobServicePortFlag, "j", 1444, "port for job service")
}
