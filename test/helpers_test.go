package test

import (
	"errors"
	"fmt"
	customerrors "gempellm/playground/internal/custom_errors"
	cache "gempellm/playground/internal/mcache"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func emulateLoad(t *testing.T, c cache.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)

		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		wg.Add(1)

		go func(k, v string) {
			storedValue, err := c.Get(k)
			if !errors.Is(err, customerrors.ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}

			wg.Done()
		}(key, value)

		wg.Add(1)

		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		wg.Wait()
	}
}
