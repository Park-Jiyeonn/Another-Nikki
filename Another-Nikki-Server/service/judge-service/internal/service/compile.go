package service

import (
	"Another-Nikki/service/judge-service/api"
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	DataFileName   = "data.in"
	CppFileName    = "c++.cpp"
	JavaFileName   = "java.java"
	GolangFileName = "golang.go"
)

func compile(ID, code, input string, language api.Language) (err error) {
	var (
		cmd        string
		compileLog []byte
	)
	switch language {
	case api.Language_Cpp:
		cmd, err = cpp_compile(ID, code, input)
	case api.Language_Java:
		cmd, err = java_compile(ID, code, input)
	case api.Language_Python:
		return
	case api.Language_Golang:
		cmd, err = golang_compile(ID, code, input)
	}
	if err != nil {
		return
	}
	if err = runCommand(cmd); err != nil {
		return
	}
	compileLog, err = os.ReadFile(fmt.Sprintf("./code/tmp-%s/compile.log", ID))
	if len(compileLog) != 0 {
		err = fmt.Errorf(string(compileLog))
		return
	}
	return
}

func cpp_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, CppFileName); err != nil {
		return
	}
	cmd = fmt.Sprintf("docker run --rm -m 256m --name cpp_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log'", ID, ID)
	return
}

func java_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, JavaFileName); err != nil {
		return
	}
	cmd = fmt.Sprintf("docker run --rm -m 256m --name java_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'javac 'Main.java' 2> compile.log'", ID, ID)
	return
}

func golang_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, GolangFileName); err != nil {
		return
	}
	cmd = fmt.Sprintf("docker run --rm -m 256m --name go_compile-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c 'go build go.go 2> compile.log'", ID, ID)
	return
}

func runCommand(s string) error {
	cmd := exec.Command("bash", "-c", s)
	err := cmd.Run()
	return err
}

func writeInFile(ID, code, input, filename string) (err error) {
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
	path := dirPath + "/" + filename
	f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer f.Close()
	if _, err = io.WriteString(f, content); err != nil {
		return
	}
	return
}
