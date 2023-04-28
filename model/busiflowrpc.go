package model

type BusiFlowRpc struct {
	Id      int `gorm:"primary_key:true"`
	Name    string
	Code    string
	Comment string
	InputParamsId    int
	InputParamsType  string
	InputParamsCode  string
	OutputParamsId   int
	OutputParamsType string
	OutputParamsCode string
}
