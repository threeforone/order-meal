package global

import "gorm.io/gorm"

var (
	DB *gorm.DB
)

func CloseDB() {
	if DB != nil {
		if db, _ := DB.DB(); db != nil {
			defer db.Close()
		}
	}
}
