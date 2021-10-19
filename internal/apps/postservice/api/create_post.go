package api

import (
	"context"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/twitchtv/twirp"
)

func (h *handler) CreatePost(ctx context.Context, req *rpc.CreatePostRequest) (*rpc.CreatePostResponse, error) {
	if req.AuthorId == "" {
		return nil, twirp.RequiredArgumentError("author_id")
	}

	if req.Title == "" {
		return nil, twirp.RequiredArgumentError("title")
	}

	if req.Content == "" {
		return nil, twirp.RequiredArgumentError("content")
	}

	createdPost, err := h.service.CreatePost(req.AuthorId, req.Title, req.Content)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	return &rpc.CreatePostResponse{
		Post: createdPost,
	}, nil
}
