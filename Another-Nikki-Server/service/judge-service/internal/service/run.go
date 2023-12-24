package service

import "fmt"

func run(ID string) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	cmd := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i data.in -o data.out ./cpp >> data.out'", ID, ID)
	err = runCommand(cmd)
	return
}

func judge(ID, problemName string) (err error) {
	defer func() {
		if err != nil {
			deleteFile(ID)
		}
	}()
	cmd := fmt.Sprintf("docker run --rm -m 256m --name cpp_run-%s -v $(pwd)/onlineJudge/tmp-%s:/dox oj:1 sh -c '../onlineJudge/judger -i ../onlineJudge/problemData/%s/input.in -a ../onlineJudge/problemData/%s/output.out -o out.out -t 1000 ./cpp > data.out'", ID, ID, problemName, problemName)
	err = runCommand(cmd)
	return
}
