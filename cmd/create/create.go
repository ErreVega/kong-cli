package create

import (
	"github.com/spf13/cobra"
)

var Create = &cobra.Command{
	Use:     "create",
	Short:   "create object of following command on kong",
	Long:    `create an object on kong`,
	Example: "kong-cli create services, kong-cli create routes",
	//	Run: func(cmd *cobra.Command, args []string) {	},+

}
