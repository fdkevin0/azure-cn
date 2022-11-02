package cdn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create nodes
//
// https://docs.azure.cn/en-us/cdn/cdn-api-create-endpoint
func (c *Client) CreateEndpoint(body CreateEndpointRequestBody) (resp *http.Response, result *CreateEndpointResponse, err error) {
	reqUrl := c.MakeRequestUrl("/endpoints?apiVersion=1.0", nil)
	postBody, _ := json.Marshal(body)
	resp, err = c.Request(http.MethodPost, reqUrl, postBody, &result)
	return resp, result, err
}

// Acceleration type
type ServiceType string

const (
	ServiceTypeWeb             ServiceType = "Web"             //Web page acceleration
	ServiceTypeDownload        ServiceType = "Download"        //Download acceleration
	ServiceTypeVOD             ServiceType = "VOD"             //On-demand acceleration
	ServiceTypeLiveStreaming   ServiceType = "LiveStreaming"   //Live-streaming acceleration
	ServiceTypeImageProcessing ServiceType = "ImageProcessing" //Image-processing acceleration
)

type CreateEndpointRequestBody struct {
	CustomDomain string //Accelerated domain names
	Host         string //Return-to-source host header
	ICP          string //ICP record number
	Origin       struct {
		Addresses []string //Return-to-source address collection
	}
	ServiceType ServiceType //Acceleration type
}

type Endpoint struct {
	EndpointID string
	Settings   struct {
		CustomDomain string
		Host         string
		ICP          string
		Origin       struct {
			Addresses []string
		}
		ServiceType string
	}
	Status struct {
		Enabled          bool
		ICPVerifyStatus  string
		LifetimeStatus   string
		CNameConfigured  bool
		FreeTrialExpired bool
		TimeLastUpdated  string
	}
}

type CreateEndpointResponse Endpoint

// Delete nodes
//
// https://docs.azure.cn/en-us/cdn/cdn-api-delete-endpoint
func (c *Client) DeleteEndpoint(request *DeleteEndpointRequest) (resp *http.Response, result *DeleteEndpointResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodPost, reqUrl, nil, &result)
	return resp, result, err
}

type DeleteEndpointRequest struct {
	EndpointID string //Target node unique identifier
}

type DeleteEndpointResponse TaskResponse

// Enable nodes
//
// https://docs.azure.cn/en-us/cdn/cdn-api-enable-endpoint
func (c *Client) EnableEndpoint(request *DeleteEndpointRequest) (resp *http.Response, result *DeleteEndpointResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/enable?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodPost, reqUrl, nil, &result)
	return resp, result, err
}

type EnableEndpointRequest struct {
	EndpointID string //Target node unique identifier
}

type EnableEndpointResponse TaskResponse

// Disable nodes
//
// https://docs.azure.cn/en-us/cdn/cdn-api-disable-endpoint
func (c *Client) DisableEndpoint(request *DisableEndpointRequest) (resp *http.Response, result *DisableEndpointResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/disable?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodPost, reqUrl, nil, &result)
	return
}

type DisableEndpointRequest struct {
	EndpointID string //Target node unique identifier
}

type DisableEndpointResponse TaskResponse

// Cache rule configuration
//
// https://docs.azure.cn/en-us/cdn/cdn-api-update-cache-policy
func (c *Client) UpdateCachePolicy(request *UpdateCachePolicyRequest) (resp *http.Response, result *TaskResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/cacherules?apiVersion=1.0", request.EndpointID), nil)
	body, _ := json.Marshal(request)
	resp, err = c.Request(http.MethodPut, reqUrl, body, &result)
	return
}

type UpdateCachePolicyRequest struct {
	EndpointID string //Target node unique identifier
	Body       *UpdateCachePolicyRequestBody
}

type CachePolicy struct {
	Rules              []CachePolicyRule
	IgnoreCacheControl bool //Indicates whether to ignore the cache-control header in the returned header and cache the request content.
	IgnoreCookie       bool //Indicates whether to ignore the set-cookie header in the returned header and cache the request content.
	IgnoreQueryString  bool //Indicates whether to ignore the query parameter and cache the request content.
}

type UpdateCachePolicyRequestBody CachePolicy

// Cache rule type
type CachePolicyRuleType string

const (
	CachePolicyRuleTypeSuffix CachePolicyRuleType = "Suffix"  //Cache based on the file extension
	CachePolicyRuleTypeDir    CachePolicyRuleType = "Dir"     //Cache all files in the specified directory
	CachePolicyRuleFullUri    CachePolicyRuleType = "FullUri" //Cache files at a specific path
)

type CachePolicyRule struct {
	Type  CachePolicyRuleType //Cache rule type
	Items []string
	TTL   int64 //Cache time, in seconds.
}

type UpdateCachePolicyResponse TaskResponse

