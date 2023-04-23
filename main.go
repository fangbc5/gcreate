package main

import (
	"gcreate/conf"
	"gcreate/flow"
)

func main() {
	//加载配置
	c := conf.Init()
	flow.Start(c)
	//加载模版引擎
	// pongo.Exec(c.Dir.Tmpl,c.Dir.Out,"")
	//创建数据库连接

	//执行生成
	// tmpl.Exec(tables)
}
