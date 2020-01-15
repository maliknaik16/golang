
package main

import (
	"log"
	"io/ioutil"
	"net/http"
)

// Source: https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
func main() {
	resp, err := http.Get("https://httpbin.org/get")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
