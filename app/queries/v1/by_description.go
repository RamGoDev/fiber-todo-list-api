package queries_v1

import (
	"gorm.io/gorm"
)

func ByDescription(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	query = query.Where(
		"LOWER(description) LIKE ?",
		"%"+value+"%",
	)

	return query
}
