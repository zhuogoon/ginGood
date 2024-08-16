package database

import (
	"fmt"
	"log"
	"youHUA/config"
	"youHUA/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.AppConfig.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	fmt.Println("Database connected successfully")

	// 可以在这里执行自动迁移
	err = DB.AutoMigrate(&models.RequestBody{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
