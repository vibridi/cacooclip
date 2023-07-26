package main // go install github.com/vibridi/cacooclip/cmd/cacooclip@latest

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/vibridi/cacooclip"
)

var (
	r = flag.Bool("r", false, "read Cacoo data from clipboard")
	w = flag.Bool("w", false, "write Cacoo data from clipboard")
)

func main() {
	flag.Parse()
	if *r {
		read()
	}
	if *w {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("scanner err:", err)
			}
			write(scanner.Text())
		}
	}
}

func read() {
	s, err := cacooclip.Read()
	if err != nil {
		fmt.Println("clipboard err:", err)
		return
	}
	fmt.Println(s)
}

func write(str string) {
	err := cacooclip.Write(str)
	if err != nil {
		fmt.Println("clipboard err:", err)
		return
	}
	fmt.Println("content written successfully")
}
