package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request\n")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "Hello at root")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "At /hello")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", getRoot)
	router.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":5000", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func testMiddleware(handler func(http.ResponseWriter, *http.Request)) {
	
}
