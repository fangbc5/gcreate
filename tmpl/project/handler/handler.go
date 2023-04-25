func (h Handler) {{Code|capfirst}}(c *gin.Context) {
	{% for busiFlow in BusiFlows %}h.FlowMain.{{busiFlow.Code|capfirst}}()
	{% endfor %}
}
