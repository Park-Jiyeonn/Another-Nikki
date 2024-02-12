package biz

type Problems struct {
	ProblemId    int64
	ProblemTitle string
	CreateTime   string
}
type ProblemRepo interface {
	CreateProblem(problemId int64, title, description, content string) (err error)
	GetProblemById(problemId int64) (problemTitle, problemDescription, problemContent, createTime string, err error)
	GetProblemByPage(pageNum, pageSize int64) (problems []*Problems, err error)
}
