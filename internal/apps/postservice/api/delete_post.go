package api

import (
	"context"
	"errors"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service/repository"
	"github.com/twitchtv/twirp"
)

func (h *handler) DeletePost(ctx context.Context, req *rpc.DeletePostRequest) (*rpc.DeletePostResponse, error) {
	if req.Id == "" {
		return nil, twirp.RequiredArgumentError("id")
	}

	err := h.service.DeletePost(req.Id)
	if err != nil {
		if errors.Is(err, repository.ErrPostDoesNotExist) {
			return nil, twirp.NotFoundError(err.Error())
		}

		return nil, twirp.InternalError(err.Error())
	}

	return &rpc.DeletePostResponse{
		Success: true,
	}, nil
}
