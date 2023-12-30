package routes

import (
	"github.com/forumGamers/nine-tails-fox/controllers"
	"github.com/gin-gonic/gin"
)

func (r *routes) postRoutes(rg *gin.RouterGroup, postController controllers.PostController) {
	uri := rg.Group("/post")

	uri.GET("/public", postController.GetPublicContent).
		GET("/tags", postController.GetTopTags).
		GET("/liked", postController.GetLikedPost).
		GET("/liked/:userId", postController.GetUserLikedPost).
		GET("/me", postController.GetUserPost).
		GET("/me/:userId", postController.GetPostByUserId).
		GET("/media", postController.GetUserMedia).
		GET("/media/:userId", postController.GetMediaByUserId)
}
