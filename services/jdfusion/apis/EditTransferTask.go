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
    jdfusion "github.com/jdcloud-api/jdcloud-sdk-go/services/jdfusion/models"
)

type EditTransferTaskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 任务ID  */
    Id string `json:"id"`

    /*  (Optional) */
    Task *jdfusion.TransferTaskInfo `json:"task"`
}

/*
 * param regionId: 地域ID (Required)
 * param id: 任务ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEditTransferTaskRequest(
    regionId string,
    id string,
) *EditTransferTaskRequest {

	return &EditTransferTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/oss_transferTasks/{id}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Id: id,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param id: 任务ID (Required)
 * param task:  (Optional)
 */
func NewEditTransferTaskRequestWithAllParams(
    regionId string,
    id string,
    task *jdfusion.TransferTaskInfo,
) *EditTransferTaskRequest {

    return &EditTransferTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_transferTasks/{id}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Task: task,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEditTransferTaskRequestWithoutParam() *EditTransferTaskRequest {

    return &EditTransferTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/oss_transferTasks/{id}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *EditTransferTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 任务ID(Required) */
func (r *EditTransferTaskRequest) SetId(id string) {
    r.Id = id
}

/* param task: (Optional) */
func (r *EditTransferTaskRequest) SetTask(task *jdfusion.TransferTaskInfo) {
    r.Task = task
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EditTransferTaskRequest) GetRegionId() string {
    return r.RegionId
}

type EditTransferTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EditTransferTaskResult `json:"result"`
}

type EditTransferTaskResult struct {
}