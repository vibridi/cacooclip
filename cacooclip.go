// Copyright 2021 The golang.design Initiative Authors, 2023 Gabriele V.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Original code written by Changkun Ou <changkun.de>
// Modified to handle Cacoo web MIME type by Gabriele V. <gabriele@nulab.com>

package cacooclip

import (
	"errors"
	"fmt"

	"github.com/vibridi/cacooclip/internal/chromium"
)

const (
	CacooWebMIMEType = "cacoo/shape"
)

var (
	errUnavailable  = errors.New("unavailable or no data from clipboard")
	errWriteFailure = errors.New("unavailable or no data was written to the clipboard")
	errNoCacooMime  = errors.New("no content with Cacoo MIME type in clipboard")
)

// Read reads Cacoo shape data from the clipboard.
// Not thread safe. Concurrent reads on some OS'es including Darwin may panic.
func Read() (string, error) {
	b, err := read()
	if err != nil {
		return "", fmt.Errorf("read clipboard err: %w", err)
	}
	wcmap, err := chromium.DecodeWebCustomMIMEData(b)
	if err != nil {
		return "", fmt.Errorf("decode err: %w", err)
	}
	content, ok := wcmap[CacooWebMIMEType]
	if !ok {
		return "", errNoCacooMime
	}
	return content, nil
}

// Write writes Cacoo shape data to the clipboard so that you can CTRL+V directly onto the Cacoo canvas.
func Write(str string) error {
	enc, err := chromium.EncodeWebCustomMIMEData(map[string]string{CacooWebMIMEType: str})
	if err != nil {
		return fmt.Errorf("encode err: %w", err)
	}
	return write(enc)
}