// Deploy HTTPS
//
// https://docs.azure.cn/en-us/cdn/cdn-create-https-binding
func (c *Client) CreateHttpsBinding(request *CreateHttpsBindingRequestBody) (resp *http.Response, result *CreateHttpsBindingResponse, err error) {
	reqUrl := c.MakeRequestUrl("/https/bindings?apiVersion=1.0", nil)
	body, _ := json.Marshal(request)
	resp, err = c.Request(http.MethodPost, reqUrl, body, &result)
	return
}

type OriginProtocol string

const (
	OriginProtocolHttp          OriginProtocol = "Http"          //Http回源
	OriginProtocolHttps         OriginProtocol = "Https"         //Https回源
	OriginProtocolFollowRequest OriginProtocol = "FollowRequest" //协议跟随回源
)

type CreateHttpsBindingRequestBody struct {
	CertificateID     string //HTTPS证书唯一标识
	EndpointID        string //节点唯一标识
	OriginProtocol    string //回源协议
	AutoHTTPSRedirect bool   //是否自动302跳转https
}

type CreateHttpsBindingResponse TaskResponse

// Update node details
//
// https://docs.azure.cn/zh-cn/cdn/cdn-api-update-endpoint
func (c *Client) UpdateEndpoint(request *UpdateEndpointRequest) (resp *http.Response, result *UpdateEndpointResponse, err error) {
	body, _ := json.Marshal(request)
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodPut, reqUrl, body, &result)
	return resp, result, err
}

type UpdateEndpointRequest struct {
	EndpointID string //Target node unique identifier
	Body       UpdateCachePolicyRequestBody
}

type UpdateEndpointRequestBody struct {
	EndpointSettings struct {
		Host   *string //Return-to-source host header
		Origin *struct {
			Addresses []string //Return-to-source address collection
		}
	}

	//Update flag
	// 	Origin: source station
	// 	HostHeader: return-to-source host header
	UpdateFlag string
}

type UpdateEndpointResponse TaskResponse

// Access control configuration
//
// https://docs.azure.cn/en-us/cdn/cdn-api-update-access-control
func (c *Client) PutAccessControlConfiguration(request *PutAccessControlConfigurationRequest) (resp *http.Response, result *PutAccessControlConfigurationResponse, err error) {
	body, _ := json.Marshal(request)
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/accesscontrol?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodPut, reqUrl, body, &result)
	return resp, result, err
}

type PutAccessControlConfigurationRequest struct {
	EndpointID string //Target node unique identifier
	Body       PutAccessControlConfigurationRequestBody
}

// Anti-theft chain type
type RefererControlType string

const (
	//A whitelist, which allows links for only the matching referrers to access the matching PathPatterns paths.
	RefererControlTypeAllowList RefererControlType = "AllowList"
	// A blacklist, with links to matching referrers who will be denied access if they attempt to access the matching PathPatterns paths.
	RefererControlTypeBlockList RefererControlType = "BlockList"
)

type PutAccessControlConfigurationRequestBody struct {
	ForbiddenIps   []string //List of forbidden IPs
	RefererControl struct {
		Enabled            bool
		PathPatterns       []string //[anti-]theft link file path connection
		Referers           []string //[anti-]theft link connection
		RefererControlType string   //Anti-theft chain type
	}
}

type PutAccessControlConfigurationResponse TaskResponse

// Get cache rule information
//
// https://docs.azure.cn/en-us/cdn/cdn-api-get-cache-policy
func (c *Client) GetCachePolicy(request *GetCachePolicyRequest) (resp *http.Response, result *GetCachePolicyResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/cacherules?apiVersion=1.0", request.EndpointID), nil)
	body, _ := json.Marshal(request)
	resp, err = c.Request(http.MethodPut, reqUrl, body, &result)
	return
}

type GetCachePolicyRequest struct {
	EndpointID string //Target node unique identifier
}

type GetCachePolicyResponse CachePolicy

// Get node information
//
// https://docs.azure.cn/en-us/cdn/cdn-api-get-endpoint
func (c *Client) GetEndpoint(request *GetEndpointRequest) (resp *http.Response, result *GetEndpointResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s?apiVersion=1.0", request.EndpointID), nil)
	resp, err = c.Request(http.MethodGet, reqUrl, nil, &result)
	return
}

type GetEndpointRequest struct {
	EndpointID string //Target node unique identifier
}

type GetEndpointResponse Endpoint

// Get information for all subscribed nodes
//
// https://docs.azure.cn/en-us/cdn/cdn-api-list-endpoints
func (c *Client) ListEndpoints() (resp *http.Response, result *ListEndpointsResponse, err error) {
	resp, err = c.Request(http.MethodGet, c.MakeRequestUrl("/endpoints?apiVersion=1.0", nil), nil, &result)
	return resp, result, err
}

type ListEndpointsResponse []Endpoint
