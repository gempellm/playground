package bench

import (
	"errors"
	"fmt"
	customerrors "gempellm/playground/internal/custom_errors"
	"sync"
	"testing"
)

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

func emulateLoad(b *testing.B, c Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)

		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)

		go func(k, v string) {
			_, err := c.Get(k)
			if err != nil {
				if !errors.Is(err, customerrors.ErrNotFound) {
					panic(err)
				}
			}

			wg.Done()
		}(key, value)

		wg.Add(1)

		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

	}
	wg.Wait()
}

func emulateLoadIntenseRead(b *testing.B, c Cache, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor/100; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)

		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

		wg.Add(1)

		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)

	}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)

		go func(k, v string) {
			_, err := c.Get(k)
			if err != nil {
				if !errors.Is(err, customerrors.ErrNotFound) {
					panic(err)
				}
			}

			wg.Done()
		}(key, value)
	}
	wg.Wait()
}
