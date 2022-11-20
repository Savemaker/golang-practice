package main

const (
	EMPTY_STRING        = ""
	ErrKeyDoesNotExist  = CacheError("key does not exist")
	ErrKeyAlreadyExists = CacheError("key already exists")
)

type InMemmoryCache map[string]string

type CacheError string

func (c CacheError) Error() string {
	return string(c)
}

func (i InMemmoryCache) Insert(key, value string) (string, error) {
	_, exists := i[key]

	if exists {
		return EMPTY_STRING, ErrKeyAlreadyExists
	} else {
		i[key] = value
		return i[key], nil
	}
}

func (i InMemmoryCache) Get(key string) (string, error) {
	val, exists := i[key]

	if exists {
		return val, nil
	} else {
		return EMPTY_STRING, ErrKeyDoesNotExist
	}
}

func (i InMemmoryCache) Update(key, value string) (string, error) {
	_, exists := i[key]

	if exists {
		i[key] = value
		return i[key], nil
	} else {
		return EMPTY_STRING, ErrKeyDoesNotExist
	}
}

func (i InMemmoryCache) Delete(key string) {
	delete(i, key)
}
