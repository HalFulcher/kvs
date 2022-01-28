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

func PingEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ping" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "%v", "pong")

}

func StoreEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/store" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":

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

	case "PUT":

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
	fmt.Println("Hitting get all entries handler")
	// i just have to return everything. I don't need anything from request body...
	response := store.GetAllEntries()
	fmt.Fprintf(w, "%v", response)
	fmt.Println(response)
}
