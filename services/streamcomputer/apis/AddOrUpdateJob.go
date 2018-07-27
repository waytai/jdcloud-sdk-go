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
    streamcomputer "github.com/jdcloud-api/jdcloud-sdk-go/services/streamcomputer/models"
)

type AddOrUpdateJobRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 创建作业的详情  */
    JobStr *streamcomputer.JobStr `json:"jobStr"`
}

/*
 * param regionId: Region ID (Required)
 * param jobStr: 创建作业的详情 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddOrUpdateJobRequest(
    regionId string,
    jobStr *streamcomputer.JobStr,
) *AddOrUpdateJobRequest {

	return &AddOrUpdateJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/job",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        JobStr: jobStr,
	}
}

/*
 * param regionId: Region ID (Required)
 * param jobStr: 创建作业的详情 (Required)
 */
func NewAddOrUpdateJobRequestWithAllParams(
    regionId string,
    jobStr *streamcomputer.JobStr,
) *AddOrUpdateJobRequest {

    return &AddOrUpdateJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/job",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        JobStr: jobStr,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddOrUpdateJobRequestWithoutParam() *AddOrUpdateJobRequest {

    return &AddOrUpdateJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/job",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddOrUpdateJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param jobStr: 创建作业的详情(Required) */
func (r *AddOrUpdateJobRequest) SetJobStr(jobStr *streamcomputer.JobStr) {
    r.JobStr = jobStr
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddOrUpdateJobRequest) GetRegionId() string {
    return r.RegionId
}

type AddOrUpdateJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddOrUpdateJobResult `json:"result"`
}

type AddOrUpdateJobResult struct {
    OkInfo streamcomputer.OkInfo `json:"okInfo"`
}