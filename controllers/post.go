package controllers

import (
	"context"

	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/pkg/user"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
)

func NewPostController(postRepo post.PostRepo, w web.ResponseWriter, r web.RequestReader) PostController {
	return &PostControllerImpl{w, r, postRepo}
}

func (pc *PostControllerImpl) GetPublicContent(c *gin.Context) {
	uuid := user.GetUser(c).UUID
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetPublicContent(context.Background(), uuid, query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	if len(datas) < 1 {
		pc.AbortHttp(c, pc.New404Error("data not found"))
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}
