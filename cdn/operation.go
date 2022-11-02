package cdn

import (
	"fmt"
	"net/http"
)

// Get operation information
//
// https://docs.azure.cn/en-us/cdn/cdn-api-get-operation
func (c *Client) GetOperation(req *GetOperationRequest) (resp *http.Response, result *GetOperationResponse, err error) {
	resp, err = c.Request(http.MethodGet, c.MakeRequestUrl(
		fmt.Sprintf("/endpoints/%s/operations/%s?apiVersion=1.0", req.EndpointID, req.OperationID), nil), nil, &result)
	return resp, result, err
}

type GetOperationRequest struct {
	EndpointID  string //Target node unique identifier
	OperationID string //Operation unique identifier
}

// Task status
type TaskStatus string

const (
	TaskStatusNotSet     TaskStatus = "NotSet"     //State not set
	TaskStatusProcessing TaskStatus = "Processing" //Currently processing
	TaskStatusSucceeded  TaskStatus = "Succeeded"  //Succeeded
	TaskStatusFailed     TaskStatus = "Failed"     //Failed
)

type GetOperationResponse struct {
	ID             string
	Type           string      //Operation type
	Status         TaskStatus  //Task status
	Message        interface{} //Operation information
	Start          string      //Start time
	End            string      //End time
	EndpointID     string      //Node unique identifier
	SubscriptionID string      //Subscription unique identifier
}
