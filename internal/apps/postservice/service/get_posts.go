package service

import (
	"time"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *service) GetPosts(cursor time.Time, limit int32) ([]*rpc.Post, error) {
	posts, err := b.repository.List(cursor, limit)
	if err != nil {
		return nil, err
	}

	var postMessages []*rpc.Post
	for _, p := range posts {
		postMessages = append(postMessages, &rpc.Post{
			Id:        p.ID,
			AuthorId:  p.AuthorID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		})
	}

	return postMessages, nil
}
