package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fname := "-"

	var r io.Reader
	switch fname {
	case "", "-":
		r = os.Stdin
	default:
		f, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "file open error.")
		}
		defer f.Close()
		r = f
	}

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "file read error.")
		}
		fmt.Print(string(buf[:n]))
	}
}

