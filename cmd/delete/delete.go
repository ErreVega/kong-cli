package delete

import (
	"github.com/spf13/cobra"
)

var Delete = &cobra.Command{
	Use:     "delete",
	Short:   "delete object of following command on kong",
	Long:    `delete an object on kong`,
	Example: "kong-cli delete services, kong-cli delete routes",
	//	Run: func(cmd *cobra.Command, args []string) {	},+

}
