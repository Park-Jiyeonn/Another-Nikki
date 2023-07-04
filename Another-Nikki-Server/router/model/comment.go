package model

type Comment struct {
	Content      string `json:"content"`
	AuthorID     int    `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
	ParentID     int    `json:"parent_id"`
	RootID       int    `json:"root_id"`
	ParentName   string `json:"parent_name"`
}
