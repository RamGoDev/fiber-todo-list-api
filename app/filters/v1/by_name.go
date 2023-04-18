package filters_v1

import (
	"strings"

	"gorm.io/gorm"
)

func ByName(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "name"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where(
			"LOWER(name) LIKE ?",
			"%"+strings.ToLower(value.(string))+"%",
		)
	}

	return query
}
