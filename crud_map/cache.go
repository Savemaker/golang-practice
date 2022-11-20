package main

type Cache interface {
	Insert(key, value string) (string, error)
	Get(key string) (string, error)
	Update(key, value string) (string, error)
	Delete(key string)
}
