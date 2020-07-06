package model

type Post struct {
	PostID   int64  `json:"id,omitempty"`
	AuthorID int64  `json:"author_id,omitempty"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
}
