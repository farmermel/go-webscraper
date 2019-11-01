package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	body := handleGet()
	println(body)
}

func handleGet() string {

	url := os.Args[1:]

	//check url[0] is a string/properly formatted
	resp, err := http.Get(url[0])

	if err != nil {
		println(err)
	}

	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			println(err)
		}

	}

	return string(bodyBytes)
}
