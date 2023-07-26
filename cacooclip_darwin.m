// Copyright 2021 The golang.design Initiative Authors, 2023 Gabriele V.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Original code written by Changkun Ou <changkun.de>
// Modified to support chromium web custom MIME type by Gabriel V. <gabriele@nulab.com>

//go:build darwin && !ios

// Interact with NSPasteboard using Objective-C
// https://developer.apple.com/documentation/appkit/nspasteboard?language=objc

#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>

NSString *const chromiumWebCustomData = @"org.chromium.web-custom-data";

int clipboard_read_cacoo(void **out) {
	NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];

 	NSArray *dataTypes = [pasteboard types];
	if (![dataTypes containsObject:chromiumWebCustomData]) {
        return -1;
	}
    NSData *data = [pasteboard dataForType:chromiumWebCustomData];
	if (data == nil) {
		return 0;
	}
	NSUInteger siz = [data length];
	*out = malloc(siz);
	[data getBytes: *out length: siz];
	return siz;
}

int clipboard_write_cacoo(const void *bytes, NSInteger n) {
	NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
	NSData *data = [NSData dataWithBytes: bytes length: n];
	[pasteboard clearContents];
	BOOL ok = [pasteboard setData: data forType:chromiumWebCustomData];
	if (!ok) {
		return -1;
	}
	return 0;
}

NSInteger clipboard_change_count() {
	return [[NSPasteboard generalPasteboard] changeCount];
}
