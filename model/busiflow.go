package model

type BusiFlow struct {
	Id      int	`gorm:"primary_key:true"`
	Name    string
	Code    string
	Comment string
	Childs  []*BusiFlowChild `gorm:"many2many:busi_flow_busi_flow_child;"`
}
