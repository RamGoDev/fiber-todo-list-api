package filters_v1

import (
	"gorm.io/gorm"
)

func ByEmail(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "email"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where("email = ?", value)
	}

	return query
}
