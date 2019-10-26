package util

import (
	"encoding/binary"
)

func Uint32ToBytes(num uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, num)
	return buf
}

func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}