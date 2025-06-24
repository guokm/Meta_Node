package controller

import (
	"github/metanode/go_blog/config"
	"github/metanode/go_blog/model/dto"
	"github/metanode/go_blog/model/entity"
	"github/metanode/go_blog/model/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatComment(c *gin.Context) {
	var addComment dto.AddComment
	if err := c.ShouldBindJSON(&addComment); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}
	userId := c.GetUint("ID")
	comment := entity.Comment{
		Content: addComment.Content,
		UserID:  userId,
		PostID:  addComment.PostID,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, vo.Error(vo.CodeParamError, "Failed to create comment"))
		return
	}
	c.JSON(http.StatusOK, vo.Success(comment))
}
func CommentListByPage(c *gin.Context) {
	var listComment dto.ListComment
	if err := c.ShouldBindJSON(&listComment); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}
	if listComment.PageSize == 0 {
		listComment.PageSize = 10
	}
	if listComment.PageNum == 0 {
		listComment.PageNum = 1
	}
	var totalCount int64
	var commentList []entity.Comment

	baseQuery := config.DB.Model(&entity.Comment{})
	if listComment.PostID != 0 {
		baseQuery.
			Where("post_id = ?", listComment.PostID)
	}
	baseQuery.Count(&totalCount)
	baseQuery.Limit(listComment.PageSize).Offset((listComment.PageNum - 1) * listComment.PageSize).Find(&commentList)
	c.JSON(http.StatusOK, vo.Success(
		vo.PageSuccess(commentList, listComment.PageNum, listComment.PageSize, int(totalCount)),
	))
}
