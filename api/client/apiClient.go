package client

import (
	"bytes"
	"coastrade/domain/model"
	"coastrade/infrastructure"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseUrl = "https://api.bitflyer.com/v1/"

type APIClient interface {
	SendOrder(order *infrastructure.Order, product string) (*infrastructure.ResponseSendChildOrder, error)
	ListOrder(query map[string]string, product string) ([]infrastructure.Order, error)
	GetBalance(product string) ([]model.Balance, error)
}

func New(apikey, secretkey string) *Client {
	return &Client{
		apikey:     apikey,
		secretkey:  secretkey,
		httpClient: &http.Client{},
		log:        &log.Logger{},
	}
}

// この動作や、値しか許容しない構造体にする
type Client struct {
	apikey, secretkey string
	httpClient        *http.Client
	baseUrl           *url.URL
	log               *log.Logger
}

func (client *Client) DoRequest(apiPath, method, product string, queryMap map[string]string, data []byte) (body []byte, err error) {
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

	request, err := http.NewRequest(method, apiPath, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// Queryでtype Values map[string][]string　を返却する
	query := request.URL.Query()
	queryMap["product"] = product
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

func (client *Client) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println(timestamp)
	message := timestamp + method + endpoint + string(body)
	mac := hmac.New(sha256.New, []byte(client.secretkey))
	mac.Write([]byte(message))
	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS-KEY":       client.apikey,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

func (client *Client) SendOrder(order *infrastructure.Order, product string) (*infrastructure.ResponseSendChildOrder, error) {
	data, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}
	logger, _ := zap.NewDevelopment()
	logger.Info("BuyOrder", zap.Time("BuyTime", time.Now()), zap.Object("Order", order))
	resp, err := client.DoRequest("POST", "me/sendchildorder", product, map[string]string{}, data)
	if err != nil {
		logger.Error("BuyOrder is Fail", zap.Time("Error Time", time.Now()), zap.Object("SendOrder", order))
		return nil, err
	}
	var response infrastructure.ResponseSendChildOrder
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *Client) GetBalance(product string) ([]model.Balance, error) {
	api := "me/getbalance"
	resp, err := client.DoRequest("GET", api, product, map[string]string{}, nil)
	log.Printf("api=%s resp=%s", api, string(resp))
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	var balance []model.Balance

	err = json.Unmarshal(resp, &balance)
	if err != nil {
		log.Printf("action=GetBalance err=%s", err.Error())
		return nil, err
	}
	return balance, nil
}

func (client *Client) ListOrder(query map[string]string, product string) ([]infrastructure.Order, error) {
	resp, err := client.DoRequest("GET", "me/getchildorders", product, query, nil)
	if err != nil {
		return nil, err
	}
	var responseListOrder []infrastructure.Order
	err = json.Unmarshal(resp, &responseListOrder)
	if err != nil {
		return nil, err
	}
	return responseListOrder, nil
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
