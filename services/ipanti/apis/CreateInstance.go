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
    ipanti "github.com/jdcloud-api/jdcloud-sdk-go/services/ipanti/models"
)

type CreateInstanceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 新购或升级实例请求参数  */
    CreateInstanceSpec *ipanti.CreateInstanceSpec `json:"createInstanceSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param createInstanceSpec: 新购或升级实例请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateInstanceRequest(
    regionId string,
    createInstanceSpec *ipanti.CreateInstanceSpec,
) *CreateInstanceRequest {

	return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CreateInstanceSpec: createInstanceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param createInstanceSpec: 新购或升级实例请求参数 (Required)
 */
func NewCreateInstanceRequestWithAllParams(
    regionId string,
    createInstanceSpec *ipanti.CreateInstanceSpec,
) *CreateInstanceRequest {

    return &CreateInstanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CreateInstanceSpec: createInstanceSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateInstanceRequestWithoutParam() *CreateInstanceRequest {

    return &CreateInstanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateInstanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param createInstanceSpec: 新购或升级实例请求参数(Required) */
func (r *CreateInstanceRequest) SetCreateInstanceSpec(createInstanceSpec *ipanti.CreateInstanceSpec) {
    r.CreateInstanceSpec = createInstanceSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateInstanceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateInstanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateInstanceResult `json:"result"`
}

type CreateInstanceResult struct {
    Code int `json:"code"`
    Message string `json:"message"`
}