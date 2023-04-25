func (fm FlowMain) {{Code|capfirst}}() {
	{% for child in Childs %}fm.FlowChild.{{child.Code|capfirst}}()
	{% endfor %}
}
