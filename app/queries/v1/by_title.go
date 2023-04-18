package queries_v1

import (
	"gorm.io/gorm"
)

func ByTitle(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	query = query.Where(
		"LOWER(title) LIKE ?",
		"%"+value+"%",
	)

	return query
}
