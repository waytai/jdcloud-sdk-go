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
    billing "github.com/jdcloud-api/jdcloud-sdk-go/services/billing/models"
)

type QueryBillSummaryRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 计费开始时间  */
    StartTime string `json:"startTime"`

    /* 计费结束时间  */
    EndTime string `json:"endTime"`

    /* 产品线代码 (Optional) */
    AppCode *string `json:"appCode"`

    /* 产品代码 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源单id列表 (Optional) */
    ResourceIds []string `json:"resourceIds"`

    /* 标签 (Optional) */
    Tags []interface{} `json:"tags"`

    /* pageIndex (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* pageSize (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param startTime: 计费开始时间 (Required)
 * param endTime: 计费结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryBillSummaryRequest(
    regionId string,
    startTime string,
    endTime string,
) *QueryBillSummaryRequest {

	return &QueryBillSummaryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/billSummary:list",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: Region ID (Required)
 * param startTime: 计费开始时间 (Required)
 * param endTime: 计费结束时间 (Required)
 * param appCode: 产品线代码 (Optional)
 * param serviceCode: 产品代码 (Optional)
 * param resourceIds: 资源单id列表 (Optional)
 * param tags: 标签 (Optional)
 * param pageIndex: pageIndex (Optional)
 * param pageSize: pageSize (Optional)
 */
func NewQueryBillSummaryRequestWithAllParams(
    regionId string,
    startTime string,
    endTime string,
    appCode *string,
    serviceCode *string,
    resourceIds []string,
    tags []interface{},
    pageIndex *int,
    pageSize *int,
) *QueryBillSummaryRequest {

    return &QueryBillSummaryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billSummary:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        AppCode: appCode,
        ServiceCode: serviceCode,
        ResourceIds: resourceIds,
        Tags: tags,
        PageIndex: pageIndex,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryBillSummaryRequestWithoutParam() *QueryBillSummaryRequest {

    return &QueryBillSummaryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billSummary:list",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryBillSummaryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 计费开始时间(Required) */
func (r *QueryBillSummaryRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 计费结束时间(Required) */
func (r *QueryBillSummaryRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param appCode: 产品线代码(Optional) */
func (r *QueryBillSummaryRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: 产品代码(Optional) */
func (r *QueryBillSummaryRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceIds: 资源单id列表(Optional) */
func (r *QueryBillSummaryRequest) SetResourceIds(resourceIds []string) {
    r.ResourceIds = resourceIds
}

/* param tags: 标签(Optional) */
func (r *QueryBillSummaryRequest) SetTags(tags []interface{}) {
    r.Tags = tags
}

/* param pageIndex: pageIndex(Optional) */
func (r *QueryBillSummaryRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: pageSize(Optional) */
func (r *QueryBillSummaryRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryBillSummaryRequest) GetRegionId() string {
    return r.RegionId
}

type QueryBillSummaryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryBillSummaryResult `json:"result"`
}

type QueryBillSummaryResult struct {
    Pagination billing.Pagination `json:"pagination"`
    Result []billing.BillSummary `json:"result"`
}