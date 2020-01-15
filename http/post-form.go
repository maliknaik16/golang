
package main

import (
	"log"
	"encoding/json"
	"net/http"
	"net/url"
)

// Source: https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
func main() {
	formData := url.Values{
		"name": {"john", "doe"},
	}

	resp, err := http.PostForm("http://httpbin.org/post", formData)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
}
