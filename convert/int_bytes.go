package convertutil

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytes(n int32) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})

	err := binary.Write(bytesBuffer, binary.BigEndian, n)

	if err != nil {
		return nil, err
	}

	return bytesBuffer.Bytes(), nil
}

func BytesToInt32(val []byte) (int32, error) {
	bytesBuffer := bytes.NewBuffer(val)

	var n int32

	err := binary.Read(bytesBuffer, binary.BigEndian, &n)

	if err != nil {
		return -1, err
	}

	return n, nil
}
