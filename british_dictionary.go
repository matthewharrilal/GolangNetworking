package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	// Do we have any added configurations to this request such as a timeout? No then we don't need a client and can execute from our http

	// We then need to read the contents until there are no contents left  from the data source we are creating
}
