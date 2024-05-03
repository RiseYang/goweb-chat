package user_service

import (
	"github.com/gin-gonic/gin"
	"goweb-blog/models"
	"goweb-blog/services/helper"
	"goweb-blog/services/session"
	"goweb-blog/services/validator"
	"net/http"
	"strconv"
)

// 用户登录
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	avatarId := c.PostForm("avatar_id")

	var u validator.User

	u.Username = username
	u.Password = password
	u.AvatarId = avatarId

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
	}

	user := models.FindUserByField("username", username)
	userInfo := user
	md5Pwd := helper.Md5Encrypt(password)

	if userInfo.ID > 0 {
		//json存在
		//验证码
		if userInfo.Password != md5Pwd {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "密码错误",
			})
			return
		}
		models.SaveAvatarId(avatarId, user)
	} else {
		//新用户
		userInfo = models.AddUser(map[string]interface{}{
			"username":  username,
			"password":  md5Pwd,
			"avatar_id": avatarId,
		})
	}

	if userInfo.ID > 0 {
		session.SaveAuthSession(c, string(strconv.Itoa(int(userInfo.ID))))
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "系统错误",
		})
		return
	}

}

func GetUserInfo(c *gin.Context) map[string]interface{} {
	return session.GetSessionUserInfo(c)
}

// 退出
func Logout(c *gin.Context) {
	session.ClearAuthSession(c)
	c.Redirect(http.StatusFound, "/")
	return
}
