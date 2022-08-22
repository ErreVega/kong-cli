package get

import (
	"fmt"
	"kong-cli/utils"
	"os"

	"sync"

	"strings"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

type ServiceStruct struct {
	ID                string      `json:"id,omitempty" yaml:"id"`
	Protocol          string      `json:"protocol,omitempty" yaml:"protocol"`
	Path              string      `json:"path,omitempty" yaml:"path"`
	Retries           int         `json:"retries,omitempty" yaml:"retries"`
	Name              string      `json:"name,omitempty" yaml:"name"`
	CaCertificates    interface{} `json:"ca_certificates,omitempty" yaml:"ca_certificates"`
	Port              int         `json:"port,omitempty" yaml:"port"`
	ClientCertificate interface{} `json:"client_certificate,omitempty" yaml:"client_certificate"`
	WriteTimeout      int         `json:"write_timeout,omitempty" yaml:"write_timeout"`
	TLSVerify         interface{} `json:"tls_verify,omitempty" yaml:"tls_verify"`
	UpdatedAt         int         `json:"updated_at,omitempty" yaml:"updated_at"`
	TLSVerifyDepth    interface{} `json:"tls_verify_depth,omitempty" yaml:"tls_verify_depth"`
	Tags              interface{} `json:"tags,omitempty" yaml:"tags"`
	ReadTimeout       int         `json:"read_timeout,omitempty" yaml:"read_timeout"`
	Enabled           bool        `json:"enabled,omitempty" yaml:"enabled"`
	ConnectTimeout    int         `json:"connect_timeout,omitempty" yaml:"connect_timeout"`
	CreatedAt         int         `json:"created_at,omitempty" yaml:"created_at"`
	Host              string      `json:"host,omitempty" yaml:"host"`
	Message           string      `json:"message,omitempty" yaml:"message"` // en caso de error
}

type getServiceStruct struct {
	Data []ServiceStruct `json:"data"`
	Next interface{}     `json:"next"`
}

var get_services = &cobra.Command{
	Use:   "services",
	Short: "list all kong services",
	Long:  `This command retrieves all kong services in a table. You can chosse what fields to get with --all or `,
	//	Example: "kong-cli services",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {

			var url string = utils.Config.GetUrl() + "/services"
			res := getServiceStruct{}
			err := utils.GetJson(url, utils.GET, &res, "")
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				serviceTables(res.Data)
			}

		} else {
			table := make([]ServiceStruct, len(args))
			// arrErr := make([]error, len(args))

			wg.Add(len(args))
			errChanel := make(chan map[string]error, len(args))

			for index, _ := range args {
				table[index] = ServiceStruct{}
				var url string = utils.Config.GetUrl() + "/services/" + args[index]

				go asynReq(url, utils.GET, &table[index], "", errChanel, args[index])
			}

			wg.Wait()
			close(errChanel)
			for v := range errChanel {
				for k, sv := range v {
					fmt.Fprintf(os.Stderr, "Error %v:  %v\n", k, sv)
				}
			}
			serviceTables(table)
		}
	},
}

func asynReq(url string, m utils.Method, i interface{}, jsonBody string, ch chan map[string]error, key string) {
	defer wg.Done()
	err := utils.GetJson(url, m, &i, jsonBody)
	if err != nil {
		errorMap := make(map[string]error)
		errorMap[key] = err
		ch <- errorMap
	}
}

func serviceTables(arrS []ServiceStruct) {
	switch FieldsFlag {
	case "":
		utils.TablePrint(arrS, "ID", "Path", "Name", "Enabled", "Host", "Port")
	case "all":
		utils.TablePrint(arrS, "ID", "Protocol", "Path", "Retries", "Name",
			"CaCertificates", "Port", "ClientCertificate", "WriteTimeout", "TLSVerify", "UpdatedAt",
			"TLSVerifyDepth", "Tags", "ReadTimeout", "Enabled", "ConnectTimeout", "CreatedAt", "Host",
		)
	default:
		fields := strings.Split(FieldsFlag, ",")
		utils.TablePrint(arrS, fields...)
	}
}

func init() {

	Get.AddCommand(get_services)
	get_services.Flags().StringVarP(&FieldsFlag, "fields", "f", "", "all: to return every field, or list selected field separated by comas")

}
