package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseUrl = "https://api.bitflyer.com/v1/"


func (client *Client)DoRequest(apiPath, method string) (body []byte, err error) {
	baseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	apiUrl, err := url.Parse(apiPath)
	if err != nil {
		return nil, err
	}

	// ResolveReference を使って絶対URLに変換(https://)
	endpoint := baseUrl.ResolveReference(apiUrl).String()
	println(endpoint)

	
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}

	// Queryでtype Values map[string][]string　を返却する
	query := request.URL.Query()
	queryMap := map[string]string{"product_code": "ETH"}
	for key, value := range queryMap {
		query.Add(key, value)
	}

	request.URL.RawQuery = query.Encode()

	for key, value := range client.header(method, request.URL.RequestURI(), nil) {
		request.Header.Add(key, value)
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// この動作や、値しか許容しない構造体にする
type Client struct {
	apikey, secretkey string
	httpClient        *http.Client
	baseUrl           *url.URL
	log               *log.Logger
}

func New(apikey, secretkey string) *Client {
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		log:        &log.Logger{},
	}
}

func (api Client) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println(timestamp)
	message := timestamp + method + endpoint + string(body)
	mac := hmac.New(sha256.New, []byte(api.secretkey))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string {
		"ACCESS-KEY": api.apikey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN": sign,
		"Content-Type": "application/json",
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
