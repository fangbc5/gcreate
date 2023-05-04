func (bfc BusiFlowChild) {{Code|capfirst}}() map[string]interface{} {
	result := make(map[string]interface{}, 16)
	{% if Action == "dao" %}
	{% for dao in ActionDao %}bfc.DaoApi.Raw("{{ dao.Sql }}"{% if dao.SqlParams != "" %}, {{ dao.SqlParams }}{% endif %}).Scan(result)
	{% endfor %}
	{% endif %}
	return result
}
