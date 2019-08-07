package conversion

import (
	"fmt"
	"reflect"
	"unsafe"
)

// byte to string
func ByteToString(b []byte) string {
	str := ""
	for i := 0; i < len(b); i++ {
		str += fmt.Sprintf("%02x", b[i])
	}
	return str
}

func BytesToStringFast(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
