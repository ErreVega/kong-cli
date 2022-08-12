package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var sum = &cobra.Command{
	Use:     "sum",
	Short:   "sum n number",
	Long:    `sum n numbers`,
	Example: "kong-cli sum 1.5 2.5 3",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		var result = 0.0

		for _, element := range args {
			num, err := strconv.ParseFloat(element, 32)
			if err == nil {
				result += num
			} else {
				fmt.Fprintln(os.Stderr, "Could not parse ", element, " to float.")
			}
		}
		print(result)
	},
}
