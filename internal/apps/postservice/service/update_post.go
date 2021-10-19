package service

import (
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *service) UpdatePost(id, title, content string) (*rpc.Post, error) {
	updatedPost, err := b.repository.Update(id, title, content)
	if err != nil {
		return nil, err
	}

	return &rpc.Post{
		Id:        updatedPost.ID,
		AuthorId:  updatedPost.AuthorID,
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		CreatedAt: timestamppb.New(updatedPost.CreatedAt),
		UpdatedAt: timestamppb.New(updatedPost.UpdatedAt),
	}, nil
}
