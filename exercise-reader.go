package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(bytes []byte) (int, error) { 
	var i int
	var ascii byte
	letter := 'A'
	
	ascii = byte(int(letter))
	
	for i = range bytes { 
		bytes[i] = ascii
	}

	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
