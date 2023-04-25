package main

import (
	"gcreate/conf"
	"gcreate/flow"
)

func main() {
	//加载配置
	c := conf.Init()
	flow.Start(c)
}
