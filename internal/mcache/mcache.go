package mcache

import (
	customerrors "gempellm/playground/internal/custom_errors"
	"sync"
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

type MCache struct {
	Storage      map[string]string
	StorageMutex *sync.Mutex
}

func NewMCache() Cache {
	return &MCache{
		Storage:      make(map[string]string),
		StorageMutex: &sync.Mutex{},
	}
}

func (c *MCache) Set(key, value string) error {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// }
	c.StorageMutex.Lock()
	defer c.StorageMutex.Unlock()
	c.Storage[key] = value

	return nil
}

func (c *MCache) Get(key string) (string, error) {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// }

	c.StorageMutex.Lock()
	value, ok := c.Storage[key]
	c.StorageMutex.Unlock()
	if ok {
		return value, nil
	} else {
		return "", customerrors.ErrNotFound
	}
}

func (c *MCache) Delete(key string) error {
	// if c.Storage == nil {
	// 	c.Storage = make(map[string]string)
	// 	return ErrNotFound
	// }

	c.StorageMutex.Lock()
	delete(c.Storage, key)
	c.StorageMutex.Unlock()
	return nil
}
