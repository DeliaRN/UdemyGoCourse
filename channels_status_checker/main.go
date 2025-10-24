package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://amazon.com",
	}

	//checkLinksInOrder(links)

	c := make(chan string) // Create a channel of type string

	for _, link := range links {
		go checkLink(link, c)
	}

	/*
		for { // infinite loop
			go checkLink(<-c, c)
		}
	*/
	for link := range c {
		// time.Sleep(5 * time.Second) WE DON'T WANT TO SLEEP THE MAIN ROUTINE
		// go checkLink(link, c)
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // receiving link on channel and sending it back to main
		return
	}
	fmt.Println(link, "is up!")
	c <- link // receiving link on channel and sending it back to main
}
