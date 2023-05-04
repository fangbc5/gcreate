package handler

import (
	"{{ProjectName}}/busiflow"
	{% if Metadata %}"{{ProjectName}}/metadata"{% endif %}
	{% if Metavo %}"{{ProjectName}}/metavo"{% endif %}
	"net/http"

	"github.com/fangbc5/gogo/core/common"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	busiflow.BusiFlow
}
