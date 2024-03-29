package service

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const OnlineJudgePath = "../../onlineJudge/"

func writeInFile(ID, code, input, filename string) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	if err = writeContentInFile(ID, code, filename); err != nil {
		return
	}
	if err = writeContentInFile(ID, input, DataFileName); err != nil {
		return
	}
	return
}

func writeContentInFile(ID, content, filename string) (err error) {
	var f *os.File
	dirPath := fmt.Sprintf(OnlineJudgePath+"tmp-%s", ID)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}
	path := dirPath + "/" + filename
	f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	if _, err = io.WriteString(f, content); err != nil {
		return
	}
	return
}

func readJudgeRet(ID string) (memoryUsed, cpuTimeUsed, JudgeResult string, err error) {
	defer deleteFile(ID)
	var ret []byte
	if ret, err = os.ReadFile(fmt.Sprintf(OnlineJudgePath+"tmp-%s/data.out", ID)); err != nil {
		return
	}
	retString := string(ret)
	ans := strings.Split(retString, "\n")
	n := len(ans)
	memoryUsed = ans[n-3] + " kb"
	cpuTimeUsed = ans[n-5] + " ms"
	JudgeResult = ans[n-1]
	return
}

func readRunRet(ID string) (output, memoryUsed, cpuTimeUsed string, err error) {
	defer deleteFile(ID)
	var ret []byte
	if ret, err = os.ReadFile(fmt.Sprintf(OnlineJudgePath+"tmp-%s/data.out", ID)); err != nil {
		return
	}
	retString := string(ret)
	ans := strings.Split(retString, "\n")
	n := len(ans)
	output = strings.Join(ans[:n-5], "\n")
	memoryUsed = ans[n-3] + " kb"
	cpuTimeUsed = ans[n-5] + " ms"
	return
}

func deleteFile(ID string) {
	dirPath := fmt.Sprintf(OnlineJudgePath+"tmp-%s", ID)
	_ = os.RemoveAll(dirPath)
}

func checkPythonSyntaxError(ID string) (string, bool) {
	ret, err := os.ReadFile(fmt.Sprintf(OnlineJudgePath+"tmp-%s/data.out", ID))
	if err != nil {
		return "", false
	}
	retString := string(ret)
	if strings.Contains(retString, "Error") {
		ans := strings.Split(retString, "\n")
		deleteFile(ID)
		return strings.Join(ans[:len(ans)-5], "\n"), true
	}
	return "", false
}
