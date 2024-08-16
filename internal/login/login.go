package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"youHUA/database"
	"youHUA/models"
)

// Login 登录
// @Tags 登录
// @Summary 登录
// @Description 登录
// @Router /login [post]
// @Param request body models.RequestBody true "username"
// @Success 200 {object} models.ResponseBody
func Login(c *gin.Context) {
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
	if !bo {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "请先注册"
		c.AbortWithStatusJSON(0, resp)
		return
	}

	pwd, err := findPwd(req.Username)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "未查询到账号数据"
		c.AbortWithStatusJSON(0, resp)
		return
	}
	if pwd == req.Password {
		resp.Code = http.StatusOK
		resp.Msg = "登录成功"
		c.AbortWithStatusJSON(0, resp)
		c.HTML(http.StatusOK, "index.html", gin.H{"username": req.Username})
	} else {
		resp.Code = http.StatusInternalServerError
		resp.Msg = "密码错误"
		c.AbortWithStatusJSON(0, resp)
		return
	}
}

// 通过用户名获取密码
func findPwd(username string) (string, error) {
	var db models.RequestBody
	err := database.DB.Model(&models.RequestBody{}).
		Select("password").
		Where("username = ?", username).
		First(&db).Error
	if err != nil {
		return "", err
	}
	return db.Password, nil
}
