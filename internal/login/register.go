package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"youHUA/database"
	"youHUA/models"
)

// Register 注册
// @Tags 注册
// @Summary 注册
// @Description 注册
// @Router /register [post]
// @Param request body models.RequestBody true "username"
// @Success 200 {object} models.ResponseBody
func Register(c *gin.Context) {
	var req models.RequestBody
	var resp models.ResponseBody
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = http.StatusBadRequest
		resp.Msg = "参数错误"
		c.AbortWithStatusJSON(0, resp)
		return
	}

	if database.DB == nil {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "Database connection failed"
		c.AbortWithStatusJSON(0, resp)
		return
	}
	bo := checkUsernameExists(database.DB, req.Username)
	if bo {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "已经注册"
		c.AbortWithStatusJSON(0, resp)
		return
	}
	// 使用封装的数据库实例
	result := database.DB.Create(&req)
	if result.Error != nil {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "注册失败"
		c.AbortWithStatusJSON(0, resp)
		return
	}
	resp.Code = http.StatusOK
	resp.Msg = "注册成功"
	c.JSON(0, resp)
}

// 查询用户名是否已经存在
func checkUsernameExists(db *gorm.DB, username string) bool {
	var count int64
	db.Model(&models.RequestBody{}).Where("username = ?", username).Count(&count)
	return count > 0
}
