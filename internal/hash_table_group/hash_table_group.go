package hash_table_group

import "CustomHashMap/internal/hash_table"

type HashTableGroup interface {
	PutTable(key string, table hash_table.HashTable)
	GetTable(key string) (interface{}, bool)
}
