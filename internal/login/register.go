package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"youHUA/database"
	"youHUA/models"
)

func Register(c *gin.Context) {
	var input models.Reg
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
	}
	bo := checkUsernameExists(database.DB, input.Username)
	if bo {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户名已存在"})
		return
	}
	// 使用封装的数据库实例
	result := database.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 查询用户名是否已经存在
func checkUsernameExists(db *gorm.DB, username string) bool {
	var count int64
	db.Model(&models.Reg{}).Where("username = ?", username).Count(&count)
	return count > 0
}
