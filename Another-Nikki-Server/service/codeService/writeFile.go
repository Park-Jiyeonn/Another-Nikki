package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteFile(ID int64) {
	dirPath := fmt.Sprintf("./code/tmp-%d", ID)
	_ = os.RemoveAll(dirPath)
}

func GetRetInFile(c *gin.Context, ID int64) (string, bool) {
	ret, err:= os.ReadFile(fmt.Sprintf("./code/tmp-%d/data.out", ID))
	if util.HandleError(c, err, "read data.out failed") {
		return "", false
	}
	return string(ret), true
}

func WriteCodeInFile(code *model.Code, ID int64, FileName string) error {
	dirPath := fmt.Sprintf("./code/tmp-%d", ID)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	cppFileName := dirPath + "/" + FileName
	f, err := os.OpenFile(cppFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, code.Code) //写入文件(字符串)
	if err != nil {
		return err
	}

	inputFileName := dirPath + "/data.in"
	f1, err := os.OpenFile(inputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		return err
	}
	defer f1.Close()
	_, err = io.WriteString(f1, code.Input) //写入文件(字符串)
	return err
}