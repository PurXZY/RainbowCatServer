package util

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	_ = binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}