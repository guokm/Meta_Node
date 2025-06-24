package dto

type AddPostDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostDto struct {
	ID      string `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type FindPostDto struct {
	PageDTO
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetOrDeletePostDto struct {
	ID string `json:"id" binding:"required"`
}
