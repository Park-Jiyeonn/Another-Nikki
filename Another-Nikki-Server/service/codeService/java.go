package service

import (
	"Another-Nikki/router/model"
	"Another-Nikki/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func JavaRun(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "Main.java")
	if util.HandleError(c, err, "写入 Main.java 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name java_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'javac 'Main.java' 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name java_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i data.in -o data.out java Main >> data.out'", ID, ID)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		// fmt.Println("compile_log = ", compile_log)
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}

func JavaJudge(c *gin.Context, code *model.Code, ID string) bool {
	err := WriteCodeInFile(code, ID, "Main.java")
	if util.HandleError(c, err, "写入 Main.java 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name java_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'javac 'Main.java' 2> compile.log'", ID, ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name java_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i ../onlineJudge/problemData/%s/input.in -a ../onlineJudge/problemData/%s/output.out -o out.out -t 2000 java Main > data.out'", ID, ID, code.ProblemName, code.ProblemName)

	err = RunCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		// fmt.Println("compile_log = ", compile_log)
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = RunCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}
