package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Cpp(c *gin.Context, code *model.Code, ID int64) bool {
	err := WriteCodeInFile(code, ID, "c++.cpp")
	if util.HandleError(c, err, "写入 c++.cpp 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID,ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'ls ./.. && ./../calc/calc1 ./cpp < data.in > data.out'", ID,ID)
	// fmt.Println(cmd1)
	// fmt.Println(cmd2)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile("./code/compile.log")
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}