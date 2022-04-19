package main

import (
	"fmt"
)
import "github.com/ghotfall/detrint/builtin/util"

func main() {
	fmt.Println("Script 2 executed!")
	util.Logger.Info("Test message from script 2")
	fmt.Println(util.Vars)
}
