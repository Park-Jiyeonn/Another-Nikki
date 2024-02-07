package service

import (
	"Another-Nikki/judge/service/api"
	"fmt"
)

const (
	CppAppName    = "c++"
	JavaAppName   = "Main"
	GolangAppName = "golang"
)

func run(ID string, language api.Language) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	var cmd string
	switch language {
	case api.Language_Cpp:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i data.in -o data.out ./%s >> data.out'", ID, CppAppName)
	case api.Language_Java:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i data.in -o data.out java %s >> data.out'", ID, JavaAppName)
	case api.Language_Python:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i data.in -o data.out -i data.in python3 -u %s >> data.out 2>&1'", ID, PythonFileName)
	case api.Language_Golang:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i data.in -o data.out ./%s >> data.out'", ID, GolangAppName)
	}
	err = runCommand(cmd)
	return
}

func judge(ID, problemName string, language api.Language) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	var cmd string
	switch language {
	case api.Language_Cpp:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i ../../onlineJudge/problemData/%s/input.in -a ../../onlineJudge/problemData/%s/output.out -o out.out -t 1000 ./%s > data.out'", ID, problemName, problemName, CppAppName)
	case api.Language_Java:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i ../../onlineJudge/problemData/%s/input.in -a ../../onlineJudge/problemData/%s/output.out -o out.out -t 2000 java %s > data.out'", ID, problemName, problemName, JavaAppName)
	case api.Language_Python:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i ../../onlineJudge/problemData/%s/input.in -a ../../onlineJudge/problemData/%s/output.out -o out.out -t 2000 python3 -u %s >> data.out 2>&1'", ID, problemName, problemName, PythonFileName)
	case api.Language_Golang:
		cmd = fmt.Sprintf("docker exec oj sh -c 'cd tmp-%s; ../../onlineJudge/judger -i ../../onlineJudge/problemData/%s/input.in -a ../../onlineJudge/problemData/%s/output.out -o out.out -t 2000 ./%s > data.out'", ID, problemName, problemName, GolangAppName)
	}
	err = runCommand(cmd)
	return
}
