package main

import (
	cfg "github.com/forumGamers/nine-tails-fox/config"
	"github.com/forumGamers/nine-tails-fox/controllers"
	"github.com/forumGamers/nine-tails-fox/errors"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	r "github.com/forumGamers/nine-tails-fox/routes"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/joho/godotenv"
)

func main() {
	errors.PanicIfError(godotenv.Load())

	cfg.Connection()
	responseWriter := web.NewResponseWriter()
	requestReader := web.NewRequestReader()

	postRepo := post.NewPostRepo()
	postController := controllers.NewPostController(postRepo, responseWriter, requestReader)

	r.NewRouters(responseWriter, postController)

}
