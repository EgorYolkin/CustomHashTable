package hash_table_group

import (
	"CustomHashMap/internal/hash_table"
	"CustomHashMap/pkg/hash"
)

func (htg *HashTableGroupRMBasic) PutTable(
	key string,
	table hash_table.HashTable,
) {
	hashV := hash.HashRMBasic(key) % htg.size
	tableBucket := &htg.tableBuckets[hashV]

	for i, p := range *tableBucket {
		if p.key == key {
			(*tableBucket)[i].table = table
			return
		}
	}

	*tableBucket = append(*tableBucket, pair{key, table})
}

func (htg *HashTableGroupRMBasic) GetTable(key string) (interface{}, bool) {
	hashV := hash.HashRMBasic(key) % htg.size
	tableBucket := &htg.tableBuckets[hashV]

	for _, p := range *tableBucket {
		if p.key == key {
			return p.table, true
		}
	}

	return nil, false
}
