package cdn

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func NewClient(keyID string, keyValue string, subscriptionID string) *Client {
	return &Client{
		RestAPIEndpoint: "restapi.cdn.azure.cn",
		HTTPClient:      http.DefaultClient,
		KeyID:           keyID,
		KeyValue:        keyValue,
		SubscriptionID:  subscriptionID,
	}
}

type Client struct {
	HTTPClient      *http.Client
	RestAPIEndpoint string
	SubscriptionID  string
	KeyID           string
	KeyValue        string
}

func (c *Client) MakeRequestUrl(path string, query url.Values) url.URL {
	if query == nil {
		query = url.Values{}
	}
	u, _ := url.Parse(fmt.Sprintf("https://%s/subscriptions/%s%s", c.RestAPIEndpoint, c.SubscriptionID, path))
	for k, v := range u.Query() {
		query[k] = v
	}
	return *u
}

func (c *Client) Request(method string, uri url.URL, body []byte, result any) (resp *http.Response, err error) {
	var (
		req          *http.Request
		responseBody []byte
	)
	requestTime := time.Now().UTC().Format("2006-01-02 15:04:05")
	if req, err = http.NewRequest(method, uri.String(), bytes.NewBuffer(body)); err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("x-azurecdn-request-date", requestTime)
	req.Header.Set("Authorization", c.CalculateAuthorizationHeader(uri, requestTime, method))
	if resp, err = c.HTTPClient.Do(req); err != nil {
		return nil, err
	}
	if responseBody, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	responseError := &ErrorResponse{}
	if err = json.Unmarshal(responseBody, &responseError); err == nil &&
		responseError.Succeeded != nil && !*responseError.Succeeded {
		return resp, responseError
	}
	if err = json.Unmarshal(responseBody, &result); err != nil {
		return resp, err
	}
	return resp, err
}

type TaskResponse struct {
	Succeeded bool
	IsAsync   bool
	AsyncInfo struct {
		TaskTrackId string
		TaskStatus  TaskStatus
	}
}

type ErrorResponse struct {
	Succeeded *bool
	ErrorInfo *struct {
		Type    string
		Message string
	}
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorInfo.Type, e.ErrorInfo.Message)
}

func (c *Client) CalculateAuthorizationHeader(requestURL url.URL, requestTime, httpMethod string) string {
	var path = requestURL.Path
	m, _ := url.ParseQuery(requestURL.RawQuery)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var orderedQueries []string
	for _, k := range keys {
		orderedQueries = append(orderedQueries, fmt.Sprintf("%s:%s", k, m[k][0]))
	}

	var queries = strings.Join(orderedQueries, ", ")
	content := fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s", path, queries, requestTime, httpMethod)
	hash := hmac.New(sha256.New, []byte(c.KeyValue))
	hash.Write([]byte(content))
	digest := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	return fmt.Sprintf("AzureCDN %s:%s", c.KeyID, digest)
}
