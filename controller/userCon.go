package controller

import (
	"fmt"
	"loginimpl/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterIndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func LoginIndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ModifyIndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "modify.html", nil)
}

//注册新用户
//待处理：用户注册之前先要对用户名进行查重，对格式进行认证，不允许空格存在
//待处理：密码加密传输，md5？
func PostController(c *gin.Context) {
	var user model.User
	c.ShouldBind(&user)
	fmt.Println(user)
	err := model.CreateAUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Create fail ": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, "创建成功")
	}
}

//获取用户的信息,根据用户名
func GetController(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := model.GetAUser(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, "用户名或密码错误")
		return
	} else {
		if user.Password == password {
			c.JSON(http.StatusOK, gin.H{
				"WELCOME": user.Name,
			})
		}
	}
}

//修改用户的信息,通常为修改密码
//功能：验证旧密码，更新新密码 √
func PutController(c *gin.Context) {
	username := c.PostForm("username") //用户名
	//根据用户名从数据库中查询旧密码
	user, err := model.GetAUser(username)
	//验证
	if err != nil {
		c.JSON(http.StatusUnauthorized, "用户名错误")
		return
	}
	oldpassword := c.PostForm("oldpassword")
	if oldpassword != user.Password {
		c.JSON(http.StatusUnauthorized, "旧密码错误")
		return
	}
	user.Password = c.PostForm("newpassword")
	err = model.UpdateAUserPassword(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "密码更新失败")
		return
	}
	//更新用户新密码
	c.JSON(http.StatusOK, "密码修改成功！")
}

//删除用户的信息
//根据用户名删除对应用户的信息
func DELETEController(c *gin.Context) {
	username := c.Param("username")
	err := model.DeleteAUser(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "删除用户失败")
	}
	c.JSON(http.StatusOK, "删除用户成功")
	//	c.JSON(http.StatusServiceUnavailable, "our base is under attack")
}
