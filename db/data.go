package db

type Data struct {
	Project
	Table
}

type Project struct {
	ProjectName string //项目名称
	PackName    string //包名
	ModuleName  string //模块名称
	HasTime     bool   //是否存在时间字段
}
