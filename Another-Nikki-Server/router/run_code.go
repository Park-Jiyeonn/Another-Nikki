package router

import (
	"Another-Nikki/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Code struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func writeCodeInFile(code *Code, ID int64) error {
	dirPath := fmt.Sprintf("./code/tmp-%d", ID)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	filename := dirPath + "/c++.cpp"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, code.Code) //写入文件(字符串)
	return err
}

func deleteFile(ID int64) {
	dirPath := fmt.Sprintf("./code/tmp-%d", ID)
	_ = os.RemoveAll(dirPath)
}

func runCommand(s string) error {
	cmd := exec.Command("bash", "-c", s)
	err := cmd.Run()
	return err
}

func RunCode(c *gin.Context) {
	var code Code
	err := c.BindJSON(&code)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	if code.Lang != "c++" {
		c.JSON(http.StatusOK, gin.H{
			"state":"error",
			"message": "暂不支持这种语言呢～",
		})
		return
	}

	ID := time.Now().UnixNano()
	err = writeCodeInFile(&code, ID)
	defer deleteFile(ID)
	if util.HandleError(c, err, "写入c++.cpp失败") {
		return
	}


	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%d -v $(pwd)/code/tmp-%d:/dox cpp_env:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID,ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%d -v $(pwd)/code/tmp-%d:/dox cpp_env:1 sh -c 'ls ./.. && ./../calc/calc ./cpp > a.out'", ID,ID)
	// fmt.Println(cmd1)
	// fmt.Println(cmd2)

	err = runCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile("./code/compile.log")
		util.HandleError(c, err, string(compile_log))
		return
	}
	err = runCommand(cmd2)
	if util.HandleError(c, err, "run failed") {
		return
	}

	ret, err:= os.ReadFile(fmt.Sprintf("./code/tmp-%d/a.out", ID))
	if util.HandleError(c, err, "read a.out failed") {
		return
	}

	ans := strings.Split(string(ret), "\n")
	n := len(ans)
	c.JSON(http.StatusOK, gin.H{
		"state":"success",
		"message":strings.Join(ans[:n-3], ""),
		"cpu_time_used":ans[n-3],
		"memory_used":ans[n-2],
		"exit_code":ans[n-1],
	})
}
