package hash_table_group

import (
	"math/rand"
	"strconv"
	"testing"

	"CustomHashMap/internal/hash_table"
)

// randomKey generator for tests
func randomKey() string {
	return "group_key_" + strconv.Itoa(rand.Intn(1000000))
}

// BenchmarkPutTable bench for tables inserting
func BenchmarkPutTable(b *testing.B) {
	htg := NewHashTableGroupRMBasic(101)
	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = randomKey()
	}

	b.ResetTimer() // После setup.
	for i := 0; i < b.N; i++ {
		dummyTable := &hash_table.HashTableRMBasic{}
		htg.PutTable(keys[i], dummyTable)
	}
}

// BenchmarkGetTable bench for founds
func BenchmarkGetTable(b *testing.B) {
	htg := NewHashTableGroupRMBasic(101)
	numItems := 1000
	keys := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		key := randomKey()
		keys[i] = key
		dummyTable := &hash_table.HashTableRMBasic{}
		htg.PutTable(key, dummyTable)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = htg.GetTable(keys[i%numItems])
	}
}

// BenchmarkGetTableMiss bench for miss founds
func BenchmarkGetTableMiss(b *testing.B) {
	htg := NewHashTableGroupRMBasic(101)
	numItems := 1000
	for i := 0; i < numItems; i++ {
		dummyTable := &hash_table.HashTableRMBasic{}
		htg.PutTable(randomKey(), dummyTable)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = htg.GetTable("missing_group_" + strconv.Itoa(i)) // Miss.
	}
}
