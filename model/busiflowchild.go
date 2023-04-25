package model

type BusiFlowChild struct {
	Id        int `gorm:"primary_key:true"`
	Name      string
	Code      string
	Comment   string
	Sort      int
	Action    string
	ActionDao []*BusiFlowDao `gorm:"many2many:busi_flow_child_busi_flow_dao;"`
	ActionDco []*BusiFlowDco `gorm:"many2many:busi_flow_child_busi_flow_dco;"`
	ActionRpc []*BusiFlowRpc `gorm:"many2many:busi_flow_child_busi_flow_rpc;"`
	ActionMsg []*BusiFlowMsg `gorm:"many2many:busi_flow_child_busi_flow_msg;"`
}
