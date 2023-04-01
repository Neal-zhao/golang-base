package models

type ArticleInfo struct {
	Id int64 `json:"id" db:"id"`
	CategoryId int64 `json:"category_id" db:"category_id"`
	Title string `json:"title" db:"title"`
	CommentCount int64 `json:"comment_count" db:"comment_count"`
	ViewCount int64 `json:"view_count" db:"view_count"`
	Username string `json:"username" db:"username"`
	Status int8 `json:"status" db:"status"`
	Summary string `json:"summary" db:"summary"`
	//CreateTime time.Duration `json:"summary" db:"summary"`
	//UpdateTime time.Duration `json:"summary" db:"summary"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}
type ArticleRecord struct {
	ArticleInfo
	Category
}