package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"youHUA/database"
	"youHUA/models"
)

func Login(c *gin.Context) {
	var log models.Reg
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
	}
	bo := checkUsernameExists(database.DB, log.Username)
	if !bo {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请先注册"})
		return
	}

	pwd, err := findPwd(log.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "未查询到账号数据"})
		return
	}
	if pwd == log.Password {
		c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
		c.HTML(http.StatusOK, "index.html", gin.H{"username": log.Username})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码错误"})
		return
	}
}

// 通过用户名获取密码
func findPwd(username string) (string, error) {
	var db models.Reg
	err := database.DB.Model(&models.Reg{}).
		Select("password").
		Where("username = ?", username).
		First(&db).Error
	if err != nil {
		return "", err
	}
	return db.Password, nil
}
