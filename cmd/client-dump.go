package cmd

import (
	"os"
	"os/exec"

	"github.com/checkpoint-restore/go-criu"
	"github.com/checkpoint-restore/go-criu/rpc"
	"github.com/nravic/cedana-client/utils"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var dump_storage_dir string
var pid int

func init() {
	clientCommand.AddCommand(dumpCommand)
	dumpCommand.Flags().StringVarP(&dump_storage_dir, "dumpdir", "d", "", "folder to dump checkpoint into")
	dumpCommand.Flags().IntVarP(&pid, "pid", "p", 0, "pid to dump")
}

// This is a direct dump command. Won't be used in practice, we want to start a daemon
var dumpCommand = &cobra.Command{
	Use:   "dump",
	Short: "Directly dump a process",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := instantiateClient()
		if err != nil {
			return err
		}
		config, err := utils.InitConfig()
		if err != nil {
			c.logger.Fatal().Err(err).Msg("Could not read config")
		}
		// load from config if flags aren't set
		if dump_storage_dir == "" {
			dump_storage_dir = config.Client.DumpStorageDir
		}

		if pid == 0 {
			pid, err = utils.GetPid(config.Client.ProcessName)
			if err != nil {
				c.logger.Err(err).Msg("Could not parse process name from config")
				return err
			}
		}

		err = c.dump(pid, dump_storage_dir)
		if err != nil {
			return err
		}

		defer c.cleanupClient()
		return nil
	},
}

func (c *Client) prepare_dump(pid int, dump_storage_dir string) {
	// copy all open file descriptors for a process
	cmd := exec.Command("ls", "-l", `/proc/${pid}/fd`)
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.logger.Fatal().Err(err).Msgf(`could not ls /proc for pid ${pid}`)
	}
	c.logger.Debug().Bytes(`open fds for pid ${pid}`, out)
	err = os.WriteFile(`${dump_storage_dir}/open_fds`, out, 0644)
}

func (c *Client) prepare_opts() rpc.CriuOpts {
	opts := rpc.CriuOpts{
		LogLevel:       proto.Int32(4),
		LogFile:        proto.String("dump.log"),
		ShellJob:       proto.Bool(false),
		LeaveRunning:   proto.Bool(true),
		TcpEstablished: proto.Bool(true),
		GhostLimit:     proto.Uint32(uint32(10000000)),
		ExtMasters:     proto.Bool(true),
	}
	return opts

}

func (c *Client) dump(pid int, dump_storage_dir string) error {

	// TODO - Dynamic storage (depending on process)
	img, err := os.Open(dump_storage_dir)
	if err != nil {
		c.logger.Fatal().Err(err).Msgf("could not open checkpoint storage dir %s", dump_storage_dir)
		return err
	}
	defer img.Close()

	// ideally we can load and unmarshal this entire struct, from a partial block in the config
	c.prepare_dump(pid, dump_storage_dir)
	opts := c.prepare_opts()
	opts.ImagesDirFd = proto.Int32(int32(img.Fd()))
	opts.Pid = proto.Int32(int32(pid))

	// perform multiple consecutive passes of the dump, altering opts as needed
	// go-CRIU doesn't expose some of this stuff, need to hand-code
	// incrementally add as you test different processes and they fail

	// fmt.Printf("starting dump with opts: %+v\n", opts)

	// do some process checks here and add opts accordingly
	c.logger.Info().Msgf(`beginning dump of pid %d`, pid)
	err = c.CRIU.Dump(opts, criu.NoNotify{})
	if err != nil {
		// TODO - better error handling
		c.logger.Fatal().Err(err).Msg("Error dumping process")
		return err
	}

	// automate later
	cmd := exec.Command("cp", "code-server.log", "code-server.log.bak")
	cmd.Run()

	return nil
}
