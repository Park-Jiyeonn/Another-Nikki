package model

type Code struct {
	ProblemName string `json:"problem_name"`
	Input       string `json:"input"`
	Code        string `json:"code"`
	Lang        string `json:"lang"`
}
