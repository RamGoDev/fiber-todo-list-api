package queries_v1

import (
	"gorm.io/gorm"
)

func ByName(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	query = query.Where(
		"LOWER(name) LIKE ?",
		"%"+value+"%",
	)

	return query
}
