package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const baseUrl = "https://apigw-integration.test.resurs.loc/api/cib_gateway_service"

func GetMe(w http.ResponseWriter, r *http.Request) (*User, error) {
	contextId := url.QueryEscape(r.PathValue("contextId"))
	url := fmt.Sprintf(baseUrl + "/users/%s/me", contextId)
	request, _ := http.NewRequest(http.MethodGet, url, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	data := User{}
	json.NewDecoder(res.Body).Decode(&data)

	return &data, nil
}