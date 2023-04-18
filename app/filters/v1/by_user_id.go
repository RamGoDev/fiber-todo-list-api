package filters_v1

import (
	"gorm.io/gorm"
)

func ByUserId(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "user_id"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where(
			"user_id = ?",
			value,
		)
	}

	return query
}
