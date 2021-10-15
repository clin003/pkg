package main

import (
	"fmt"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/pkg/config"

	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	defer func() {
		fmt.Scanln()
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Debugf("run time panic:%v\n", err)
		}
	}()
	// cfg := pflag.StringP("config", "c", "", "config file path")
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Info("config 初始化完成")
}
