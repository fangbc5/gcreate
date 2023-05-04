package busiflowchild

import (
	{% if Metadata %}"{{ProjectName}}/metadata"{% endif %}
	{% if Metavo %}"{{ProjectName}}/metavo"{% endif %}
	
	"gorm.io/gorm"
)

type BusiFlowChild struct {
	DaoApi *gorm.DB
}
