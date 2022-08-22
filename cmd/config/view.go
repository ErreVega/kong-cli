package config

import (
	"fmt"
	"kong-cli/utils"

	"github.com/spf13/cobra"
)

var view = &cobra.Command{
	Use:   "view",
	Short: "use to retrieve the current configuration",
	// Long:    `Retieve a list of objects to interact with kong services`,
	// Example: "kong-cli get services, kong-cli get routes",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("%+v", utils.Config.GetConfig())
	},
}

func init() {

	Config.AddCommand(view)

}
