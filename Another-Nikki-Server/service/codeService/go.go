package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func GoRun(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "go.go")
	if util.HandleError(c, err, "写入 go.go 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name go_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'go build go.go 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name go_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i data.in -o data.out ./go >> data.out'", ID, ID)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}

func GoJudge(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "go.go")
	if util.HandleError(c, err, "写入 go.go 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name go_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'go build go.go 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name java_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i ../onlineJudge/problemData/%s/input.in -a ../onlineJudge/problemData/%s/output.out -o out.out -t 2000 ./go > data.out'", ID, ID, code.ProblemName, code.ProblemName)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}
