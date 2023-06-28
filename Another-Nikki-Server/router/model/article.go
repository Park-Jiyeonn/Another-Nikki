package model
type Article struct {
	Title string `json:"title"`
    Description string `json:"description"`
    Content string `json:"content"`
}

type ArticleUpdate struct {
    ID  int  `json:"ID"`
	Title string `json:"title"`
    Description string `json:"description"`
    Content string `json:"content"`
}