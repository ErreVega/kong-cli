package get

import (
	"kong-cli/config"
	"kong-cli/utils"

	"strings"

	"github.com/spf13/cobra"
)

type pluginsStruct struct {
	Data []struct {
		Route   interface{} `json:"route"`
		ID      string      `json:"id"`
		Service struct {
			ID string `json:"id"`
		} `json:"service"`
		Protocols []string    `json:"protocols"`
		Enabled   bool        `json:"enabled"`
		Consumer  interface{} `json:"consumer"`
		CreatedAt int         `json:"created_at"`
		Config    interface{} `json:"config,omitempty"`
		Tags      interface{} `json:"tags"`
		Name      string      `json:"name"`
	} `json:"data"`
	Next interface{} `json:"next"`
}

var get_plugins = &cobra.Command{
	Use:   "plugins",
	Short: "list all kong plugins",
	Long:  `This command retrieves all kong plugins in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: pluginFunc,
}

func pluginFunc(cmd *cobra.Command, args []string) {

	var url string = config.Config.GetUrl() + "/plugins"
	res := pluginsStruct{}
	utils.GetJson(url, &res)

	switch FieldsFlag {
	case "":
		utils.TablePrint(res.Data, "Route", "ID", "Name", "Service", "Enabled")
	case "all":
		utils.TablePrint(res.Data, "Route", "ID", "Service", "Protocols", "Enabled",
			"Consumer", "CreatedAt", "Config", "Tags", "Name",
		)

	default:
		fields := strings.Split(FieldsFlag, ",")
		utils.TablePrint(res.Data, fields...)
	}

}

func init() {

	Get.AddCommand(get_plugins)
	get_plugins.Flags().StringVarP(&FieldsFlag, "fields", "f", "", "all: to return every field, or list selected field separated by comas")

}
