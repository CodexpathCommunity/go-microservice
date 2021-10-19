package service

import "fmt"

func (b *service) DeletePost(id string) error {
	if err := b.repository.Delete(id); err != nil {
		return err
	}

	// delete data if it exists in redis db
	_ = b.clients.Redis.Delete(fmt.Sprintf("post::%s", id))

	return nil
}
