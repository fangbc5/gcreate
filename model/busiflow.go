package model

type BusiFlow struct {
	Id               int `gorm:"primary_key:true"`
	Name             string
	Code             string
	Comment          string
	ActionDao        bool
	ActionDco        bool
	ActionRpc        bool
	ActionMsg        bool
	InputParamsId    int
	InputParamsType  string
	InputParamsCode  string
	OutputParamsId   int
	OutputParamsType string
	OutputParamsCode string
	Childs           []*BusiFlowChild `gorm:"many2many:busi_flow_busi_flow_child;"`
}
