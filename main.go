package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func getMe(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf(
		"https://apigw-integration.test.resurs.loc/api/cib_gateway_service/users/%s/me",
		url.QueryEscape(r.PathValue("contextId")),
	)

	request, _ := http.NewRequest(http.MethodGet, requestUrl, nil)
	request.Header.Set("Apikey", r.Header.Get("Apikey"))
	request.Header.Set("Authorization", r.Header.Get("Authorization"))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}
	
	var data interface{}
	json.NewDecoder(res.Body).Decode(&data)
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/api/{contextId}/me", getMe)

	err := http.ListenAndServe(":5000", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
