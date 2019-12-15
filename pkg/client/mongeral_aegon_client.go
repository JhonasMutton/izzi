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

type IMongeralAegonClient interface {
	Ping() error
	Simulation(body []byte) (*http.Response, error)
}

type MongeralAegonClient struct {
	baseUrl    *url.URL
	httpClient *http.Client
	xApiKey    string
	cnpjPartner string
}

func (client MongeralAegonClient) Ping() error {
	panic("implement me")
}

const transactionIdMessage = "\t TransactionId: "

func (client MongeralAegonClient) Simulation(body []byte) (*http.Response, error) {
	result, err := client.CallMongeralAegon(body, "POST", "simulacao?cnpj="+client.cnpjPartner+"&&codigoModeloProposta=YX")
	if err != nil {
		infra.Logger.Error("Error to call Mongeral Aegon:", err.Error())
		return nil, err
	}

	if 200 != result.StatusCode {
		infra.Logger.Error("Call to Mongeral Aegon was unsuccessful")
		return nil, errors.ErrInternalServer
	}

	return result, nil
}

func (client MongeralAegonClient) CallMongeralAegon(body []byte, method, path string) (*http.Response, error) {
	callUrl := client.baseUrl.String() + path


	request, err := client.newRequest(method, callUrl, body)
	if err != nil {
		return nil, err
	}

	response, err := client.doRequest(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client MongeralAegonClient) doRequest(req *http.Request) (*http.Response, error) {
	now := time.Now().Unix()

	infra.Logger.Debug(transactionIdMessage, now, " | Calling Mongeral Aegon in address:", req.URL)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		infra.Logger.Error(transactionIdMessage, now, " | Error in Mongeral Aegon request")
		return nil, err
	}

	infra.Logger.Debug(transactionIdMessage, now, " | Mongeral Aegon result:", resp.StatusCode)
	return resp, err
}

func (client MongeralAegonClient) newRequest( method, callUrl string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, callUrl, bytes.NewReader(body))
	if err != nil {
		infra.Logger.Error("Error to create new request to Mongeral Aegon")
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("X-Api-Key", client.xApiKey)
	return req, nil
}

func NewMongeralAegonClient() MongeralAegonClient {
	httpClient := http.DefaultClient
	baseUrlString := os.Getenv("MONGERAL_AEGON_HOST")

	baseUrl, err := url.Parse(baseUrlString)
	if err != nil {
		infra.Logger.Fatal("Error to parse MONGERAL_AEGON_HOST to URL.")
	}

	xApiKey, ok := os.LookupEnv("MONGERAL_AEGON_X_API_KEY")
	if !ok {
		infra.Logger.Fatal("Error to find API X KEY.")
	}

	cnpjPartner, ok := os.LookupEnv("MONGERAL_AEGON_CNPJ_PARTNER")
	if !ok {
		infra.Logger.Fatal("Error to find CNPJ Partner.")
	}

	return MongeralAegonClient{baseUrl: baseUrl, httpClient: httpClient, xApiKey: xApiKey, cnpjPartner:cnpjPartner}
}
