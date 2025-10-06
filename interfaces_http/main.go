package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Creating a custom writer :)
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(resp)

	/*
	   	bs := make([]byte, 99999) // byte slice, empty but with a length of 99999
	   	resp.Body.Read(bs) // Read from the response body and put it inside the byte slice
	   	fmt.Println(string(bs)) // convert byte slice to string and print it

	   This is a way to read from a stream (resp.Body) but it's not efficient
	*/

	//io.Copy(os.Stdout, resp.Body)

	// copies from a source (resp.Body) to a destination (os.Stdout)
	// This is the most efficient way to read from a stream

	// TEST CUSTOM WRITER:
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// Just by defining this method, logWriter now implements the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil

}
