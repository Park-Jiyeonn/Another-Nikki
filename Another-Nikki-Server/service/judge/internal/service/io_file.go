package service

import (
	"Another-Nikki/service/judge/api"
	"fmt"
	"io"
	"os"
	"strings"
)

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
	dirPath := fmt.Sprintf("./onlineJudge/tmp-%s", ID)
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

func readJudgeRet(ID string) (resp *api.JudgeResp, err error) {
	defer deleteFile(ID)
	resp = new(api.JudgeResp)
	var ret []byte
	if ret, err = os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/data.out", ID)); err != nil {
		return
	}
	retString := string(ret)
	ans := strings.Split(retString, "\n")
	n := len(ans)
	resp.MemoryUsed = ans[n-3] + " kb"
	resp.CpuTimeUsed = ans[n-5] + " ms"
	resp.JudgeResult = ans[n-1]
	return
}

func readRunRet(ID string) (resp *api.OnlineRunResp, err error) {
	defer deleteFile(ID)
	resp = new(api.OnlineRunResp)
	var ret []byte
	if ret, err = os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/data.out", ID)); err != nil {
		return
	}
	retString := string(ret)
	ans := strings.Split(retString, "\n")
	n := len(ans)
	resp.CompileState = "success"
	resp.Output = strings.Join(ans[:n-5], "\n")
	resp.MemoryUsed = ans[n-3] + " kb"
	resp.CpuTimeUsed = ans[n-5] + " ms"
	return
}

func deleteFile(ID string) {
	dirPath := fmt.Sprintf("./onlineJudge/tmp-%s", ID)
	_ = os.RemoveAll(dirPath)
}
