package api

import (
	"context"
	"time"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/twitchtv/twirp"
)

func (h *handler) GetPosts(ctx context.Context, req *rpc.GetPostsRequest) (*rpc.GetPostsResponse, error) {
	cursor, err := time.Parse(time.RFC3339, req.Cursor)
	if err != nil {
		return nil, twirp.RequiredArgumentError("curosr")
	}

	posts, err := h.service.GetPosts(cursor, req.Limit)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	return &rpc.GetPostsResponse{
		Posts: posts,
	}, nil
}
