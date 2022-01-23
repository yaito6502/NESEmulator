package pkg

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Btouint8(b bool) uint8 {
	return uint8(btoi(b))
}

func Btouint16(b bool) uint16 {
	return uint16(btoi(b))
}

//対称性を保つために実装
func Uint8tob(i uint8) bool {
	return i != 0
}

func Uint16tob(i uint16) bool {
	return i != 0
}
