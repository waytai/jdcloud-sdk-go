// Copyright 2018-2025 JDCLOUD.COM
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
    . "github.com/jdcloud-api/jdcloud-sdk-go/core"
    "reflect"
)

type CreateDatabaseRequest struct {

    JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 数据库名  */
    DatabaseName string `json:"databaseName"`

    /* 实例名称  */
    InstanceName string `json:"instanceName"`

    /* 描述信息 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: 地域ID 
 * param databaseName: 数据库名 
 * param instanceName: 实例名称 
 * param description: 描述信息 (Optional)
 */
func NewCreateDatabaseRequest(
    regionId string,
    databaseName string,
    instanceName string,
) *CreateDatabaseRequest {

	return &CreateDatabaseRequest{
        JDCloudRequest: JDCloudRequest{
			URL:     "/regions/{regionId}/dwDatabase/{databaseName}",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DatabaseName: databaseName,
        InstanceName: instanceName,
	}
}

func (r *CreateDatabaseRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *CreateDatabaseRequest) SetDatabaseName(databaseName string) {
    r.DatabaseName = databaseName
}

func (r *CreateDatabaseRequest) SetInstanceName(instanceName string) {
    r.InstanceName = instanceName
}

func (r *CreateDatabaseRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateDatabaseRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type CreateDatabaseResponse struct {
    RequestID string `json:"requestId"`
    Error ErrorResponse `json:"error"`
    Result CreateDatabaseResult `json:"result"`
}

type CreateDatabaseResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}