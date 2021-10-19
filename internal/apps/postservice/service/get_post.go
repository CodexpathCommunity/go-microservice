package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/rpc"
	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/service/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *service) GetPost(id string) (*rpc.Post, error) {
	cachedPost, err := b.clients.Redis.Get(fmt.Sprintf("post::%s", id))
	if err != nil {
		postItem, err := b.repository.Find(id)
		if err != nil {
			return nil, err
		}

		p, err := json.Marshal(postItem)
		if err != nil {
			return nil, err
		}

		err = b.clients.Redis.SetWithTTL(fmt.Sprintf("post::%s", id), p, 60*time.Second)
		if err != nil {
			return nil, err
		}

		return &rpc.Post{
			Id:        postItem.ID,
			AuthorId:  postItem.AuthorID,
			Title:     postItem.Title,
			Content:   postItem.Content,
			CreatedAt: timestamppb.New(postItem.CreatedAt),
			UpdatedAt: timestamppb.New(postItem.UpdatedAt),
		}, nil
	}

	var postItem repository.Post
	if err := json.Unmarshal([]byte(cachedPost), &postItem); err != nil {
		return nil, err
	}

	return &rpc.Post{
		Id:        postItem.ID,
		AuthorId:  postItem.AuthorID,
		Title:     postItem.Title,
		Content:   postItem.Content,
		CreatedAt: timestamppb.New(postItem.CreatedAt),
		UpdatedAt: timestamppb.New(postItem.UpdatedAt),
	}, nil
}
