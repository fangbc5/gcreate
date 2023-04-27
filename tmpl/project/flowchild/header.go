package flowchild

import (
	{% if ActionDao %}"gorm.io/gorm"
	{% endif %}
)

type FlowChild struct {
	{% if ActionDao %}DaoApi *gorm.DB
	{% endif %}
}
