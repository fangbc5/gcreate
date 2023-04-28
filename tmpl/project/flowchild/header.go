package flowchild

import (
	{% if InputParamsType == "metadata" || OutputParamsType == "metadata" %}"{{ProjectName}}/metadata"{% endif %}
	{% if InputParamsType == "metavo" || OutputParamsType == "metavo" %}"{{ProjectName}}/metavo"{% endif %}
	
	{% if ActionDao %}"gorm.io/gorm"
	database "github.com/fangbc5/gogo/core/database/mysql"
	{% endif %}
)

type FlowChild struct {
	{% if ActionDao %}DaoApi *gorm.DB
	{% endif %}
}
