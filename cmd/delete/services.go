package delete

import (
	"fmt"
	"kong-cli/cmd/get"
	"kong-cli/utils"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var delete_services = &cobra.Command{
	Use:   "services",
	Short: "list all kong services",
	Long: `This command retrieves all kong services in a table. You can chosse what fields to get with --all
	

name				optional	The Service name.
retries				optional	The number of retries to execute upon failure to proxy. Default: 5.
protocol			The protocol used to communicate with the upstream. Accepted values are: "grpc", "grpcs", "http", "https", "tcp", "tls", "tls_passthrough", "udp". Default: "http".
host				The host of the upstream server. Note that the host value is case sensitive.
port				The upstream server port. Default: 80.
path				optional	The path to be used in requests to the upstream server.
connect_timeout		optional	The timeout in milliseconds for establishing a connection to the upstream server. Default: 60000.
write_timeout		optional	The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server. Default: 60000.
read_timeout		optional	The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server. Default: 60000.
tags				optional	An optional set of strings associated with the Service for grouping and filtering.
client_certificate	optional	Certificate to be used as client certificate while TLS handshaking to the upstream server. With form-encoded, the notation is client_certificate.id=<client_certificate id>. With JSON, use “"client_certificate":{"id":"<client_certificate id>"}.
tls_verify			optional	Whether to enable verification of upstream server TLS certificate. If set to null, then the Nginx default is respected.
tls_verify_depth	optional	Maximum depth of chain while verifying Upstream server’s TLS certificate. If set to null, then the Nginx default is respected. Default: null.
ca_certificates		optional	Array of CA Certificate object UUIDs that are used to build the trust store while verifying upstream server’s TLS certificate. If set to null when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted). With form-encoded, the notation is ca_certificates[]=4e3ad2e4-0bc4-4638-8e34-c84a417ba39b&ca_certificates[]=51e77dc2-8f3e-4afa-9d0e-0e3bbbcfd515. With JSON, use an Array.
					enabled	Whether the Service is active. If set to false, the proxy behavior will be as if any routes attached to it do not exist (404). Default: true. Default: true.
url					shorthand-attribute	Shorthand attribute to set protocol, host, port and path at once. This attribute is write-only (the Admin API never returns the URL).
	
	
	
	`,
	Args: cobra.MinimumNArgs(1),
	//	Example: "kong-cli services",
	Run: func(cmd *cobra.Command, args []string) {

		res := get.ServiceStruct{}

		//read file
		dat, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		//parse yaml to struc
		err = yaml.Unmarshal(dat, &res)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(res.ID)
		fmt.Println(res.Name)

		var url string = utils.Config.GetUrl() + "/services/" + res.Name

		//HTTP Call
		err = utils.GetJson(url, utils.DELETE, &res, "")

		table := []get.ServiceStruct{res}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not delete")
			fmt.Fprintln(os.Stderr, err)
		} else {
		}
		utils.TablePrint(table, "ID", "Path", "Name", "Enabled", "Host", "Port")

	},
}

func init() {

	Delete.AddCommand(delete_services)

}
