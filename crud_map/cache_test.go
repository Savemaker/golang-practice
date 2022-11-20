package main

import (
	"strconv"
	"testing"
)

const (
	FIRST_KEY               = "1"
	FIRST_KEY_VALUE         = "Great secret"
	FIRST_KEY_UPDATED_VALUE = "Great exclusive secret"
)

var cache Cache = InMemmoryCache{}

func TestCache(t *testing.T) {

	var assertError = func(t *testing.T, errGot error, errWanted error) {
		if errGot != errWanted {
			t.Fatalf("wanted error %q but got %q", errWanted, errGot)
		}
	}

	var assertNoError = func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatalf("Expected no errors but got %q", err)
		}
	}

	t.Run("test insert to empty cache", func(t *testing.T) {
		got, err := cache.Insert(FIRST_KEY, FIRST_KEY_VALUE)

		assertNoError(t, err)

		if got != FIRST_KEY_VALUE {
			t.Errorf("got %q expected %q", got, FIRST_KEY_VALUE)
		}

	})

	t.Run("test update cache for existing key", func(t *testing.T) {
		got, err := cache.Update(FIRST_KEY, FIRST_KEY_UPDATED_VALUE)

		assertNoError(t, err)

		if got != FIRST_KEY_UPDATED_VALUE {
			t.Fatalf("wanted %q but got %q", FIRST_KEY_UPDATED_VALUE, got)
		}
	})

	t.Run("test get from cache", func(t *testing.T) {
		got, err := cache.Get(FIRST_KEY)

		assertNoError(t, err)

		if got != FIRST_KEY_UPDATED_VALUE {
			t.Fatalf("got %q but wanted %q", got, FIRST_KEY_UPDATED_VALUE)
		}
	})

	t.Run("test delete from cache", func(t *testing.T) {
		cache.Delete(FIRST_KEY)

		_, err := cache.Get(FIRST_KEY)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func BenchmarkMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Insert(strconv.FormatInt(int64(i), 10), "Some string")
	}
}

func BenchmarkMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Get(strconv.FormatInt(int64(i), 10))
	}
}

func BenchmarkMapDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.Delete(strconv.FormatInt(int64(i), 10))
	}
}
