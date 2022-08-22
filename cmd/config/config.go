package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var KONG_HOST_Flag string
var KONG_PORT_Flag int

var Config = &cobra.Command{
	Use:   "config",
	Short: "use to configure the app",
	// Long:    `Retieve a list of objects to interact with kong services`,
	// Example: "kong-cli get services, kong-cli get routes",
	Run: func(cmd *cobra.Command, args []string) {

		err := os.MkdirAll("./config/", os.ModePerm)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		f, err := os.Create("./config/" + args[0])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		f.Write([]byte("holaaaaaa"))
		fmt.Printf("OK\n")

	},
}
