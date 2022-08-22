package utils

import (
	"fmt"
	"os"
)

func PrintErr(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}
