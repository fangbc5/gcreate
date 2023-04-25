package model

type Interface struct {
	Id             int `gorm:"primary_key:true"`
	Name           string
	Code           string
	Comment        string
	InputParamsId  int
	OutputParamsId int
	Url            string
	Method         string
	BusiFlows      []*BusiFlow `gorm:"many2many:interface_busi_flow;"`
}