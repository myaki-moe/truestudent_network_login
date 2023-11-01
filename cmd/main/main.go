package main

import (
	"os"
	"truestudent_network_login/pkg/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(-1)
	}
}
