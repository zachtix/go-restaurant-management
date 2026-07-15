package helper

import (
	"gorm.io/gorm"
)

func CountTotal(query *gorm.DB, limit int) (total int64, totalPage int64, err error) {
	if err := query.Count(&total).Error; err != nil {
		return 0, 0, err
	}
	if limit <= 0 {
		limit = 1
	}
	totalPage = (total + int64(limit) - 1) / int64(limit)
	return total, totalPage, nil
}
