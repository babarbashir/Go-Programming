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

func worker(urlCh chan string, sizeCh chan string, id int) {

	for {
		url := <- urlCh
		length, err := getPage(url)
		if(err == nil) {
			sizeCh <- fmt.Sprintf("length of %s is %d (%id)",url, length, id)
		} else {
			sizeCh <- fmt.Sprintf("Error getting %s is %s",url, err)
		}
	}

	
}

func generator(url string, urlCh chan string) {
	urlCh <- url
}

func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com","http://www.gbmme.com","http://www.bing.com"}
	urlCh := make(chan string)
	sizeCh := make(chan string)
	for i:=0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}
	for _, url := range urls {
		go generator(url, urlCh)
	}
	
	for i := 0; i < len(urls); i++ {
	fmt.Printf("%s\n", <-sizeCh)
	}
}