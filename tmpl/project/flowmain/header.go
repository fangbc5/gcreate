package flowmain

import (
	"{{ProjectName}}/flowchild"
	{% if InputParamsType == "metadata" || OutputParamsType == "metadata" %}"{{ProjectName}}/metadata"{% endif %}
	{% if InputParamsType == "metavo" || OutputParamsType == "metavo" %}"{{ProjectName}}/metavo"{% endif %}
)
	
type FlowMain struct {
	flowchild.FlowChild
}
