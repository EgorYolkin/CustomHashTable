package hash_table

import (
	"math/rand"
	"strconv"
	"testing"
)

// randomKey generator for tests
func randomKey() string {
	return "key_" + strconv.Itoa(rand.Intn(1000000))
}

// BenchmarkPut bench for inserting
func BenchmarkPut(b *testing.B) {
	ht := NewHashTableRMBasic(101)
	keys := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		keys[i] = randomKey()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ht.Put(keys[i], i)
	}
}

// BenchmarkGet bench for founds
func BenchmarkGet(b *testing.B) {
	ht := NewHashTableRMBasic(101)
	numItems := 1000
	keys := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		key := randomKey()
		keys[i] = key
		ht.Put(key, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ht.Get(keys[i%numItems])
	}
}

// BenchmarkGetMiss bench for miss founds
func BenchmarkGetMiss(b *testing.B) {
	ht := NewHashTableRMBasic(101)
	numItems := 1000
	for i := 0; i < numItems; i++ {
		ht.Put(randomKey(), i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ht.Get("missing_key_" + strconv.Itoa(i)) // Miss.
	}
}
