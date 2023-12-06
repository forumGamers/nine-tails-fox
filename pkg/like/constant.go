package like

import (
	"context"

	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/pkg/post"
	"github.com/forumGamers/nine-tails-fox/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeRepo interface {
	FindUserLikedPost(ctx context.Context, userId string, query web.GetPostParams) ([]post.PostResponse, error)
	CountPostLikes(ctx context.Context, ids []primitive.ObjectID) ([]PostLikes, error)
}

type LikeRepoImpl struct {
	b.BaseRepo
}
