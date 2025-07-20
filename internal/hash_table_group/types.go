package hash_table_group

import "CustomHashMap/internal/hash_table"

type HashTableGroupRMBasic struct {
	tableBuckets [][]pair
	size         uint64
}

type pair struct {
	key   string
	table hash_table.HashTable
}

func NewHashTableGroupRMBasic(size uint64) *HashTableGroupRMBasic {
	return &HashTableGroupRMBasic{
		tableBuckets: make([][]pair, size),
		size:         size,
	}
}
