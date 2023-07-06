package model

type Problem struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ProblemUpdate struct {
	Problem
	ID int `json:"ID"`
}
