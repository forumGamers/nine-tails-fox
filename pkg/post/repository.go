package post

import (
	"context"
	"time"

	"github.com/forumGamers/nine-tails-fox/errors"
	h "github.com/forumGamers/nine-tails-fox/helpers"
	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/web"
	"go.mongodb.org/mongo-driver/bson"
)

func NewPostRepo() PostRepo {
	return &PostRepoImpl{b.NewBaseRepo(b.GetCollection(b.Post))}
}

func (r *PostRepoImpl) GetPublicContent(ctx context.Context, userId string, query web.GetPostParams) ([]PostResponse, error) {
	now := time.Now().UTC()
	curr, err := r.BaseRepo.Aggregations(ctx, bson.A{
		bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "createdAt",
						Value: bson.D{
							{Key: "$gte", Value: now.AddDate(0, -1, 0)},
							{Key: "$lte", Value: now},
						},
					},
				},
			},
		},
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
										{Key: "from", Value: "comment"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "comment"},
									},
								},
							},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "like"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "like"},
									},
								},
							},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "share"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "share"},
									},
								},
							},
							bson.D{
								{Key: "$addFields",
									Value: bson.D{
										{Key: "countLike", Value: bson.D{{Key: "$size", Value: "$like"}}},
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
										{Key: "isLiked",
											Value: bson.D{
												{Key: "$reduce",
													Value: bson.D{
														{Key: "input", Value: "$like"},
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
					{Key: "_id", Value: "$datas._id"},
					{Key: "userId", Value: "$datas.userId"},
					{Key: "text", Value: "$datas.text"},
					{Key: "media", Value: "$datas.media"},
					{Key: "allowComment", Value: "$datas.allowComment"},
					{Key: "createdAt", Value: "$datas.createdAt"},
					{Key: "updatedAt", Value: "$datas.updatedAt"},
					{Key: "countLike", Value: "$datas.countLike"},
					{Key: "countComment", Value: "$datas.countComment"},
					{Key: "countShare", Value: "$datas.countShare"},
					{Key: "isLiked", Value: "$datas.isLiked"},
					{Key: "isShared", Value: "$datas.isShared"},
					{Key: "tags", Value: "$datas.tags"},
					{Key: "privacy", Value: "$datas.privacy"},
					{Key: "totalData", Value: "$total.total"},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	var datas []PostResponse
	for curr.Next(context.TODO()) {
		var data PostResponse
		if err := curr.Decode(&data); err != nil {
			return nil, err
		}
		data.Text = h.Decryption(data.Text)
		datas = append(datas, data)
	}

	if len(datas) < 1 {
		return datas, errors.NewError("data not found", 404)
	}

	return datas, nil
}

func (r *PostRepoImpl) GetUserPost(ctx context.Context, userId string, query web.GetPostParams) ([]PostResponse, error) {
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
										{Key: "from", Value: "comment"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "comment"},
									},
								},
							},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "like"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "like"},
									},
								},
							},
							bson.D{
								{Key: "$lookup",
									Value: bson.D{
										{Key: "from", Value: "share"},
										{Key: "localField", Value: "_id"},
										{Key: "foreignField", Value: "postId"},
										{Key: "as", Value: "share"},
									},
								},
							},
							bson.D{
								{Key: "$addFields",
									Value: bson.D{
										{Key: "countLike", Value: bson.D{{Key: "$size", Value: "$like"}}},
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
										{Key: "isLiked",
											Value: bson.D{
												{Key: "$reduce",
													Value: bson.D{
														{Key: "input", Value: "$like"},
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
					{Key: "_id", Value: "$datas._id"},
					{Key: "userId", Value: "$datas.userId"},
					{Key: "text", Value: "$datas.text"},
					{Key: "media", Value: "$datas.media"},
					{Key: "allowComment", Value: "$datas.allowComment"},
					{Key: "createdAt", Value: "$datas.createdAt"},
					{Key: "updatedAt", Value: "$datas.updatedAt"},
					{Key: "countLike", Value: "$datas.countLike"},
					{Key: "countComment", Value: "$datas.countComment"},
					{Key: "countShare", Value: "$datas.countShare"},
					{Key: "isLiked", Value: "$datas.isLiked"},
					{Key: "isShared", Value: "$datas.isShared"},
					{Key: "tags", Value: "$datas.tags"},
					{Key: "privacy", Value: "$datas.privacy"},
					{Key: "totalData", Value: "$total.total"},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	var datas []PostResponse
	for curr.Next(context.TODO()) {
		var data PostResponse
		if err := curr.Decode(&data); err != nil {
			return nil, err
		}

		data.Text = h.Decryption(data.Text)
		datas = append(datas, data)
	}

	if len(datas) < 1 {
		return datas, errors.NewError("data not found", 404)
	}

	return datas, nil
}
