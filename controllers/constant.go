package controllers

import (
	"github.com/forumGamers/nine-tails-fox/pkg/comment"
	"github.com/forumGamers/nine-tails-fox/pkg/like"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	GetPublicContent(c *gin.Context)
	GetLikedPost(c *gin.Context)
	GetUserPost(c *gin.Context)
	GetUserMedia(c *gin.Context)
	GetPostByUserId(c *gin.Context)
	GetMediaByUserId(c *gin.Context)
	GetUserLikedPost(c *gin.Context)
	GetTopTags(c *gin.Context)
}

type PostControllerImpl struct {
	web.ResponseWriter
	web.RequestReader
	PostRepo post.PostRepo
	LikeRepo like.LikeRepo
}

type CommentController interface {
	FindPostComment(c *gin.Context)
}

type CommentControllerImpl struct {
	CommentRepo comment.CommentRepo
	web.ResponseWriter
	web.RequestReader
}
