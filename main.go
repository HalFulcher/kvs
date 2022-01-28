package main

import (
	"http-server-3/logfile"
	"http-server-3/server"
	"http-server-3/store"
	"log"
	"net/http"
	"os"
)

func main() {

	// listening on channel for jobs

	store.StartListening()

	// logging

	fileName := "logfile/httprequests.log"

	// https://www.socketloop.com/tutorials/golang-how-to-save-log-messages-to-file
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to webrequests.log
	log.SetOutput(logFile)

	//endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", server.PingEndpointHandler)
	mux.HandleFunc("/store", server.StoreEndpointHandler)
	mux.HandleFunc("/list", server.GetAllEntriesHandler)

	//start server
	http.ListenAndServe("127.0.0.1:8080", logfile.RequestLogger(mux))
}
