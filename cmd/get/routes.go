package get

import (
	"kong-cli/config"
	"kong-cli/utils"

	"strings"

	"github.com/spf13/cobra"
)

type routeStruct struct {
	Data []struct {
		ID                string      `json:"id"`
		Snis              interface{} `json:"snis"`
		ResponseBuffering bool        `json:"response_buffering"`
		UpdatedAt         int         `json:"updated_at"`
		Name              string      `json:"name"`
		PreserveHost      bool        `json:"preserve_host"`
		Paths             []string    `json:"paths"`
		Methods           interface{} `json:"methods"`
		Sources           interface{} `json:"sources"`
		Destinations      interface{} `json:"destinations"`
		Service           struct {
			ID string `json:"id"`
		} `json:"service"`
		PathHandling            string      `json:"path_handling"`
		StripPath               bool        `json:"strip_path"`
		Protocols               []string    `json:"protocols"`
		CreatedAt               int         `json:"created_at"`
		Headers                 interface{} `json:"headers"`
		RegexPriority           int         `json:"regex_priority"`
		Hosts                   interface{} `json:"hosts"`
		HTTPSRedirectStatusCode int         `json:"https_redirect_status_code"`
		Tags                    interface{} `json:"tags"`
		RequestBuffering        bool        `json:"request_buffering"`
	} `json:"data"`
	Next interface{} `json:"next"`
}

var get_routes = &cobra.Command{
	Use:   "routes",
	Short: "list all kong routes",
	Long:  `This command retrieves all kong routes in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: func(cmd *cobra.Command, args []string) {

		var url string = config.Config.GetUrl() + "/routes"
		res := routeStruct{}
		utils.GetJson(url, &res)

		switch FieldsFlag {
		case "":
			utils.TablePrint(res.Data, "ID", "Name", "Service", "Hosts", "Protocols", "Methods")
		case "all":
			utils.TablePrint(res.Data, "ID", "Snis", "ResponseBuffering", "UpdatedAt", "Name",
				"PreserveHost", "Paths", "Methods", "Sources", "Destinations", "Service",
				"PathHandling", "StripPath", "Protocols", "CreatedAt", "Headers", "RegexPriority",
				"Hosts", "HTTPSRedirectStatusCode", "Tags", "RequestBuffering",
			)
		default:
			fields := strings.Split(FieldsFlag, ",")
			utils.TablePrint(res.Data, fields...)
		}

	},
}

func init() {

	Get.AddCommand(get_routes)
	get_routes.Flags().StringVarP(&FieldsFlag, "fields", "f", "", "all: to return every field, or list selected field separated by comas")

}
