package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//getPage func

func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if(err != nil) {
		return 0, err
	}

	defer resp.Body.Close()  

	body, err := ioutil.ReadAll(resp.Body)
	if(err != nil) {
		return 0, err
	}

	return len(body), nil

}

func getter(url string, size chan string) {
	length, err := getPage(url)
	if(err == nil) {
		size <- fmt.Sprintf("length of %s is %d",url, length)
	}
	
}

func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com","http://www.gbmme.com","http://www.bing.com"}

	size := make(chan string)

	for _, url := range urls {
		go getter(url, size)		
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-size)
	}
}