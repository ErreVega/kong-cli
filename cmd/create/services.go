package create

import (
	"kong-cli/config"
	"kong-cli/utils"

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

var create_services = &cobra.Command{
	Use:   "services",
	Short: "list all kong services",
	Long:  `This command retrieves all kong services in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: func(cmd *cobra.Command, args []string) {

		var url string = config.Config.GetUrl() + "/services"
		res := serviceStruct{}
		utils.GetJson(url, &res)

	},
}

func init() {

	Create.AddCommand(create_services)

}
