package utils

// returns true if given []bool are equals
func EqualBits(a, b []bool) bool {
	// length must match
	if len(a) != len(b) {
		return false
	}

	// bits with same index must match
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func ByteSliceToBoolSlice(byteSlice []byte) []bool {
	boolSlice := make([]bool, 8*len(byteSlice))
	for i, b := range byteSlice {
		for j := 0; j < 8; j++ {
			boolSlice[8*i+j] = (b>>uint(7-j))&1 == 1
		}
	}
	return boolSlice
}
