package filters_v1

import (
	"strings"

	"gorm.io/gorm"
)

func ByDescription(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "description"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where(
			"LOWER(description) LIKE ?",
			"%"+strings.ToLower(value.(string))+"%",
		)
	}

	return query
}
