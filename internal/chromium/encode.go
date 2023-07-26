package chromium

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
)

var u16null = []byte{0x00, 0x00}

func EncodeWebCustomMIMEData(m map[string]string) (_ []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("write web custom mime data: %v", r)
		}
	}()

	payload := make([][]byte, 1+len(m)*4)

	// entries
	payload[0] = writeInt(len(m))

	// map entries
	i := 1
	for k, v := range m {
		// write key
		u16key := utf16.Encode([]rune(k))
		payload[i] = writeInt(len(u16key))
		payload[i+1] = writeStr(u16key)
		// write value
		u16val := utf16.Encode([]rune(v))
		payload[i+2] = writeInt(len(u16val))
		payload[i+3] = writeStr(u16val)

	}

	out := bytes.Join(payload, u16null)
	msgsize := writeInt(len(out))
	return bytes.Join([][]byte{msgsize, u16null, out}, nil), nil

}

func writeInt(n int) []byte {
	dst := make([]byte, 2)
	binary.LittleEndian.PutUint16(dst, uint16(n))
	return dst
}

func writeStr(u16str []uint16) []byte {
	dst := make([]byte, len(u16str)*2)
	for i, n := range u16str {
		binary.LittleEndian.PutUint16(dst[i*2:i*2+2], n)
	}
	return dst
}
