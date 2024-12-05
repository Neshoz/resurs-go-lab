package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	ccm.AuthService.getServiceToken()

	req, err := http.NewRequest("GET", ccm.ServiceURL+"/organisations", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}
	req.Header.Add("Apikey", ccm.ApiKey)

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	io.WriteString(w, string(body))
}

func (auth *AuthService) getServiceToken() string {
	_, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(auth.PrivateKeyPEM))
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	if err != nil {
		log.Printf("Error parsing private key: %v", err)
	}
	token.Header["kid"] = "1"

	req, err := http.NewRequest("GET", auth.ServiceURL+"/token", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}
	req.Header.Add("Apikey", auth.ApiKey)
	req.Header.Add("Authorization", "Bearer ")

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	return string(body)
}

func (auth *AuthService) getGovermentID(personID string) string {
	req, err := http.NewRequest("GET", auth.ServiceURL+"/government_id", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}
	req.Header.Add("Apikey", auth.ApiKey)

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	return string(body)
}

type AuthService struct {
	ApiKey        string
	ServiceURL    string
	PrivateKeyPEM string
}

type CCMService struct {
	ApiKey      string
	ServiceURL  string
	AuthService AuthService
}

type Config struct {
	ApiKey        string
	ApiGatewayURL string
}

func main() {
	cfg := Config{
		ApiKey:        "59d51b41e42f4f4bb07104371b19d152",
		ApiGatewayURL: "https://apigw-integration.test.resurs.loc/api",
	}
	auth := AuthService{
		ApiKey:     cfg.ApiKey,
		ServiceURL: cfg.ApiGatewayURL + "/auth_service",
	}
	ccm := CCMService{
		ApiKey:      cfg.ApiKey,
		ServiceURL:  cfg.ApiGatewayURL + "/corporate_customer_manager_service",
		AuthService: auth,
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
