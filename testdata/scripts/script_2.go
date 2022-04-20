package main

import (
	"fmt"
	"github.com/ghotfall/detrint/builtin/shell"
	"strconv"
)
import "github.com/ghotfall/detrint/builtin/util"

func main() {
	fmt.Println("Script 2 executed!")
	util.Logger.Info("Test message from script 2")
	fmt.Println(util.Vars)
	stdout, stderr, code, err := shell.Execute("echol test")
	if err != nil {
		util.Logger.Debug(err.Error())
	}
	util.Logger.Debug(stdout)
	util.Logger.Debug(stderr)
	util.Logger.Debug(strconv.Itoa(code))
}
