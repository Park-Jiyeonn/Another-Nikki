package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var userLogin model.UserLogin
	err := c.BindJSON(&userLogin)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	user, err := dal.GetUserByName(context.Background(), userLogin.UserName)
	if util.HandleError(c, err, "数据库查询用户存不存在失败！") {
		return
	}
	if user == nil {
		util.HandleError(c, util.NewErrNo("用户不存在"), "")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if util.HandleError(c, err, "用户名或密码错误") {
		return
	}

	token, err := util.GenToken(int(user.ID), userLogin.UserName)
	if util.HandleError(c, err, "颁发 Token 失败") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"data": gin.H {
			"token":token,
		},
		"message":"success",
	})
}

func Register(c *gin.Context) {
	var userRg model.UserRegister
	err := c.BindJSON(&userRg)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	if userRg.Password1 != userRg.Password2 {
		util.HandleError(c, util.NewErrNo("密码不一致捏"), "")
		return
	}
	if user, _ := dal.GetUserByName(context.Background(), userRg.UserName); user != nil {
		util.HandleError(c, util.NewErrNo("用户已存在"), "")
		return
	}

	EncodePassword, err := bcrypt.GenerateFromPassword([]byte(userRg.Password1), bcrypt.DefaultCost)
	if util.HandleError(c, err, "密码加密失败") {
		return
	}
	id, err := dal.CreateUser(context.Background(), userRg.UserName, string(EncodePassword))
	if util.HandleError(c, err, "创建用户失败") {
		return
	}


	token, err := util.GenToken(id, userRg.UserName)
	if util.HandleError(c, err, "jwt 生成 Token 失败") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":200,
		"data": gin.H {
			"token":token,
		},
		"message":"success",
	})
}