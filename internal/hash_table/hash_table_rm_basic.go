package hash_table

import "CustomHashMap/pkg/hash"

func (ht *HashTableRMBasic) Put(
	key string,
	value interface{},
) {
	hashV := hash.HashRMBasic(key) % ht.size
	bucket := &ht.buckets[hashV]

	for i, p := range *bucket {
		if p.key == key {
			(*bucket)[i].value = value
			return
		}
	}

	*bucket = append(*bucket, pair{key, value})
}

func (ht *HashTableRMBasic) Get(key string) (interface{}, bool) {
	hashV := hash.HashRMBasic(key) % ht.size
	bucket := ht.buckets[hashV]

	for _, p := range bucket {
		if p.key == key {
			return p.value, true
		}
	}

	return nil, false
}
