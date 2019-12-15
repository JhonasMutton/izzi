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

type IBigIDClient interface {
	Ping() error
	VerifyRG(body []byte) (*http.Response, error)
}

type BigIDClient struct {
	baseUrl    *url.URL
	httpClient *http.Client
	xApiKey    string
}

func (client BigIDClient) VerifyRG(body []byte) (*http.Response, error) {
	result, err := client.CallMongeralAegon(body, "POST", "VerifyID")
	if err != nil {
		infra.Logger.Error("Error to call Bid ID:", err.Error())
		return nil, err
	}

	if 200 != result.StatusCode {
		infra.Logger.Error("Call to Bid ID was unsuccessful")
		return nil, errors.ErrInternalServer
	}

	return result, nil
}

func (client BigIDClient) Ping() error {
	panic("implement me")
}


func (client BigIDClient) CallMongeralAegon(body []byte, method, path string) (*http.Response, error) {
	callUrl := client.baseUrl.String() + path


	request, err := client.newRequest( method, callUrl, body)
	if err != nil {
		return nil, err
	}

	response, err := client.doRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client BigIDClient) doRequest(req *http.Request) (*http.Response, error) {
	now := time.Now().Unix()

	infra.Logger.Debug(transactionIdMessage, now, " | Calling BigID in address:", req.URL)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		infra.Logger.Error(transactionIdMessage, now, " | Error inBigID request")
		return nil, err
	}

	infra.Logger.Debug(transactionIdMessage, now, " | BigID result:", resp.StatusCode)
	return resp, err
}

func (client BigIDClient) newRequest( method, callUrl string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, callUrl, bytes.NewReader(body))
	if err != nil {
		infra.Logger.Error("Error to create new request to BigID")
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("X-Api-Key", client.xApiKey)

	return req, nil
}

func NewBigIDClient() BigIDClient {
	httpClient := http.DefaultClient
	baseUrlString := os.Getenv("BIG_ID_HOST")

	baseUrl, err := url.Parse(baseUrlString)
	if err != nil {
		infra.Logger.Fatal("Error to parse BIG_ID_HOST to URL.")
	}

	xApiKey, ok := os.LookupEnv("BIG_ID_X_API_KEY")
	if !ok {
		infra.Logger.Fatal("Error to find API X KEY.")
	}

	return BigIDClient{baseUrl: baseUrl, httpClient: httpClient, xApiKey:xApiKey}
}
