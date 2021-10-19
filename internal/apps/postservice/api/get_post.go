package api

import (
	"context"
	"errors"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service/repository"
	"github.com/twitchtv/twirp"
)

func (h *handler) GetPost(ctx context.Context, req *rpc.GetPostRequest) (*rpc.GetPostResponse, error) {
	if req.Id == "" {
		return nil, twirp.RequiredArgumentError("id")
	}

	postItem, err := h.service.GetPost(req.Id)
	if err != nil {
		if errors.Is(err, repository.ErrPostDoesNotExist) {
			return nil, twirp.NotFoundError(err.Error())
		}

		return nil, twirp.InternalError(err.Error())
	}

	return &rpc.GetPostResponse{
		Post: postItem,
	}, nil
}
