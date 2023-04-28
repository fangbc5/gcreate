package handler

import (
	"{{ProjectName}}/flowmain"
	{% if metadata %}"{{ProjectName}}/metadata"{% endif %}
	{% if metavo %}"{{ProjectName}}/metavo"{% endif %}
	"net/http"

	"github.com/fangbc5/gogo/core/common"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	flowmain.FlowMain
}
