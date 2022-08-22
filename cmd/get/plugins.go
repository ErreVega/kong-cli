package get

import (
	"fmt"
	"kong-cli/utils"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

type pluginsStruct struct {
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
}
type getPluginsStruct struct {
	Data []pluginsStruct `json:"data"`
	Next interface{}     `json:"next"`
}

var get_plugins = &cobra.Command{
	Use:   "plugins",
	Short: "list all kong plugins",
	Long:  `This command retrieves all kong plugins in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: get_plugins_func,
}

var get_services_plugins = &cobra.Command{
	Use:   "plugins",
	Short: "list all plugins asociated with a service",
	Long:  `This command retrieves all plugins of a service in a table. You can chosse what fields to get with --all or `,
	Args:  cobra.RangeArgs(1, 1),
	//	Example: "kong-cli services",
	Run: get_services_plugins_func,
}

func get_plugins_func(cmd *cobra.Command, args []string) {

	var url string = utils.Config.GetUrl() + "/plugins"
	res := getPluginsStruct{}
	err := utils.GetJson(url, utils.GET, &res, "")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		pluginTables(res.Data)
	}

}

func get_services_plugins_func(cmd *cobra.Command, args []string) {
	var url string = utils.Config.GetUrl() + "/services/" + args[0] + "/plugins"

	res := getPluginsStruct{}
	err := utils.GetJson(url, utils.GET, &res, "")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		pluginTables(res.Data)
	}

}

func pluginTables(pg []pluginsStruct) {
	switch FieldsFlag {
	case "":
		utils.TablePrint(pg, "Route", "ID", "Name", "Service", "Enabled")
	case "all":
		utils.TablePrint(pg, "Route", "ID", "Service", "Protocols", "Enabled",
			"Consumer", "CreatedAt", "Config", "Tags", "Name",
		)

	default:
		fields := strings.Split(FieldsFlag, ",")
		utils.TablePrint(pg, fields...)
	}

}

func init() {

	Get.AddCommand(get_plugins)
	get_services.AddCommand(get_services_plugins)
	get_plugins.Flags().StringVarP(&FieldsFlag, "fields", "f", "", "all: to return every field, or list selected field separated by comas")

}
