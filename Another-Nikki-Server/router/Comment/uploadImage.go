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
	if file.Size > 20*1024*1024 {
		util.SendResp(c, 404, "", "文件过大，请上传大小在 20MB 以内的文件")
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
