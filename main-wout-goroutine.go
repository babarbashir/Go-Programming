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


func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com","http://www.gbmme.com","http://www.bing.com"}


	for _, url := range urls {
		length, err := getPage(url)
		if (err != nil) {
			fmt.Println(err)
		}
		fmt.Println("Length == ", length)
	}
}