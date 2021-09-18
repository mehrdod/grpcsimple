package cache

type Cache interface {
	Set(key string, value int64) error
	Get(key string) (int64, error)
}
