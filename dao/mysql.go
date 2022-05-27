package dao

import (
	"todo-list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



var (
	DB *gorm.DB
	// TODO: read from the config
	SQLLogMode logger.LogLevel = logger.Silent
)

func InitMySQL() (err error) {
	dsn := "root:1qaz9ol.@(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(SQLLogMode),
	})
	return err
}

func AutoMigrate() error {
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		return err
	}
	return nil
}