package utils

import "gorm.io/gorm"

// GetPageData 获取分页数据
func GetPageData(page, pageSize uint) (limit, offset uint) {
	limit = pageSize
	offset = pageSize * (page - 1)
	return
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 1
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateWithSkip(page, pageSize, skip int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 1
		}
		if skip < 0 {
			skip = 0
		}

		offset := skip + (page-1)*pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PreloadChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", func(db *gorm.DB) *gorm.DB {
		return PreloadChildren(db) // 递归调用自己
	})
}
