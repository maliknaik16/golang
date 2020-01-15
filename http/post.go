
package main

import (
	"encoding/json"
	"io/ioutil"
	"bytes"
	"log"
	"net/http"
)

// Source: https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
func main() {
	requestBody, err := json.Marshal(map[string]string{
		"name": "John Doe",
		"email": "johndoe@gmail.com",
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))

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

