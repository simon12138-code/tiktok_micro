package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"tiktok/cmd/video/config"
)

var DB *gorm.DB

// Init init DB
func InitDB() {
	var err error
	mysqlInfo := config.MysqlConfig{
		Name:     config.Settings.Mysqlinfo.Name,
		Host:     config.Settings.Mysqlinfo.Host,
		Password: config.Settings.Mysqlinfo.Password,
		DBName:   config.Settings.Mysqlinfo.DBName,
		Port:     config.Settings.Mysqlinfo.Port,
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
		mysqlInfo.Port, mysqlInfo.DBName)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// New constructs a new plugin based opentracing. It supports to trace all operations in gorm,
	// so if you have already traced your servers, now this plugin will perfect your tracing job.
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}
