package models

type PostData struct {
	ID        int    `json:"userId"`
	TodoTitle string `json:"title"`
}

type AddPost struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}
