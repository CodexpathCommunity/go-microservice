package api

import (
	"context"
	"errors"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service/repository"
	"github.com/twitchtv/twirp"
)

func (h *handler) UpdatePost(ctx context.Context, req *rpc.UpdatePostRequest) (*rpc.UpdatePostResponse, error) {
	if req.Id == "" {
		return nil, twirp.RequiredArgumentError("id")
	}

	if req.Title == "" {
		return nil, twirp.RequiredArgumentError("title")
	}

	if req.Content == "" {
		return nil, twirp.RequiredArgumentError("content")
	}

	updatedPost, err := h.service.UpdatePost(req.Id, req.Title, req.Content)
	if err != nil {
		if errors.Is(err, repository.ErrPostDoesNotExist) {
			return nil, twirp.NotFoundError(err.Error())
		}

		return nil, twirp.InternalError(err.Error())
	}

	return &rpc.UpdatePostResponse{
		Post: updatedPost,
	}, nil
}
