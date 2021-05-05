package client

import (
	config "coastrade/configs"
	"errors"
	"log"
	"net/http"
	"net/url"
)

func doRequest(apiPath string) (string, error) {
	client, err := NewClient(config.Config.ApiKey,
		config.Config.ApiSecret,
		config.Config.BaseUrl)
	if err != nil {
		return "", err
	}
	return client.apikey, nil
}

// この動作や、値しか許容しない構造体にする
type Client struct {
	apikey, secretkey string
	httpClient        *http.Client
	baseUrl           *url.URL
	log               *log.Logger
}

func New(apikey, secretkey string, baseUrl *url.URL) *Client {
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		baseUrl:    baseUrl,
		log:        &log.Logger{},
	}
}

func NewClient(apikey, secretkey, baseUrlstr string) (*Client, error) {
	baseurl, err := url.ParseRequestURI(baseUrlstr)
	if err != nil {
		return nil, err
	}
	if len(apikey) == 0 {
		return nil, errors.New("apikey is empty")
	}
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		baseUrl:    baseurl,
		log:        &log.Logger{},
	}, nil
}
