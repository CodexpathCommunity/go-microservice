package kv

import "time"

type KV interface {
	// Get returns the value for the given key.
	Get(key string) (string, error)
	// Set sets the value for the given key.
	Set(key string, value interface{}) error
	// SetWithTTL sets the value for the given key with a TTL.
	SetWithTTL(key string, value interface{}, ttl time.Duration) error
	// Delete deletes the value for the given key.
	Delete(key string) error
}
