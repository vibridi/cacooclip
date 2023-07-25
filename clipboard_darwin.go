// Copyright 2021 The golang.design Initiative Authors, 2023 Gabriele V.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Original code written by Changkun Ou <changkun.de>
// Modified by vibridi (Gabriele V.) <gabriele@nulab.com>

//go:build darwin && !ios

package cacooclip

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

int clipboard_read_cacoo(void **out);
int clipboard_write_cacoo(const void *bytes, NSInteger n);
NSInteger clipboard_change_count();
*/
import "C"
import (
	"time"
	"unsafe"
)

func read() (buf []byte, err error) {
	var (
		data unsafe.Pointer
		n    C.int
	)
	n = C.clipboard_read_cacoo(&data)
	if n == -1 {
		return nil, errNoCacooMime
	}
	if n == 0 && data == nil {
		return nil, errUnavailable
	}
	defer C.free(data)
	if n == 0 {
		return nil, nil
	}
	return C.GoBytes(data, C.int(n)), nil
}

// write writes the given data to clipboard and
// returns true if success or false if failed.
func write(buf []byte) (<-chan struct{}, error) {
	var ok C.int
	if len(buf) == 0 {
		ok = C.clipboard_write_cacoo(unsafe.Pointer(nil), 0)
	} else {
		ok = C.clipboard_write_cacoo(unsafe.Pointer(&buf[0]), C.NSInteger(len(buf)))
	}
	if ok != 0 {
		return nil, errUnavailable
	}

	// use unbuffered data to prevent goroutine leak
	changed := make(chan struct{}, 1)
	cnt := C.long(C.clipboard_change_count())
	go func() {
		for {
			// not sure if we are too slow or the user too fast :)
			time.Sleep(time.Second)
			cur := C.long(C.clipboard_change_count())
			if cnt != cur {
				changed <- struct{}{}
				close(changed)
				return
			}
		}
	}()
	return changed, nil
}
