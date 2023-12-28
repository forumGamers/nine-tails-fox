package controllers

import (
	"context"

	"github.com/forumGamers/nine-tails-fox/pkg/like"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/pkg/user"
	"github.com/forumGamers/nine-tails-fox/web"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewPostController(postRepo post.PostRepo, w web.ResponseWriter, r web.RequestReader, l like.LikeRepo) PostController {
	return &PostControllerImpl{w, r, postRepo, l}
}

func (pc *PostControllerImpl) GetPublicContent(c *gin.Context) {
	uuid := user.GetUser(c).UUID
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)
	//get $or userId yang di follow + $or tags teratas dalam 1 hari
	datas, err := pc.PostRepo.GetPublicContent(context.Background(), uuid, query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetUserPost(c *gin.Context) {
	uuid := user.GetUser(c).UUID
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetUserPost(context.Background(), uuid, query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetLikedPost(c *gin.Context) {
	uuid := user.GetUser(c).UUID
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.LikeRepo.FindUserLikedPost(context.Background(), uuid, query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	var postIds []primitive.ObjectID
	for _, post := range datas {
		postIds = append(postIds, post.Id)
	}

	metas, err := pc.LikeRepo.CountPostLikes(context.Background(), postIds)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	for _, meta := range metas {
		for i := 0; i < len(datas); i++ {
			if meta.Id == datas[i].Id {
				datas[i].CountLike = meta.TotalLike
			}
		}
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetUserMedia(c *gin.Context) {
	uuid := user.GetUser(c).UUID
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetUserPostMedia(context.Background(), uuid, query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetPostByUserId(c *gin.Context) {
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetUserPost(context.Background(), c.Param("userId"), query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetMediaByUserId(c *gin.Context) {
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetUserPostMedia(context.Background(), c.Param("userId"), query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetUserLikedPost(c *gin.Context) {
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.LikeRepo.FindUserLikedPost(context.Background(), c.Param("userId"), query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	var postIds []primitive.ObjectID
	for _, post := range datas {
		postIds = append(postIds, post.Id)
	}

	metas, err := pc.LikeRepo.CountPostLikes(context.Background(), postIds)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	for _, meta := range metas {
		for i := 0; i < len(datas); i++ {
			if meta.Id == datas[i].Id {
				datas[i].CountLike = meta.TotalLike
			}
		}
	}

	pc.Write200ResponseWithMetadata(c, "OK", datas, web.MetaData{
		Total: datas[0].TotalData,
		Page:  query.Page,
		Limit: query.Limit,
	})
}

func (pc *PostControllerImpl) GetTopTags(c *gin.Context) {
	var query web.GetPostParams
	pc.GetParams(c, &query)
	pc.DefaultPage(&query)
	pc.DefaultLimit(&query)

	datas, err := pc.PostRepo.GetTopTags(context.Background(), query)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write200Response(c, "OK", datas)
}
