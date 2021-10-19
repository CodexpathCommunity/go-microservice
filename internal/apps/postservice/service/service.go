package service

import (
	"time"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/clients"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service/repository"
)

// Service methods that contains post use cases.
type Service interface {
	GetPosts(cursor time.Time, limit int32) ([]*rpc.Post, error)
	GetPost(id string) (*rpc.Post, error)
	CreatePost(authorID, title, content string) (*rpc.Post, error)
	UpdatePost(id, title, content string) (*rpc.Post, error)
	DeletePost(id string) error
}

type service struct {
	clients    clients.Clients
	repository repository.PostRepository
}

// Validate service interface implementation.
var _ Service = &service{}

// New creates a new service service.
func New(clients clients.Clients) Service {
	return &service{
		clients:    clients,
		repository: repository.NewPostRepository(clients.Database),
	}
}
