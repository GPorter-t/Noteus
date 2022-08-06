package initialize

import (
	"Noteus/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	var db *gorm.DB
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		db = GormMysql()
	default:
		db = GormMysql()
	}
	return db
}
