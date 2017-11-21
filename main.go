package main

import (
	"fmt"
	"os"

	"github.com/rai-project/dlframework/framework/cmd/server"
)

func main() {
	if err := server.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
