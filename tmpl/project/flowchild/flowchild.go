func (fc FlowChild) {{Code|capfirst}}() map[string]interface{} {
	{% if Action == "dao" %}fc.DaoApi = database.GetGormApi(){% endif %}
	result := make(map[string]interface{}, 16)
	{% if Action == "dao" %}
	{% for dao in ActionDao %}fc.DaoApi.Raw("{{ dao.Sql }}"{% if dao.SqlParams != "" %}, {{ dao.SqlParams }}{% endif %}).Scan(result)
	{% endfor %}
	{% endif %}
	return result
}
