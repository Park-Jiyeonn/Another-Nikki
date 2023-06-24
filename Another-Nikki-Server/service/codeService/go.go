package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Go(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "go.go")
	if util.HandleError(c, err, "写入 go.go 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name go_compile-%s -v $(pwd)/code/tmp-%s:/dox oj:1 sh -c 'go build go.go 2> compile.log'", ID,ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name go_run-%s -v $(pwd)/code/tmp-%s:/dox oj:1 sh -c './../calc/calc2 ./go < data.in > data.out'", ID,ID)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./code/tmp-%s/compile.log", ID))
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}