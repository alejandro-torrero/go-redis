package store

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Keys() []string
	Del(key string)
	Len() int
	Incr(key string) (int, error)
}
