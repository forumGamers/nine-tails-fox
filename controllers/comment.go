package controllers

import (
	"context"

	"github.com/forumGamers/nine-tails-fox/pkg/comment"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCommentController(repo comment.CommentRepo, w web.ResponseWriter, r web.RequestReader) CommentController {
	return &CommentControllerImpl{repo, w, r}
}

func (cc *CommentControllerImpl) FindPostComment(c *gin.Context) {
	postId, err := primitive.ObjectIDFromHex(c.Param("postId"))
	if err != nil {
		cc.AbortHttp(c, cc.NewInvalidObjectIdError())
		return
	}

	var query web.GetPostParams
	cc.GetParams(c, &query)
	cc.DefaultPage(&query)
	cc.DefaultLimit(&query)

	datas, err := cc.CommentRepo.FindPostComment(context.Background(), postId, query)
	if err != nil {
		cc.AbortHttp(c, err)
		return
	}

	cc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}
