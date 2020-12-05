package convertutil

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytes(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})

	binary.Write(bytesBuffer, binary.BigEndian, n)

	return bytesBuffer.Bytes()
}

func BytesToInt32(val []byte) int32 {
	bytesBuffer := bytes.NewBuffer(val)

	var n int32

	binary.Read(bytesBuffer, binary.BigEndian, &n)

	return n
}
