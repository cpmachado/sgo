package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(p []byte) (n int, err error) {
	p[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
