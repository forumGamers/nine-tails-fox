package routes

import (
	"github.com/forumGamers/nine-tails-fox/controllers"
	"github.com/gin-gonic/gin"
)

func (r *routes) postRoutes(rg *gin.RouterGroup, postController controllers.PostController) {
	uri := rg.Group("/post")

	uri.GET("/public", postController.GetPublicContent)
	uri.GET("/liked", postController.GetLikedPost)
	uri.GET("/liked/:userId", postController.GetUserLikedPost)
	uri.GET("/me", postController.GetUserPost)
	uri.GET("/me/:userId", postController.GetPostByUserId)
	uri.GET("/media", postController.GetUserMedia)
	uri.GET("/media/:userId", postController.GetMediaByUserId)
}
