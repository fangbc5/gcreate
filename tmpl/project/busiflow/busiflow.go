func (bf BusiFlow) {{Code|capfirst}}() map[string]interface{} {
	d := database.GetGormApi()
	tx := d.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	bf.BusiFlowChild = &busiflowchild.BusiFlowChild{DaoApi: tx}
	{% for child in Childs %}{% if child.Code|lower == OutputParamsCode|lower %}{{child.Code|lower}} := {% endif %}bf.BusiFlowChild.{{child.Code|capfirst}}({% if child.InputParamsCode != "" %}{{child.InputParamsCode|lower}} {{child.InputParamsCode}}{% endif %})
	{% endfor %}
	tx.Commit()
	return {{OutputParamsCode|lower}}
}
