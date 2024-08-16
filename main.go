package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"youHUA/config"
	"youHUA/database"
	"youHUA/router"
)

func main() {

	r := gin.Default()

	// 读取数据库配置文件
	err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load database configuration: ", err)
	}

	database.InitDB()

	router.Router(r)

	err = r.Run(":8080")
	if err != nil {
		fmt.Println("启动失败")
	}
}
