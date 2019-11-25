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
    apigateway "github.com/jdcloud-api/jdcloud-sdk-go/services/apigateway/models"
)

type QueryKeysRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 页码, 默认为1, 取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20，取值范围：[10,100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 排序类型 (Optional) */
    OrderBy *string `json:"orderBy"`

    /* 用户类型 (Optional) */
    UserType *string `json:"userType"`

    /* 密钥Id (Optional) */
    KeyId *string `json:"keyId"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryKeysRequest(
    regionId string,
) *QueryKeysRequest {

	return &QueryKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/kms",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pageNumber: 页码, 默认为1, 取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小，默认为20，取值范围：[10,100] (Optional)
 * param orderBy: 排序类型 (Optional)
 * param userType: 用户类型 (Optional)
 * param keyId: 密钥Id (Optional)
 */
func NewQueryKeysRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    orderBy *string,
    userType *string,
    keyId *string,
) *QueryKeysRequest {

    return &QueryKeysRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        OrderBy: orderBy,
        UserType: userType,
        KeyId: keyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryKeysRequestWithoutParam() *QueryKeysRequest {

    return &QueryKeysRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/kms",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *QueryKeysRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码, 默认为1, 取值范围：[1,∞)(Optional) */
func (r *QueryKeysRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20，取值范围：[10,100](Optional) */
func (r *QueryKeysRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param orderBy: 排序类型(Optional) */
func (r *QueryKeysRequest) SetOrderBy(orderBy string) {
    r.OrderBy = &orderBy
}

/* param userType: 用户类型(Optional) */
func (r *QueryKeysRequest) SetUserType(userType string) {
    r.UserType = &userType
}

/* param keyId: 密钥Id(Optional) */
func (r *QueryKeysRequest) SetKeyId(keyId string) {
    r.KeyId = &keyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryKeysRequest) GetRegionId() string {
    return r.RegionId
}

type QueryKeysResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryKeysResult `json:"result"`
}

type QueryKeysResult struct {
    Total int `json:"total"`
    Data []apigateway.KeyInfo `json:"data"`
}