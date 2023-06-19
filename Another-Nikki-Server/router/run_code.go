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
	Input	string `json:"input"`
	Code 	string `json:"code"`
	Lang 	string `json:"lang"`
}

func writeCodeInFile(code *Code, ID int64, FileName string) error {
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

func deleteFile(ID int64) {
	dirPath := fmt.Sprintf("./code/tmp-%d", ID)
	_ = os.RemoveAll(dirPath)
}

func runCommand(s string) error {
	cmd := exec.Command("bash", "-c", s)
	err := cmd.Run()
	return err
}

func python(c *gin.Context, code *Code, ID int64) bool {
	err := writeCodeInFile(code, ID, "py.py")
	if util.HandleError(c, err, "写入 py.py 或 data.in 失败") {
		return false
	}
	cmd := fmt.Sprintf("docker run --rm -m 256m --name py_run-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'ls ./.. && ./../calc/calc2 python3 py.py < data.in > data.out'", ID,ID)
	err = runCommand(cmd)
	return !util.HandleError(c, err, "run faild")
}

func cpp(c *gin.Context, code *Code, ID int64) bool {
	err := writeCodeInFile(code, ID, "c++.cpp")
	if util.HandleError(c, err, "写入 c++.cpp 或 data.in 失败") {
		return false
	}
	cmd1 := fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID,ID)
	cmd2 := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%d -v $(pwd)/code/tmp-%d:/dox oj:1 sh -c 'ls ./.. && ./../calc/calc1 ./cpp < data.in > data.out'", ID,ID)
	// fmt.Println(cmd1)
	// fmt.Println(cmd2)

	err = runCommand(cmd1)
	if err != nil {
		compile_log, _ := os.ReadFile("./code/compile.log")
		util.HandleError(c, err, string(compile_log))
		return false
	}
	err = runCommand(cmd2)
	return !util.HandleError(c, err, "run faild")
}

func RunCode(c *gin.Context) {
	var code Code
	err := c.BindJSON(&code)
	if util.HandleError(c, err, "参数错误") {
		return
	}

	ID := time.Now().UnixNano()
	defer deleteFile(ID)
	
	if code.Lang == "c++" {
		if !cpp(c, &code, ID) {
			return
		}
	} else if code.Lang == "python" {
		if !python(c, &code, ID) {
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state":"error",
			"message": "暂不支持这种语言呢～",
		})
		return
	}

	ret, err:= os.ReadFile(fmt.Sprintf("./code/tmp-%d/data.out", ID))
	if util.HandleError(c, err, "read data.out failed") {
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
