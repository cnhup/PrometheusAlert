package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBInstancePerformance invokes the rds.DescribeDBInstancePerformance API synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceperformance.html
func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (response *DescribeDBInstancePerformanceResponse, err error) {
	response = CreateDescribeDBInstancePerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancePerformanceWithChan invokes the rds.DescribeDBInstancePerformance API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancePerformanceWithChan(request *DescribeDBInstancePerformanceRequest) (<-chan *DescribeDBInstancePerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancePerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstancePerformance(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDBInstancePerformanceWithCallback invokes the rds.DescribeDBInstancePerformance API asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstanceperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstancePerformanceWithCallback(request *DescribeDBInstancePerformanceRequest, callback func(response *DescribeDBInstancePerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancePerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstancePerformance(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDBInstancePerformanceRequest is the request struct for api DescribeDBInstancePerformance
type DescribeDBInstancePerformanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Key                  string           `position:"Query" name:"Key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBInstancePerformanceResponse is the response struct for api DescribeDBInstancePerformance
type DescribeDBInstancePerformanceResponse struct {
	*responses.BaseResponse
	RequestId       string                                         `json:"RequestId" xml:"RequestId"`
	DBInstanceId    string                                         `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine          string                                         `json:"Engine" xml:"Engine"`
	StartTime       string                                         `json:"StartTime" xml:"StartTime"`
	EndTime         string                                         `json:"EndTime" xml:"EndTime"`
	PerformanceKeys PerformanceKeysInDescribeDBInstancePerformance `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

// CreateDescribeDBInstancePerformanceRequest creates a request to invoke DescribeDBInstancePerformance API
func CreateDescribeDBInstancePerformanceRequest() (request *DescribeDBInstancePerformanceRequest) {
	request = &DescribeDBInstancePerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancePerformance", "rds", "openAPI")
	return
}

// CreateDescribeDBInstancePerformanceResponse creates a response to parse from DescribeDBInstancePerformance response
func CreateDescribeDBInstancePerformanceResponse() (response *DescribeDBInstancePerformanceResponse) {
	response = &DescribeDBInstancePerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
