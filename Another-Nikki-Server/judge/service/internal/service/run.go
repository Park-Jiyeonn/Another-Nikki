package service

import (
	"Another-Nikki/judge/service/api"
	"Another-Nikki/pkg/log"
	"fmt"
	"golang.org/x/net/context"
)

const (
	CppAppName    = "c++"
	JavaAppName   = "Main"
	GolangAppName = "golang"
)

const runCmd = "docker exec oj sh -c 'cd tmp-%s; ../judger -i data.in -o data.out -T 5000 "

func run(ctx context.Context, ID string, language api.Language) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	var cmd string
	switch language {
	case api.Language_Cpp:
		cmd = fmt.Sprintf(runCmd+"./%s >> data.out'", ID, CppAppName)
	case api.Language_Java:
		cmd = fmt.Sprintf(runCmd+"java %s >> data.out'", ID, JavaAppName)
	case api.Language_Python:
		cmd = fmt.Sprintf(runCmd+"-i data.in python3 -u %s >> data.out 2>&1'", ID, PythonFileName)
	case api.Language_Golang:
		cmd = fmt.Sprintf(runCmd+"./%s >> data.out'", ID, GolangAppName)
	}
	err = runCommand(ctx, cmd)
	if err != nil {
		log.Error(ctx, "run code error: %v", err)
	}
	return
}

const judgeCmd = "docker exec oj sh -c 'cd tmp-%s; ../judger -i ../problemData/%s/input.in -a ../problemData/%s/output.out -o out.out -T 5000 "

func judge(ctx context.Context, ID, problemName string, language api.Language) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	var cmd string
	switch language {
	case api.Language_Cpp:
		cmd = fmt.Sprintf(judgeCmd+"-t 1000 ./%s > data.out'", ID, problemName, problemName, CppAppName)
	case api.Language_Java:
		cmd = fmt.Sprintf(judgeCmd+"-t 2000 java %s > data.out'", ID, problemName, problemName, JavaAppName)
	case api.Language_Python:
		cmd = fmt.Sprintf(judgeCmd+"-t 2000 python3 -u %s >> data.out 2>&1'", ID, problemName, problemName, PythonFileName)
	case api.Language_Golang:
		cmd = fmt.Sprintf(judgeCmd+"-t 2000 ./%s > data.out'", ID, problemName, problemName, GolangAppName)
	}
	err = runCommand(ctx, cmd)
	if err != nil {
		log.Error(ctx, "judge code error: %v", err)
	}
	return
}
