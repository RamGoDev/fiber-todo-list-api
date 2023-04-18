package queries_v1

import (
	"gorm.io/gorm"
)

func ByEmail(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	query = query.Where("email = ?", value)

	return query
}
