package service

import (
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *service) CreatePost(authorID, title, content string) (*rpc.Post, error) {
	newPost, err := b.repository.Create(authorID, title, content)
	if err != nil {
		return nil, err
	}

	return &rpc.Post{
		Id:        newPost.ID,
		AuthorId:  newPost.AuthorID,
		Title:     newPost.Title,
		Content:   content,
		CreatedAt: timestamppb.New(newPost.CreatedAt),
		UpdatedAt: timestamppb.New(newPost.UpdatedAt),
	}, nil
}
