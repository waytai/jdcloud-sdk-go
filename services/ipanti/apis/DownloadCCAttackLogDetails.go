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

type DownloadCCAttackLogDetailsRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 开始时间, 只能下载最近 60 天以内的数据, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ  */
    StartTime string `json:"startTime"`

    /* 查询的结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ  */
    EndTime string `json:"endTime"`

    /* 高防实例 ID  */
    InstanceId int `json:"instanceId"`

    /* 子域名 (Optional) */
    SubDomain []string `json:"subDomain"`
}

/*
 * param regionId: 区域 Id (Required)
 * param startTime: 开始时间, 只能下载最近 60 天以内的数据, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param instanceId: 高防实例 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDownloadCCAttackLogDetailsRequest(
    regionId string,
    startTime string,
    endTime string,
    instanceId int,
) *DownloadCCAttackLogDetailsRequest {

	return &DownloadCCAttackLogDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/attacklog:CCDetail/download",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param startTime: 开始时间, 只能下载最近 60 天以内的数据, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param endTime: 查询的结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ (Required)
 * param instanceId: 高防实例 ID (Required)
 * param subDomain: 子域名 (Optional)
 */
func NewDownloadCCAttackLogDetailsRequestWithAllParams(
    regionId string,
    startTime string,
    endTime string,
    instanceId int,
    subDomain []string,
) *DownloadCCAttackLogDetailsRequest {

    return &DownloadCCAttackLogDetailsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/attacklog:CCDetail/download",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        StartTime: startTime,
        EndTime: endTime,
        InstanceId: instanceId,
        SubDomain: subDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDownloadCCAttackLogDetailsRequestWithoutParam() *DownloadCCAttackLogDetailsRequest {

    return &DownloadCCAttackLogDetailsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/attacklog:CCDetail/download",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *DownloadCCAttackLogDetailsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param startTime: 开始时间, 只能下载最近 60 天以内的数据, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DownloadCCAttackLogDetailsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询的结束时间, UTC 时间, 格式：yyyy-MM-dd'T'HH:mm:ssZ(Required) */
func (r *DownloadCCAttackLogDetailsRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param instanceId: 高防实例 ID(Required) */
func (r *DownloadCCAttackLogDetailsRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param subDomain: 子域名(Optional) */
func (r *DownloadCCAttackLogDetailsRequest) SetSubDomain(subDomain []string) {
    r.SubDomain = subDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DownloadCCAttackLogDetailsRequest) GetRegionId() string {
    return r.RegionId
}

type DownloadCCAttackLogDetailsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DownloadCCAttackLogDetailsResult `json:"result"`
}

type DownloadCCAttackLogDetailsResult struct {
}