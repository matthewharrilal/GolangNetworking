package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type BookInfo struct {
	Result Result `json:"result"`
}

type Result struct {
	Definition string `json:"definition"`

	Label string `json:"label"`

	PrimaryTopic PrimaryTopic `json:"primaryTopic"`
}

type PrimaryTopic struct {
	Subject []Subject `json:"subject"`
	Title   string    `json:"title"`
}

type Subject struct {
	About string `json:"_about"`

	Label string `json:"label"`
}

func main() {

	// SO let us analyze the components of a network request and see what we are missing

	// We need first the url that we are going to be hitting
	url := "http://bnb.data.bl.uk/doc/resource/012690955.json"

	// Do we have any added configurations to this request such as a timeout? No then we don't need a client and can execute from our http
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	// Once we have that response body we then need to unmarshal that data into our Book info object ... but first we need to read out our output from the response body and the use it
	body, readErr := ioutil.ReadAll(resp.Body) // Read

	if readErr != nil {
		log.Fatal(readErr) // If there was a problem other than reading to the end of the given datasource
	}

	// Now that we have read that data we have to unmarshal that data to our book object
	bookInfo := &BookInfo{} // Instantiate a book info variable to be a pointer to our book info object that has been instantiated

	fmt.Println(string(body))
	unMarshaledBody := json.Unmarshal([]byte(body), *bookInfo)

	fmt.Println(unMarshaledBody)

}
