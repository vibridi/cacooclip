package main

import (
	"fmt"

	"github.com/vibridi/cacooclip"
)

func main() {
	s, err := cacooclip.Read()
	if err != nil {
		fmt.Println("clipboard err:", err)
		return
	}
	fmt.Println("clipboard content:")
	fmt.Println(s)
}
