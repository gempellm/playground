package test

import (
	cache "gempellm/playground/internal/mcache"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cache(t *testing.T) {
	t.Parallel()

	testCache := cache.NewMCache()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)

		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storedValue)
	})

	t.Run("correctly updated value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)

		newValue := "someValue#2"
		err = testCache.Set(key, newValue)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)
		newStoredValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, newValue, newStoredValue)

		anotherValue := "someValue#3"
		testCache.Storage[key] = anotherValue
		newAnotherValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, anotherValue, newAnotherValue)

	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoad(t, testCache, parallelFactor)
	})
}
