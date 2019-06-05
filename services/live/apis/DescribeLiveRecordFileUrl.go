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

type DescribeLiveRecordFileUrlRequest struct {

    core.JDCloudRequest

    /* 录制文件ID
  */
    FileId string `json:"fileId"`

    /* 地址有效期，单位：秒，默认：3600，最大支持7天
 (Optional) */
    AuthExpire *int `json:"authExpire"`
}

/*
 * param fileId: 录制文件ID
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveRecordFileUrlRequest(
    fileId string,
) *DescribeLiveRecordFileUrlRequest {

	return &DescribeLiveRecordFileUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/recordFiles/{fileId}:url",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        FileId: fileId,
	}
}

/*
 * param fileId: 录制文件ID
 (Required)
 * param authExpire: 地址有效期，单位：秒，默认：3600，最大支持7天
 (Optional)
 */
func NewDescribeLiveRecordFileUrlRequestWithAllParams(
    fileId string,
    authExpire *int,
) *DescribeLiveRecordFileUrlRequest {

    return &DescribeLiveRecordFileUrlRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordFiles/{fileId}:url",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        FileId: fileId,
        AuthExpire: authExpire,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveRecordFileUrlRequestWithoutParam() *DescribeLiveRecordFileUrlRequest {

    return &DescribeLiveRecordFileUrlRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/recordFiles/{fileId}:url",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param fileId: 录制文件ID
(Required) */
func (r *DescribeLiveRecordFileUrlRequest) SetFileId(fileId string) {
    r.FileId = fileId
}

/* param authExpire: 地址有效期，单位：秒，默认：3600，最大支持7天
(Optional) */
func (r *DescribeLiveRecordFileUrlRequest) SetAuthExpire(authExpire int) {
    r.AuthExpire = &authExpire
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveRecordFileUrlRequest) GetRegionId() string {
    return ""
}

type DescribeLiveRecordFileUrlResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveRecordFileUrlResult `json:"result"`
}

type DescribeLiveRecordFileUrlResult struct {
    FileId string `json:"fileId"`
    FileUrl string `json:"fileUrl"`
}