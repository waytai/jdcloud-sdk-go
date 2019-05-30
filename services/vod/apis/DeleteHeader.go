// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DeleteHeaderRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`

    /* 头参数名 (Optional) */
    HeaderName *string `json:"headerName"`

    /* 头参数值 (Optional) */
    HeaderValue *string `json:"headerValue"`

    /* 头参数类型 (Optional) */
    HeaderType *string `json:"headerType"`
}

/*
 * param domainId: 域名ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteHeaderRequest(
    domainId int,
) *DeleteHeaderRequest {

	return &DeleteHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:deleteHeader",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 * param headerName: 头参数名 (Optional)
 * param headerValue: 头参数值 (Optional)
 * param headerType: 头参数类型 (Optional)
 */
func NewDeleteHeaderRequestWithAllParams(
    domainId int,
    headerName *string,
    headerValue *string,
    headerType *string,
) *DeleteHeaderRequest {

    return &DeleteHeaderRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:deleteHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
        HeaderName: headerName,
        HeaderValue: headerValue,
        HeaderType: headerType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteHeaderRequestWithoutParam() *DeleteHeaderRequest {

    return &DeleteHeaderRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:deleteHeader",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *DeleteHeaderRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param headerName: 头参数名(Optional) */
func (r *DeleteHeaderRequest) SetHeaderName(headerName string) {
    r.HeaderName = &headerName
}

/* param headerValue: 头参数值(Optional) */
func (r *DeleteHeaderRequest) SetHeaderValue(headerValue string) {
    r.HeaderValue = &headerValue
}

/* param headerType: 头参数类型(Optional) */
func (r *DeleteHeaderRequest) SetHeaderType(headerType string) {
    r.HeaderType = &headerType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteHeaderRequest) GetRegionId() string {
    return ""
}

type DeleteHeaderResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteHeaderResult `json:"result"`
}

type DeleteHeaderResult struct {
}