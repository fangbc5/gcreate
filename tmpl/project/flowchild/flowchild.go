func (fc FlowChild) {{Code|capfirst}}() {
	{% if Action == "dao" %}
	{% for dao in ActionDao %}fc.DaoApi.Exec("{{ dao.Sql }}"{% if dao.SqlParams != "" %}, {{ dao.SqlParams }}{% endif %})
	{% endfor %}
	{% endif %}
}
