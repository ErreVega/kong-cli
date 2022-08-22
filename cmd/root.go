package cmd

import (
	"fmt"
	"kong-cli/cmd/config"
	"kong-cli/cmd/create"
	"kong-cli/cmd/delete"
	"kong-cli/cmd/get"

	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kong-cli",
	Short: "kong-cli is a tool to intercat with the kong admin API",
	Long: `kong-cli is a tool 
			to intercat with the 
			kong admin API`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// 	print("kong-cli is ok")
	// },
}

func Execute() {
	RootCmd.AddCommand(sum)
	RootCmd.AddCommand(get.Get)
	RootCmd.AddCommand(create.Create)
	RootCmd.AddCommand(delete.Delete)
	RootCmd.AddCommand(config.Config)

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
