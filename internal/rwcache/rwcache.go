package rwcache

import (
	customerrors "gempellm/playground/internal/custom_errors"
	"sync"
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type RWCache struct {
	Storage        map[string]string
	StorageRWMutex *sync.RWMutex
}

func NewRWCache() *RWCache {
	return &RWCache{
		Storage:        make(map[string]string),
		StorageRWMutex: &sync.RWMutex{},
	}
}

func (c *RWCache) Set(key, value string) error {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// }
	c.StorageRWMutex.Lock()
	defer c.StorageRWMutex.Unlock()
	c.Storage[key] = value

	return nil
}

func (c *RWCache) Get(key string) (string, error) {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// }

	c.StorageRWMutex.RLock()
	value, ok := c.Storage[key]
	c.StorageRWMutex.RUnlock()
	if ok {
		return value, nil
	} else {
		return "", customerrors.ErrNotFound
	}

}

func (c *RWCache) Delete(key string) error {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// 	return ErrNotFound
	// }

	c.StorageRWMutex.Lock()
	delete(c.Storage, key)
	c.StorageRWMutex.Unlock()
	return nil
}
