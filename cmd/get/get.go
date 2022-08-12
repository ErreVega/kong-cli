package get

import (
	"github.com/spf13/cobra"
)

var FieldsFlag string

var Get = &cobra.Command{
	Use:     "get",
	Short:   "retrieve objects of following command",
	Long:    `Retieve a list of objects to interact with kong services`,
	Example: "kong-cli get services, kong-cli get routes",
	//	Run: func(cmd *cobra.Command, args []string) {	},+

}
