package cmd

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"

	"github.com/AbderrahimBouhdida/AzH/config"
	"github.com/AbderrahimBouhdida/AzH/constants"
)

var (
	rootCmd = &cobra.Command{
		Use:     constants.Name,
		Long:    constants.Description,
		Version: constants.Version,
	}
	log logr.Logger
)

func init() {
	config.Init(rootCmd, config.GlobalConfig)
}

func Execute() error {
	return rootCmd.Execute()
}
