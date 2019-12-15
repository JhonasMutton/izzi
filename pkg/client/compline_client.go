package client

import (
	"bytes"
	"github.com/JhonasMutton/izzi/internal/errors"
	"github.com/JhonasMutton/izzi/internal/infra"
	"net/http"
	"net/url"
	"os"
	"time"
)

type IComplineClient interface {
	Ping() error
	Search(body []byte, customerId string) (*http.Response, error)
}

type ComplineClient struct {
	baseUrl    *url.URL
	httpClient *http.Client
}

func (client ComplineClient) Ping() error {
	panic("implement me")
}

func (client ComplineClient) Search(body []byte, customerId string) (*http.Response, error) {
	result, err := client.CallMongeralAegon(body, "POST", "/trace/search", customerId)
	if err != nil {
		infra.Logger.Error("Error to call Compline:", err.Error())
		return nil, err
	}

	if 200 != result.StatusCode {
		infra.Logger.Error("Call to Compline was unsuccessful")
		return nil, errors.ErrInternalServer
	}

	return result, nil
}

func (client ComplineClient) CallMongeralAegon(body []byte, method, path, customerId string) (*http.Response, error) {
	callUrl := client.baseUrl.String() + path


	request, err := client.newRequest(customerId, method, callUrl, body)
	if err != nil {
		return nil, err
	}

	response, err := client.doRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client ComplineClient) doRequest(req *http.Request) (*http.Response, error) {
	now := time.Now().Unix()

	infra.Logger.Debug(transactionIdMessage, now, " | Calling Compline in address:", req.URL)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		infra.Logger.Error(transactionIdMessage, now, " | Error inCompline request")
		return nil, err
	}

	infra.Logger.Debug(transactionIdMessage, now, " | Compline result:", resp.StatusCode)
	return resp, err
}

func (client ComplineClient) newRequest(customerId, method, callUrl string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, callUrl, bytes.NewReader(body))
	if err != nil {
		infra.Logger.Error("Error to create new request to Compline")
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("customerId", customerId)
	return req, nil
}

func NewComplineClientt() ComplineClient {
	httpClient := http.DefaultClient
	baseUrlString := os.Getenv("COMPLINE_HOST")

	baseUrl, err := url.Parse(baseUrlString)
	if err != nil {
		infra.Logger.Fatal("Error to parse COMPLINE_HOST to URL.")
	}

	return ComplineClient{baseUrl: baseUrl, httpClient: httpClient}
}
