package get

import (
	"kong-cli/config"
	"kong-cli/utils"

	"strings"

	"github.com/spf13/cobra"
)

type serviceStruct struct {
	Data []struct {
		ID                string      `json:"id"`
		Protocol          string      `json:"protocol"`
		Path              string      `json:"path"`
		Retries           int         `json:"retries"`
		Name              string      `json:"name"`
		CaCertificates    interface{} `json:"ca_certificates"`
		Port              int         `json:"port"`
		ClientCertificate interface{} `json:"client_certificate"`
		WriteTimeout      int         `json:"write_timeout"`
		TLSVerify         interface{} `json:"tls_verify"`
		UpdatedAt         int         `json:"updated_at"`
		TLSVerifyDepth    interface{} `json:"tls_verify_depth"`
		Tags              interface{} `json:"tags"`
		ReadTimeout       int         `json:"read_timeout"`
		Enabled           bool        `json:"enabled"`
		ConnectTimeout    int         `json:"connect_timeout"`
		CreatedAt         int         `json:"created_at"`
		Host              string      `json:"host"`
	} `json:"data"`
	Next interface{} `json:"next"`
}

var get_services = &cobra.Command{
	Use:   "services",
	Short: "list all kong services",
	Long:  `This command retrieves all kong services in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: func(cmd *cobra.Command, args []string) {

		var url string = config.Config.GetUrl() + "/services"
		res := serviceStruct{}
		utils.GetJson(url, &res)

		switch FieldsFlag {
		case "":
			utils.TablePrint(res.Data, "ID", "Path", "Name", "Enabled", "Host", "Port")
		case "all":
			utils.TablePrint(res.Data, "ID", "Protocol", "Path", "Retries", "Name",
				"CaCertificates", "Port", "ClientCertificate", "WriteTimeout", "TLSVerify", "UpdatedAt",
				"TLSVerifyDepth", "Tags", "ReadTimeout", "Enabled", "ConnectTimeout", "CreatedAt", "Host",
			)

		default:
			fields := strings.Split(FieldsFlag, ",")
			utils.TablePrint(res.Data, fields...)
		}

	},
}

func init() {

	Get.AddCommand(get_services)
	get_services.Flags().StringVarP(&FieldsFlag, "fields", "f", "", "all: to return every field, or list selected field separated by comas")

}
