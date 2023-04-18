package queries_v1

import (
	"gorm.io/gorm"
)

func ById(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	query = query.Where("id = ?", value)

	return query
}
