package cdn

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Get bandwidth information
// https://docs.azure.cn/en-us/cdn/cdn-api-get-endpoint-bandwidth
func (c *Client) GetEndpointBandwidth(req *GetEndpointBandwidthRequest) (resp *http.Response, result *GetEndpointBandwidthResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/bandwidth?apiVersion=1.0", req.EndpointId), url.Values{
		"startTime": {req.StartTime.UTC().Format("2006-01-02T15:04:05Z")},
		"endTime":   {req.EndTime.UTC().Format("2006-01-02T15:04:05Z")},
	})
	resp, err = c.Request(http.MethodGet, reqUrl, nil, &result)
	return resp, result, err
}

type GetEndpointBandwidthRequest struct {
	EndpointId string    //Target node unique identifier
	StartTime  time.Time //The bandwidth query start time must be a UTC time in the 'yyyy-MM-ddThh:mm:ssZ’ format.
	EndTime    time.Time //The bandwidth query end time must be a UTC time in the 'yyyy-MM-ddThh:mm:ssZ’ format.
}

type GetEndpointBandwidthResponse struct {
	DomainName string
	Items      []struct {
		Timestamp             string
		BandwidthInMbps       int64
		OriginBandwidthInMbps int64
	}
	PeakBandwidthInMbps         int64 //CDN bandwidth peak value
	ValleyBandwidthInMbps       int64 //CDN bandwidth trough value
	PeakOriginBandwidthInMbps   int64 //Return-to-source bandwidth peak value
	ValleyOriginBandwidthInMbps int64 //eturn-to-source bandwidth trough value
}

// Get traffic information
// https://docs.azure.cn/en-us/cdn/cdn-api-get-endpoint-volume
func (c *Client) GetEndpointVolume(req *GetEndpointVolumeRequest) (resp *http.Response, result *GetEndpointVolumeResponse, err error) {
	reqUrl := c.MakeRequestUrl(fmt.Sprintf("/endpoints/%s/volume?apiVersion=1.0", req.EndpointID), url.Values{
		"granularity": {req.Granularity},
		"startTime":   {req.StartTime.Format("2006-01-02T15:04:05Z")},
		"endTime":     {req.EndTime.Format("2006-01-02T15:04:05Z")},
	})

	resp, err = c.Request(http.MethodGet, reqUrl, nil, &result)
	return resp, result, err
}

// Traffic statistic granularity
type Granularity string

const (
	GranularityPerFiveMinutes Granularity = "PerFiveMinutes" //Per five minutes
	GranularityPerHour        Granularity = "PerHour"        //Per hour
	GranularityPerDay         Granularity = "PerDay"         //Per day
)

type GetEndpointVolumeRequest struct {
	EndpointID  string    //Target node unique identifier
	Granularity string    //Traffic statistic granularity
	StartTime   time.Time //The traffic query start time must be a UTC time in the 'yyyy-MM-ddThh:mm:ssZ’ format.
	EndTime     time.Time //The traffic query end time must be a UTC time in the 'yyyy-MM-ddThh:mm:ssZ’ format.
}

type GetEndpointVolumeResponse struct {
	DomainName string //Accelerated domain names
	Items      []struct {
		Timestamp        string
		VolumeInMB       int64 //CDN traffic
		OriginVolumeInMB int64 //Back to source traffic
	}
	TotalCDNVolumeInMB    int64 //CDN total traffic
	TotalOriginVolumeInMB int64 //Back to source total traffic
}
