package main

import (
	"gcreate/conf"
	"gcreate/pongo"
)

func main() {
	//加载配置
	conf.Init()

	//加载模版引擎
	pongo.Exec("")
	//创建数据库连接

	//执行生成
	// tmpl.Exec(tables)
}
