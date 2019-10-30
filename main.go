package main

import "net/http"

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := 
	}
}
