package cdn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Cache refreshing
//
// https://docs.azure.cn/en-us/cdn/cdn-api-add-purge
func (c *Client) AddPurge(request *AddPurgeRequest) (resp *http.Response, result *AddPurgeResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/purges?apiVersion=1.0", request.EndpointID), nil)
	postBody, _ := json.Marshal(request.Body)
	resp, err = c.Request(http.MethodPost, reqUrl, postBody, &result)
	return
}

type AddPurgeRequest struct {
	EndpointID string //Target node unique identifier
	Body       AddPurgeRequestBody
}

type AddPurgeRequestBody struct {
	//When refreshing the file list, the path must be an absolute path. For example: http://example.com/pictures/city.png
	Files []string
	//When refreshing the directory list, the path must be an absolute path. For example: http://example.com/pictures/
	Directories []string
}

type AddPurgeResponse TaskResponse

// Query prefetch progress
//
// https://docs.azure.cn/en-us/cdn/cdn-api-query-preload
func (c *Client) QueryPreload(request *QueryPreloadRequest) (resp *http.Response, result *QueryPreloadResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/preloads/%s?apiVersion=1.0", request.EndpointID, request.PreloadID), nil)
	resp, err = c.Request(http.MethodGet, reqUrl, nil, &result)
	return
}

type QueryPreloadRequest struct {
	EndpointID string //Target node unique identifier
	PreloadID  string //Prefetch operation unique identifier
}

// Prefetch status
type PrefetchStatus string

const (
	PrefetchStatusWaiting PrefetchStatus = "Waiting" //Waiting
	PrefetchStatusRunning PrefetchStatus = "Running" //Running
	PrefetchStatusSucceed PrefetchStatus = "Succeed" //Succeeded
	PrefetchStatusFailed  PrefetchStatus = "Failed"  //Failed
)

type QueryPreloadResponse struct {
	Files []struct {
		Url    string
		Status PrefetchStatus // Prefetch status
	} // File prefetch results
}

// Preloading
//
// https://docs.azure.cn/en-us/cdn/cdn-api-add-preload
func (c *Client) AddPreload(request *AddPreloadRequest) (resp *http.Response, result *AddPreloadResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/preloads?apiVersion=1.0", request.EndpointID), nil)
	postBody, _ := json.Marshal(request.Body)
	resp, err = c.Request(http.MethodPost, reqUrl, postBody, &result)
	return
}

type AddPreloadRequest struct {
	EndpointID string //Target node unique identifier
	Body       AddPreloadRequestBody
}

type AddPreloadRequestBody struct {
	//When refreshing the file list, the path must be an absolute path. For example:http://example.com/pictures/city.png
	Files []string
}

type AddPreloadResponse TaskResponse

// Check cache refresh progress
//
// https://docs.azure.cn/en-us/cdn/cdn-api-query-purge
func (c *Client) QueryPurge(request *QueryPurgeRequest) (resp *http.Response, result *QueryPurgeResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/purges/%s?apiVersion=1.0", request.EndpointID, request.PurgeID), nil)
	resp, err = c.Request(http.MethodGet, reqUrl, nil, &result)
	return
}

type QueryPurgeRequest struct {
	EndpointID string //Target node unique identifier
	PurgeID    string //Cache refresh operation unique identifier
}

// Refresh status
type RefreshStatus string

const (
	RefreshStatusRunning RefreshStatus = "Running" //Running
	RefreshStatusSucceed RefreshStatus = "Succeed" //Succeeded
	RefreshStatusFailed  RefreshStatus = "Failed"  //Failed
)

type QueryPurgeResponse struct {
	Files []struct {
		Url    string
		Status RefreshStatus //Refresh status
	} // File refresh results
	Directories []struct {
		Url    string
		Status RefreshStatus //Refresh status
	} // Directory refresh results
}
