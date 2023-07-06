package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func CppRun(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "c++.cpp")
	if util.HandleError(c, err, "写入 c++.cpp 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i data.in -o data.out ./cpp >> data.out'", ID, ID)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./code/tmp-%s/compile.log", ID))
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}

func CppJudge(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "c++.cpp")
	if util.HandleError(c, err, "写入 c++.cpp 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i ../onlineJudge/problemData/%s/input.in -a ../onlineJudge/problemData/%s/output.out -o out.out -t 1000 ./cpp > data.out'", ID, ID, code.ProblemName, code.ProblemName)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run failed")
}
