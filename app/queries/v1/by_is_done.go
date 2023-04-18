package queries_v1

import (
	"strconv"

	"gorm.io/gorm"
)

func ByIsDone(value string, query *gorm.DB) *gorm.DB {
	if value == "" {
		return query
	}

	boolValue, _ := strconv.ParseBool(value)
	query = query.Where("is_done = ?", boolValue)

	return query
}
