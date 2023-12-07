package routes

import (
	"github.com/forumGamers/nine-tails-fox/controllers"
	"github.com/gin-gonic/gin"
)

func (r *routes) commentRoutes(rg *gin.RouterGroup, c controllers.CommentController) {
	uri := rg.Group("/comment")

	uri.GET("/:postId", c.FindPostComment)
}
