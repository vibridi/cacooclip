// Copyright 2023 Gabriele V. <gabriele@nulab.com>

package chromium

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"unicode/utf16"
)

type u16string []uint16

func DecodeWebCustomMIMEData(b []byte) (_ map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("read web custom mime data: %v", r)
		}
	}()
	dumpU16(b)
	buf := bytes.NewBuffer(b)
	res := map[string]string{}

	read := reader(buf)

	_ = readNumber(read) // msg size

	entries := readNumber(read)
	for i := 0; i < entries; i++ {
		webmime := readU16String(read)
		content := readU16String(read)
		res[webmime] = content
	}
	return res, nil
}

func reader(buf *bytes.Buffer) func(any) {
	return func(v any) {
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			panic(err)
		}
	}
}

func dumpU16(b []byte) {
	buf := bytes.NewBuffer(b)
	for {
		var n uint16
		err := binary.Read(buf, binary.LittleEndian, &n)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func advance(read func(any)) {
	var null uint16
	read(&null)
}

func readNumber(read func(any)) int {
	var n uint16
	read(&n)
	advance(read)
	return int(n)
}

func readU16String(read func(any)) string {
	slen := readNumber(read)
	u16s := make(u16string, slen)
	read(u16s)
	advance(read)
	return string(utf16.Decode(u16s))
}
