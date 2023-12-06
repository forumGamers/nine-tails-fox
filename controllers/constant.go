package controllers

import (
	"github.com/forumGamers/nine-tails-fox/pkg/like"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	GetPublicContent(c *gin.Context)
	GetLikedPost(c *gin.Context)
	GetUserPost(c *gin.Context)
}

type PostControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	PostRepo post.PostRepo
	LikeRepo like.LikeRepo
}
