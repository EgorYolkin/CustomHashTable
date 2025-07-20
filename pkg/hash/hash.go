package hash

func HashRMBasic(s string) uint64 {
	var hash uint64 = 0xDEADBEEF
	for _, b := range []byte(s) {
		hash = ((hash << 5) + (hash << 2) + hash) ^ uint64(b) + uint64(b)
	}
	hash ^= hash >> 32
	return hash
}
