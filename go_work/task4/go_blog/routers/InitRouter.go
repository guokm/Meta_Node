package routers

import (
	"github/metanode/go_blog/controller"
	"github/metanode/go_blog/middleware"

	"github.com/gin-gonic/gin"
)

var (
	publicGroup  *gin.RouterGroup
	privateGroup *gin.RouterGroup
)

func InitRouter(router *gin.Engine) {
	apiVersion := router.Group("/api/v1")
	{
		// 公开路由组
		publicGroup = apiVersion.Group("")

		// 受保护路由组（需要JWT）
		privateGroup = apiVersion.Group("", middleware.JWTAuth())
	}

	RegisterUserRoutes()
	RegisterPostRoutes()
	RegisterCommentRoutes()
}

func RegisterUserRoutes() {
	publicGroup.Group("/user")
	{
		publicGroup.POST("/register", controller.Register)
		publicGroup.POST("/login", controller.Login)
	}

}

func RegisterPostRoutes() {
	privateGroup.Group("/post")
	{
		privateGroup.POST("/creatPost", controller.CreatPost)
		privateGroup.POST("/postListByPage", controller.PostListByPage)
		privateGroup.GET("/posts/:id", controller.FindPostById)
		privateGroup.POST("/updatePost", controller.UpdatePost)
		privateGroup.POST("/deletePost", controller.DeletePost)
	}
}

func RegisterCommentRoutes() {
	publicGroup.Group("/comment")
	{
		publicGroup.POST("/commentListByPage", controller.CommentListByPage)
	}

	privateGroup.Group("/comment")
	{
		privateGroup.POST("/creatComment", controller.CreatComment)
	}
}
