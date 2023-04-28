package handler

import (
	"{{ProjectName}}/flowmain"
	{% if InputParamsType == "metadata" || OutputParamsType == "metadata" %}"{{ProjectName}}/metadata"{% endif %}
	{% if InputParamsType == "metavo" || OutputParamsType == "metavo" %}"{{ProjectName}}/metavo"{% endif %}
	"net/http"

	"github.com/fangbc5/gogo/core/common"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	flowmain.FlowMain
}
