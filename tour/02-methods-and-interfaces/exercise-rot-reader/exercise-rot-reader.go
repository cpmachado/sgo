package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}

	// needs to be Go 1.16 compliant
	for i := 0; i < n; i++ {
		c := rune(p[i])
		if !unicode.IsLetter(c) {
			continue
		}
		if unicode.IsUpper(c) {
			c = ((c - 'A' + 13) % 26) + 'A'
		} else {
			c = ((c - 'a' + 13) % 26) + 'a'
		}
		p[i] = byte(c)
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
