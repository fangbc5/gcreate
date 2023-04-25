package model

type BusiFlowMsg struct {
	Id      int `gorm:"primary_key:true"`
	Name    string
	Code    string
	Comment string
}
