package cmd

// This file contains all the config-related commands when starting `cedana config ...`

import (
	"encoding/json"
	"fmt"

	"github.com/cedana/cedana/utils"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show currently set configuration",
	RunE: func(_ *cobra.Command, _ []string) error {
		config, err := utils.GetConfig()
		if err != nil {
			logger.Error().Err(err).Msg("failed to get config")
			return err
		}

		prettycfg, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			logger.Error().Err(err).Msg("failed to parse json")
			return err
		}
		fmt.Printf("config: %v\n", string(prettycfg))

		return nil
	},
}

func init() {
	logger = utils.GetLogger()
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(showCmd)
}
