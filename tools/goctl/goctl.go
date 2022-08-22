package main

import (
	"github.com/Holyson/test-go-zero-cors/core/load"
	"github.com/Holyson/test-go-zero-cors/core/logx"
	"github.com/Holyson/test-go-zero-cors/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
