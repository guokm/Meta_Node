package dto

type AddComment struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"blog_id" binding:"required"`
}

type ListComment struct {
	PostID uint `json:"blog_id" binding:"required"`
	PageDTO
}
