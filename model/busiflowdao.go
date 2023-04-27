package model

type BusiFlowDao struct {
	Id      int `gorm:"primary_key:true"`
	Name    string
	Code    string
	Comment string
	Sql       string
	SqlParams string
}
