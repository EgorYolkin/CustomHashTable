package hash_table

type HashTableRMBasic struct {
	buckets [][]pair
	size    uint64
}

type pair struct {
	key   string
	value interface{}
}

func NewHashTableRMBasic(size uint64) *HashTableRMBasic {
	return &HashTableRMBasic{
		buckets: make([][]pair, size),
		size:    size,
	}
}
