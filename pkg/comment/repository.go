package comment

import (
	"context"

	"github.com/forumGamers/nine-tails-fox/errors"
	h "github.com/forumGamers/nine-tails-fox/helpers"
	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/utils"
	"github.com/forumGamers/nine-tails-fox/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCommentRepo(qu utils.QueryUtils) CommentRepo {
	return &CommentRepoImpl{b.NewBaseRepo(b.GetCollection(b.Comment)), qu}
}

func (r *CommentRepoImpl) FindPostComment(ctx context.Context, postId primitive.ObjectID, query web.GetPostParams) ([]CommentResponse, error) {
	curr, err := r.Aggregations(ctx, bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "postId", Value: postId}}}},
		bson.D{
			{Key: "$facet",
				Value: bson.D{
					{Key: "total",
						Value: bson.A{
							bson.D{{Key: "$count", Value: "total"}},
						},
					},
					{Key: "data",
						Value: bson.A{
							r.NewSkip((query.Page - 1) * query.Limit),
							r.NewLimit(query.Limit),
							bson.D{{Key: "$sort", Value: bson.D{{Key: "createdAt", Value: -1}}}},
						},
					},
				},
			},
		},
		r.NewRawUnwind("$total"),
		r.NewRawUnwind("$data"),
		bson.D{
			{Key: "$project",
				Value: bson.D{
					{Key: "_id", Value: "$data._id"},
					{Key: "text", Value: "$data.text"},
					{Key: "postId", Value: "$data.postId"},
					{Key: "userId", Value: "$data.userId"},
					{Key: "createdAt", Value: "$data.createdAt"},
					{Key: "updatedAt", Value: "$data.updatedAt"},
					{Key: "reply", Value: "$data.reply"},
					{Key: "totalData", Value: "$total.total"},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())

	var datas []CommentResponse
	for curr.Next(context.TODO()) {
		var data CommentResponse
		if err := curr.Decode(&data); err != nil {
			return datas, err
		}

		if len(data.Reply) > 0 {
			for i := 0; i < len(data.Reply); i++ {
				data.Reply[i].Text = h.Decryption(data.Reply[i].Text)
			}
		}
		datas = append(datas, data)
	}

	if len(datas) < 1 {
		return datas, errors.NewError("data not found", 404)
	}
	return datas, nil
}
