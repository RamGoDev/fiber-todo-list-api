package filters_v1

import (
	"strings"

	"gorm.io/gorm"
)

func ByTitle(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "title"

	if value, exists := filter[key]; exists && value != "" {
		query = query.Where(
			"LOWER(title) LIKE ?",
			"%"+strings.ToLower(value.(string))+"%",
		)
	}

	return query
}
