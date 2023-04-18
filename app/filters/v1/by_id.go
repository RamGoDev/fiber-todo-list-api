package filters_v1

import (
	"gorm.io/gorm"
)

func ById(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "id"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where(
			"id = ?",
			value,
		)
	}

	return query
}
