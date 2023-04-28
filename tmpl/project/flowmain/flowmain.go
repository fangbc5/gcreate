func (fm FlowMain) {{Code|capfirst}}() map[string]interface{} {
	{% for child in Childs %}{% if child.Code|lower == OutputParamsCode|lower %}{{child.Code|lower}} := {% endif %}fm.FlowChild.{{child.Code|capfirst}}({% if child.InputParamsCode != "" %}{{child.InputParamsCode|lower}} {{child.InputParamsCode}}{% endif %})
	{% endfor %}
	return {{OutputParamsCode|lower}}
}
