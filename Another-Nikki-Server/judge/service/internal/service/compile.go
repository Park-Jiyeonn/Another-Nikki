package service

import (
	"Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"os/exec"
)

const (
	DataFileName   = "data.in"
	CppFileName    = "c++.cpp"
	JavaFileName   = "Main.java"
	GolangFileName = "golang.go"
	PythonFileName = "py.py"
)

func compile(ctx context.Context, ID, code, input string, language api.Language) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
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
		_, _ = python_compile(ID, code, input)
		return
	case api.Language_Golang:
		cmd, err = golang_compile(ID, code, input)
	}
	if err != nil {
		log.Error(ctx, "write code into file err: %v", err)
		return
	}
	if err = runCommand(ctx, cmd); err != nil {
		log.Error(ctx, "compile error, err = %v, code = %v, language = %v", err, code, language)
		compileLog, readErr := os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
		log.Error(ctx, "compile log = %v, read compile err = %v", string(compileLog), readErr)
		return
	}
	compileLog, err = os.ReadFile(fmt.Sprintf("./onlineJudge/tmp-%s/compile.log", ID))
	if err != nil {
		log.Error(ctx, "ReadFile err: %v", err)
		return
	}
	if len(compileLog) != 0 {
		err = fmt.Errorf(string(compileLog))
		log.Error(ctx, "compile error: %v", string(compileLog))
		return
	}
	return
}

func cpp_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, CppFileName); err != nil {
		return
	}
	cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; g++ %s -o %s -O2 -std=c++11 2> compile.log'", ID, CppFileName, CppAppName)
	return
}

func java_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, JavaFileName); err != nil {
		return
	}
	cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; javac %s 2> compile.log'", ID, JavaFileName)
	return
}

func golang_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, GolangFileName); err != nil {
		return
	}
	// docker run -d -i -m 256m --name oj -v $(pwd)/onlineJudge:/dox oj:1
	cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; go build %s 2> compile.log'", ID, GolangFileName)
	return
}

func python_compile(ID, code, input string) (cmd string, err error) {
	if err = writeInFile(ID, code, input, PythonFileName); err != nil {
		return
	}
	return
}

func runCommand(ctx context.Context, s string) error {
	cmd := exec.Command("bash", "-c", s)
	err := cmd.Run()
	return err
}
