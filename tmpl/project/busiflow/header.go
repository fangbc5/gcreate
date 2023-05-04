package busiflow

import (
	"{{ProjectName}}/busiflowchild"
	{% if Metadata %}"{{ProjectName}}/metadata"{% endif %}
	{% if Metavo %}"{{ProjectName}}/metavo"{% endif %}

	database "github.com/fangbc5/gogo/core/database/mysql"
)
	
type BusiFlow struct {
	*busiflowchild.BusiFlowChild
}
