package initialize

import (
	"Noteus/config"
	"Noteus/global"
	"Noteus/initialize/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() (db *gorm.DB) {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	var err error
	if db, err = gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return
	}
	return
}

func GormMysqlByConfig(m config.Mysql) (db *gorm.DB) {
	if m.Dbname == "" {
		return
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	var err error
	if db, err = gorm.Open(mysql.New(mysqlConfig)); err != nil {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return
	}
	return
}
