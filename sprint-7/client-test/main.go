package main

import (
	"io"
	"net/http"
)

type PostalCodeClient struct {
	Client  *http.Client
	baseURL string
}

func (api *PostalCodeClient) PostalCode(address string) (string, error) {
	req, err := http.NewRequest("GET", api.baseURL+"/getPostalCode", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("address", address)
	req.URL.RawQuery = q.Encode()

	resp, err := api.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {

}
