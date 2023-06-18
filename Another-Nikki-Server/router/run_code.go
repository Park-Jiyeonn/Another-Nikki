package router

import (
	"Another-Nikki/util"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Code struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}

func writeCodeInFile(code *Code) error {
	var filename = "./code/c++.cpp"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.WriteString(f, code.Code) //写入文件(字符串)
	return err
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

	err = writeCodeInFile(&code)
	if util.HandleError(c, err, "写入c++.cpp失败") {
		return
	}

	err = runCommand("docker run --rm -m 256m --name cpp_compile -v $(pwd)/code:/dox cpp_env:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'")
	if util.HandleError(c, err, "compile failed") {
		return
	}
	err = runCommand("docker run --rm -m 256m --name cpp_run -v $(pwd)/code:/dox cpp_env:1 sh -c './calc ./cpp > a.out'")
	if util.HandleError(c, err, "run failed") {
		return
	}

	ret, err:= os.ReadFile("./code/a.out")
	if util.HandleError(c, err, "read a.out failed") {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":"success",
		"message":string(ret),
	})
}
