package db

import (
	"database/sql"
	"sync"
)

type Datasource struct {
	*sql.DB
	lock     sync.Mutex
	username string
	password string
	driver   string
	url      string
	schema   string
}

type Table struct {
	ProjectName string  //项目名称
	ModuleName  string  //模块名称
	HasTime     bool    //是否存在时间字段
	TableName   string  //表名
	ModelName   string  //模型名
	Fields      []Field //字段列表
}

type Field struct {
	Primary   string
	FieldName string
	FieldDesc string
	DataType  string
	IsNull    string
}
