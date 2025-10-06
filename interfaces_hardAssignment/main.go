package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct{}

func main() {

	fmt.Println(os.Args) // command line arguments
	file := os.Args[1]   // give mes the first argument after the program name

	resp, err := os.Open(file)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fr := fileReader{}
	io.Copy(fr, resp)
}

func (fileReader) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
