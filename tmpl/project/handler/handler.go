func (h Handler) {{Code|capfirst}}(c *gin.Context) {
	{%if InputParamsCode != ""%}{{InputParamsCode|lower}} := &{{InputParamsType}}.{{InputParamsCode}}{}
	c.ShouldBindJSON({{InputParamsCode|lower}}){%endif%}
	{% for busiFlow in BusiFlows %}{% if busiFlow.Code|lower == OutputParamsCode|lower %}{{busiFlow.Code|lower}} := {% endif %}h.FlowMain.{{busiFlow.Code|capfirst}}({% if busiFlow.InputParamsCode != "" %}{{busiFlow.InputParamsCode|lower}}{% endif %})
	{% endfor %}
	c.JSON(http.StatusOK,common.GetRspData({{OutputParamsCode|lower}}))
}
