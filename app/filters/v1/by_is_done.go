package filters_v1

import (
	"strconv"

	"gorm.io/gorm"
)

func ByIsDone(filter map[string]interface{}, query *gorm.DB) *gorm.DB {
	var key = "is_done"

	if value, exists := filter[key]; exists && value != "" {
		value, _ = strconv.ParseBool(filter[key].(string))
		query = query.Where(
			"id_done = ?",
			value,
		)
	}

	return query
}
