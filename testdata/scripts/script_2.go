package main

import (
	"fmt"
	"github.com/ghotfall/detrint/builtin/file"
)
import "github.com/ghotfall/detrint/builtin/util"

func main() {
	fmt.Println("Script 2 executed!")
	util.Logger.Info("Test message from script 2")
	fmt.Println(util.Vars)
	var info file.Info
	info, err := file.GetStat("C:/Users/Ghotfall/Desktop/test")
	if err != nil {
		util.Logger.Debug(err.Error())
	} else {
		fmt.Println(info)
	}
}
