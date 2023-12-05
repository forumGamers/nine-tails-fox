package routes

import (
	"github.com/forumGamers/nine-tails-fox/controllers"
	"github.com/gin-gonic/gin"
)

func (r *routes) postRoutes(rg *gin.RouterGroup, postController controllers.PostController) {
	uri := rg.Group("/post")

	uri.GET("/public", postController.GetPublicContent)
}
