package like

import (
	"context"

	"github.com/forumGamers/nine-tails-fox/errors"
	h "github.com/forumGamers/nine-tails-fox/helpers"
	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewLikeRepo() LikeRepo {
	return &LikeRepoImpl{b.NewBaseRepo(b.GetCollection(b.Like))}
}

func (r *LikeRepoImpl) FindUserLikedPost(ctx context.Context, userId string, query web.GetPostParams) ([]post.PostResponse, error) {
	curr, err := r.Aggregations(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "userId", Value: userId}}}},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "createdAt", Value: -1}}}},
		bson.D{
			{Key: "$facet",
				Value: bson.D{
					{Key: "total",
						Value: bson.A{
							bson.D{{Key: "$count", Value: "total"}},
						},
					},
					{Key: "datas",
						Value: bson.A{
							bson.D{{Key: "$skip", Value: (query.Page - 1) * query.Limit}},
							bson.D{{Key: "$limit", Value: query.Limit}},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "post"},
										{Key: "localField", Value: "postId"},
										{Key: "foreignField", Value: "_id"},
										{Key: "as", Value: "post"},
									},
								},
							},
							bson.D{{Key: "$unwind", Value: "$post"}},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "comment"},
										{Key: "localField", Value: "post._id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "comment"},
									},
								},
							},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "share"},
										{Key: "localField", Value: "post._id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "share"},
									},
								},
							},
							bson.D{
								{Key: "$addFields",
									Value: bson.D{
										{Key: "countShare", Value: bson.D{{Key: "$size", Value: "$share"}}},
										{Key: "countComment",
											Value: bson.D{
												{Key: "$sum",
													Value: bson.A{
														bson.D{{Key: "$size", Value: "$comment"}},
														bson.D{{Key: "$size", Value: "$comment.reply"}},
													},
												},
											},
										},
										{Key: "isShared",
											Value: bson.D{
												{Key: "$reduce",
													Value: bson.D{
														{Key: "input", Value: "$share"},
														{Key: "initialValue", Value: false},
														{Key: "in",
															Value: bson.D{
																{Key: "$cond",
																	Value: bson.A{
																		bson.D{
																			{Key: "$eq",
																				Value: bson.A{
																					"$$this.userId",
																					userId,
																				},
																			},
																		},
																		true,
																		"$$value",
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$datas"}},
		bson.D{{Key: "$unwind", Value: "$total"}},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "post", Value: "$datas.post"},
					{Key: "countComment", Value: "$datas.countComment"},
					{Key: "countShare", Value: "$datas.countShare"},
					{Key: "isShared", Value: "$datas.isShared"},
					{Key: "total", Value: "$total.total"},
				},
			},
		},
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: "$post._id"},
					{Key: "userId", Value: "$post.userId"},
					{Key: "text", Value: "$post.text"},
					{Key: "media", Value: "$post.media"},
					{Key: "allowComment", Value: "$post.allowComment"},
					{Key: "createdAt", Value: "$post.createdAt"},
					{Key: "updatedAt", Value: "$post.updatedAt"},
					{Key: "countComment", Value: 1},
					{Key: "isShared", Value: 1},
					{Key: "tags", Value: "$post.tags"},
					{Key: "privacy", Value: "$post.privacy"},
					{Key: "totalData", Value: "$total"},
					{Key: "countShare", Value: 1},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	var datas []post.PostResponse
	for curr.Next(context.TODO()) {
		var data post.PostResponse
		if err := curr.Decode(&data); err != nil {
			return datas, err
		}

		data.IsLiked = true
		data.Text = h.Decryption(data.Text)

		datas = append(datas, data)
	}

	if len(datas) < 1 {
		return datas, errors.NewError("data not found", 404)
	}

	return datas, nil
}

func (r *LikeRepoImpl) CountPostLikes(ctx context.Context, ids []primitive.ObjectID) ([]PostLikes, error) {
	curr, err := r.Aggregations(ctx, bson.A{
		bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "postId",
						Value: bson.D{
							{Key: "$in", Value: ids},
						},
					},
				},
			},
		},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$postId"},
					{Key: "totalLike", Value: bson.D{{Key: "$sum", Value: 1}}},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	var datas []PostLikes
	for curr.Next(context.TODO()) {
		var data PostLikes
		if err := curr.Decode(&data); err != nil {
			return datas, err
		}

		datas = append(datas, data)
	}

	if len(datas) < 1 {
		return datas, errors.NewError("data not found", 404)
	}

	return datas, nil
}
