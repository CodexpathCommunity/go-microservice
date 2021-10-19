package api

import (
	"net/http"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service"
	"github.com/go-chi/chi/v5"
	"github.com/twitchtv/twirp"
)

// NewServer returns a new post service http multiplexer.
func NewServer(b service.Service) http.Handler {
	mux := chi.NewMux()
	handler := rpc.NewPostsServer(NewHandler(b), twirp.WithServerPathPrefix("/v1"))

	mux.Mount(handler.PathPrefix(), handler)

	return mux
}

type handler struct {
	service service.Service
}

// NewHandler returns a new handler with the Posts interface.
func NewHandler(service service.Service) rpc.Posts {
	return &handler{
		service: service,
	}
}
