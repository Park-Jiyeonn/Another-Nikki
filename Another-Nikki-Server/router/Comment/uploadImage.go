package Comment

import (
	"Another-Nikki/util"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if util.HandleError(c, err, "参数填写错误") {
		return
	}
	path := "./tmp/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if util.HandleError(c, err, "文件保存失败") {
		return
	}
	url, err := util.UploadCos(path, file.Filename)
	if util.HandleError(c, err, "文件上传腾讯云失败") {
		return
	}
	util.SendResp(c, 200, url, "上传成功")
}
