package model

type BusiFlowDco struct {
	Id      int `gorm:"primary_key:true"`
	Name    string
	Code    string
	Comment string
}