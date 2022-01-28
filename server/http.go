package server

import (
	"fmt"
	"http-server-3/store"
	"net/http"
	// "encoding/json"
)

// logging func

// handler functions that deal with http requests / responses.
// Get what you need from request object to pass into store methods/funcs. Then send func output to
//response obj for things that need to return something, i.e. GET requests.
// url query format: http://localhost:8080/store?key={key-you're-searching-for}

func PingEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ping" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("-----------\n")
	fmt.Println("pong")
	fmt.Fprintf(w, "%v", "pong")
	fmt.Println("-----------\n")
}

func StoreEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/store" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		fmt.Println("-----------\n")
		fmt.Println("Received a GET request")
		// implement get method
		params := r.URL.Query()

		var getKey string

		for k, _ := range params {
			getKey = k
		}

		//   now that we've extracted specific val from URL can pass it into
		//   store GetEntry method

		response := store.GetEntry(getKey)
		// write response to response writer.
		fmt.Fprintf(w, "%v", response)
		fmt.Println(response)

		fmt.Println("-----------\n")

	case "PUT":

		fmt.Println("-----------\n")

		fmt.Println("Received a PUT request")
		// implement get method
		putParams := r.URL.Query()
		// this accesses username
		username, _, _ := r.BasicAuth()
		var putKey string
		var putValue string

		for k, v := range putParams {
			putKey = k
			putValue = v[0]
		}

		//now that we've extracted specific val from URL can pass it into
		//store method. But should values be sent in request body? Not as query params.
		store.AddEntry(putKey, putValue, username)

	case "DELETE":

		fmt.Println("-----------\n")

		fmt.Println("Received a DELETE request")
		// implement delete method
		params := r.URL.Query()
		username, _, _ := r.BasicAuth()
		var deleteKey string
		// this accesses username

		for k, _ := range params {
			deleteKey = k
		}
		fmt.Println(deleteKey)
		//   now that we've extracted specific val from URL can pass it into
		//   store DeleteEntry method

		store.DeleteEntry(deleteKey, username)
		// success message

		fmt.Println("-----------\n")

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Not a valid endpoint.\n"))

	}

}

func GetAllEntriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("-----------\n")
	fmt.Println("Hitting get all entries handler")
	// i just have to return everything. I don't need anything from request body...
	response := store.GetAllEntries()
	fmt.Fprintf(w, "%v", response)
	fmt.Println(response)
	fmt.Println("-----------\n")
}
