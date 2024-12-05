package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

func (ccm *CCMService) getOrganisations(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Name string `json:"name"`
	}

	io.WriteString(w, "At /organisations")
}

type CCMService struct {
	serviceURL string
}

func main() {
	ccm := CCMService{
		serviceURL: "https://apigw.integration.resurs.com/api/customer_manager_service",
	}
	router := http.NewServeMux()
	router.HandleFunc("GET /", getRoot)
	router.HandleFunc("GET /hello", getHello)
	router.HandleFunc("GET /organisations", ccm.getOrganisations)

	err := http.ListenAndServe(":8080", RequestLoggerMiddleware(router))

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s %s\n", time.Now().Format("2006-01-02 15:04:05.000000000"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
