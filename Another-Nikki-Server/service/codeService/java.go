package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Java(c *gin.Context, code *model.Code, ID int64) bool {
	err := WriteCodeInFile(code, ID, "Main.java")
	if util.HandleError(c, err, "写入 Main.java 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name java_compile-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'javac 'Main.java' 2> compile.log'", ID,ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name java_run-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c './../calc/calc2 java Main < data.in > data.out'", ID,ID)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile("./code/compile.log")
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}