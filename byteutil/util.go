package byteutil

// Copy copies a slice to make it immutable
func Copy(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}
