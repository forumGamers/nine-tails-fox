package routes

import (
	"os"

	"github.com/forumGamers/nine-tails-fox/controllers"
	md "github.com/forumGamers/nine-tails-fox/middlewares"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRouters(writer web.ResponseWriter, postController controllers.PostController) {
	r := routes{gin.Default()}

	groupRoutes := r.router.Group("/api/v1")

	mds := md.NewMiddlewares(writer)

	r.router.Use(mds.Authentication)
	r.postRoutes(groupRoutes, postController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4301"
	}

	r.router.Run(":" + port)
}
