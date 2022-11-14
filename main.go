package main

import (
	"gcreate/conf"
	"gcreate/db"
	"gcreate/interface/configloader"
	"gcreate/interface/tmpl"
	"gcreate/templates"
)

func main() {
	//加载配置
	loadConfig()
	//加载模版引擎
	tmpl := getTempEngine()
	//创建数据库连接
	db := getDbEngine()
	defer db.Close()
	tables := db.GetTables(conf.MakeConfig().Names...)
	//执行生成
	tmpl.Exec(tables)
}

func loadConfig() conf.Configuration {
	var c configloader.Config
	c = conf.MakeConfig()
	return *c.Load()
}

func getTempEngine() tmpl.Tmpl {
	var t tmpl.Tmpl
	t = templates.MakePongo()
	return t
}

func getDbEngine() db.Datasource {
	d := db.MakeMysql() //创建mysql数据库对象
	d.Conn()            //连接mysql数据库
	return *d
}
