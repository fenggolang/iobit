package main

import (
	"fmt"
)

func i64tob(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}

func btoi64(val []byte) uint64 {
	r := uint64(0)
	for i := uint64(0); i < 8; i++ {
		r |= uint64(val[i]) << (8 * i)
	}
	return r
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func btoi32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}

func main() {
	//var a uint64 = 0x0102030405060708
	//fmt.Println("Uint64:", a, "Bytes array:", i64tob(a))
	//fmt.Println("Uint64:", btoi64(i64tob(a)), "Bytes array:", i64tob(btoi64(i64tob(a))))
	//var b uint32 = 0x01020304

	var b uint32 = 118440758
	fmt.Println("Uint32:", b, "Bytes array:", i32tob(b))
	//fmt.Println("Uint32:", btoi32(i32tob(b)), "Bytes array:", i32tob(btoi32(i32tob(b))))
}