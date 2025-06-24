package controller

import (
	"errors"
	"github/metanode/go_blog/config"
	"github/metanode/go_blog/model/dto"
	"github/metanode/go_blog/model/entity"
	"github/metanode/go_blog/model/vo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatPost(c *gin.Context) {
	var addPost dto.AddPostDto
	if err := c.ShouldBindJSON(&addPost); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}

	userId := c.GetUint("ID")
	post := entity.Post{
		Title:   addPost.Title,
		Content: addPost.Content,
		UserID:  userId,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, vo.Success(post))
}

func PostListByPage(c *gin.Context) {
	var findPost dto.FindPostDto
	if err := c.ShouldBindJSON(&findPost); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}
	if findPost.PageSize == 0 {
		findPost.PageSize = 10
	}
	if findPost.PageNum == 0 {
		findPost.PageNum = 1
	}

	var totalCount int64
	var posts []entity.Post

	baseQuery := config.DB.Model(&entity.Post{}).Omit("content")
	if findPost.Title != "" {
		baseQuery = baseQuery.Where("title LIKE ?", "%"+findPost.Title+"%")
	}

	baseQuery.Count(&totalCount)
	baseQuery.
		Order("created_at DESC"). // 按创建时间倒序
		Limit(findPost.PageSize).
		Offset((findPost.PageNum - 1) * findPost.PageSize).
		Find(&posts)

	c.JSON(http.StatusOK,
		vo.PageSuccess(posts, findPost.PageNum, findPost.PageSize, int(totalCount)))
}

func FindPostById(c *gin.Context) {
	id := c.Param("id")

	var post entity.Post
	err := config.DB.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "FindPostById is error "+err.Error()))
		return
	}
	c.JSON(http.StatusOK, vo.Success(post))
}

func UpdatePost(c *gin.Context) {
	var updatePost dto.UpdatePostDto

	if err := c.ShouldBindJSON(&updatePost); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}
	if updatePost.Content == "" && updatePost.Title == "" {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "Failed to update posts"))
		return
	}

	var post entity.Post
	err := config.DB.
		Select("id,user_id").
		Take(&post, updatePost.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "无数据"))
			return
		}
	}
	if post.UserID != c.GetUint("ID") {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "无权限"))
		return
	}

	post.Title = updatePost.Title
	post.Content = updatePost.Content

	err = config.DB.Model(&entity.Post{}).
		Where("id = ?", post.ID).
		Updates(post).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "更新失败"))
		log.Fatalf("failed to find posts: %v", err)
		return
	}

	c.JSON(http.StatusOK, vo.Success(post))
}

func DeletePost(c *gin.Context) {
	var getPost dto.GetOrDeletePostDto
	if err := c.ShouldBindJSON(&getPost); err != nil {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, err.Error()))
		return
	}

	var post entity.Post
	err := config.DB.
		Select("id,user_id").
		Take(&post, getPost.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "无数据"))
			return
		}
	}
	if post.UserID != c.GetUint("ID") {
		c.JSON(http.StatusBadRequest, vo.Error(vo.CodeParamError, "无权限"))
		return
	}

	if err := config.DB.Delete(&entity.Post{}, getPost.ID).Error; err == nil {
		c.JSON(http.StatusOK, vo.Success("删除成功"))
	} else {
		c.JSON(http.StatusOK, vo.Success("删除失败"))
	}
}
