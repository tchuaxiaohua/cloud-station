package main

import (
	"fmt"

	"github.com/tchuaxiaohua/cloud-station/cmd"
)

func main() {
	if err := cmd.StartCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
