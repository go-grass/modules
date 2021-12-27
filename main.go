package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://blog.logrocket.com/making-http-requests-in-go/")
	if err != nil {
		log.Fatalln(err)
	}

	// Read the response body on the line below.
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body type to string
	sb := string(body)
	log.Printf(sb)
}
