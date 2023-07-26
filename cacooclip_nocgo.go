//go:build !windows && !cgo

package cacooclip

func read() (buf []byte, err error) {
	panic("clipboard: cannot use when CGO_ENABLED=0")
}

func write(buf []byte) (<-chan struct{}, error) {
	panic("clipboard: cannot use when CGO_ENABLED=0")
}
