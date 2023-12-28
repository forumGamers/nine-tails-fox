package post

import (
	"context"

	b "github.com/forumGamers/nine-tails-fox/pkg/base"
	"github.com/forumGamers/nine-tails-fox/web"
)

type PostRepo interface {
	GetPublicContent(ctx context.Context, userId string, query web.GetPostParams) ([]PostResponse, error)
	GetUserPost(ctx context.Context, userId string, query web.GetPostParams) ([]PostResponse, error)
	GetUserPostMedia(ctx context.Context, userId string, query web.GetPostParams) ([]PostResponse, error)
	GetTopTags(ctx context.Context, query web.GetPostParams) ([]TopTags, error)
}

type PostRepoImpl struct {
	b.BaseRepo
}
