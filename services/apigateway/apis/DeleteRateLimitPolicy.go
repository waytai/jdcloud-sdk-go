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

type DeleteRateLimitPolicyRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 限流策略ID  */
    PolicyId string `json:"policyId"`
}

/*
 * param regionId: 地域ID (Required)
 * param policyId: 限流策略ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteRateLimitPolicyRequest(
    regionId string,
    policyId string,
) *DeleteRateLimitPolicyRequest {

	return &DeleteRateLimitPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/rateLimitPolicies/{policyId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PolicyId: policyId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param policyId: 限流策略ID (Required)
 */
func NewDeleteRateLimitPolicyRequestWithAllParams(
    regionId string,
    policyId string,
) *DeleteRateLimitPolicyRequest {

    return &DeleteRateLimitPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies/{policyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PolicyId: policyId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteRateLimitPolicyRequestWithoutParam() *DeleteRateLimitPolicyRequest {

    return &DeleteRateLimitPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/rateLimitPolicies/{policyId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DeleteRateLimitPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param policyId: 限流策略ID(Required) */
func (r *DeleteRateLimitPolicyRequest) SetPolicyId(policyId string) {
    r.PolicyId = policyId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteRateLimitPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteRateLimitPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteRateLimitPolicyResult `json:"result"`
}

type DeleteRateLimitPolicyResult struct {
    PolicyId string `json:"policyId"`
}