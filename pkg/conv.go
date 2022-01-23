package pkg

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

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

func GetFuncName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func ConvUpperHexString(i uint64) string {
	str := strings.ToUpper(strconv.FormatUint(i, 16))
	if len(str) <= 1 {
		return "0" + str
	}
	return str
}
