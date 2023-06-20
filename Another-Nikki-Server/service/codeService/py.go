package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Python(c *gin.Context, code *model.Code, ID int64) bool {
	err := WriteCodeInFile(code, ID, "py.py")
	if util.HandleError(c, err, "写入 py.py 或 data.in 失败") {
		return false
	}
	cmd := fmt.Sprintf("docker run --rm -m 256m --name py_run-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'ls ./.. && ./../calc/calc2 python3 py.py < data.in > data.out'", ID,ID)
	err = RunCommand(cmd)
	return !util.HandleError(c, err, "run faild")
}