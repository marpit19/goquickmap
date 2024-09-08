package hash

import "math/bits"

// hash computes a hash value for the given string
func Hash(s string) uint64 {
	var h uint64 = 14695981039346656037 // FNV-1 64 bit
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211 // FNV prime
	}
	return bits.RotateLeft64(h, 13)
}
