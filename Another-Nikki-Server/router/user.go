package router

import (
	"Another-Nikki/dal"
	"Another-Nikki/router/model"
	"Another-Nikki/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var user dal.User

func Login(c *gin.Context) {
	var userLogin model.UserLogin
	err := c.BindJSON(&userLogin)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	user, err := user.GetUserByName(userLogin.UserName)
	if util.HandleError(c, err, "数据库查询用户存不存在失败！") {
		return
	}
	if user == nil {
		util.SendResp(c, 404, nil, "用户不存在")
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
	util.SendResp(c, 200, gin.H{
		"token":     token,
		"user_id":   user.ID,
		"user_name": user.Username,
		"avatar":    user.Avatar,
	}, "success")
}

func Register(c *gin.Context) {
	var userRg model.UserRegister
	err := c.BindJSON(&userRg)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	if userRg.Password1 != userRg.Password2 {
		util.SendResp(c, 404, nil, "密码不一致捏")
		return
	}
	if user, _ := user.GetUserByName(userRg.UserName); user != nil {
		util.SendResp(c, 404, nil, "用户已存在")
		return
	}

	EncodePassword, err := bcrypt.GenerateFromPassword([]byte(userRg.Password1), bcrypt.DefaultCost)
	if util.HandleError(c, err, "密码加密失败") {
		return
	}
	id, err := user.CreateUser(userRg.UserName, string(EncodePassword))
	if util.HandleError(c, err, "创建用户失败") {
		return
	}

	token, err := util.GenToken(id, userRg.UserName)
	if util.HandleError(c, err, "jwt 生成 Token 失败") {
		return
	}

	util.SendResp(c, 200, gin.H{
		"token":     token,
		"user_id":   id,
		"user_name": userRg.UserName,
	}, "success")
}
