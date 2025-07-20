package hash_table

type HashTable interface {
	Put(key string, value interface{})
	Get(key string) (interface{}, bool)
}
