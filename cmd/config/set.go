package config

import (
	"kong-cli/utils"
	"strings"

	"github.com/spf13/cobra"
)

var ConfigFlags utils.Configuration

var set = &cobra.Command{
	Use:   "set",
	Short: "use to set each value of configuration",
	// Long:    `Retieve a list of objects to interact with kong services`,
	Example: "kong-cli config set KONG_PORT=8001",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		for _, v := range args {
			a := strings.Split(v, "=")
			utils.Config.SetConfigProperty(a[0], a[1])
		}
	},
}

var set2 = &cobra.Command{
	Use:   "set2",
	Short: "use to set each value of configuration",
	// Long:    `Retieve a list of objects to interact with kong services`,
	Example: "kong-cli config set --KONG_PORT=8001",
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flags().Lookup("KONG_HOST").Changed {
			value, err := cmd.Flags().GetString("KONG_HOST")
			utils.PrintErr(err)
			utils.Config.SetKONG_HOST(value)
		}

		if cmd.Flags().Lookup("KONG_PORT").Changed {
			value, err := cmd.Flags().GetInt("KONG_PORT")
			utils.PrintErr(err)
			utils.Config.SetKONG_PORT(value)
		}

	},
}

func init() {

	Config.AddCommand(set)
	Config.AddCommand(set2)
	set2.Flags().StringVar(&ConfigFlags.KONG_HOST, "KONG_HOST", ConfigFlags.KONG_HOST, "--KONG_HOST=myhost")
	set2.Flags().IntVar(&ConfigFlags.KONG_PORT, "KONG_PORT", ConfigFlags.KONG_PORT, "--KONG_PORT=0000")

}
