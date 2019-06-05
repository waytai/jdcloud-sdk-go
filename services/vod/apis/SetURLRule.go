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

type SetURLRuleRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`

    /* 规则类型，取值 'url' (Optional) */
    RuleType *string `json:"ruleType"`

    /* 规则配置对象 (Optional) */
    Config *interface{} `json:"config"`

    /* 是否启用该规则 (Optional) */
    Enabled *bool `json:"enabled"`
}

/*
 * param domainId: 域名ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetURLRuleRequest(
    domainId int,
) *SetURLRuleRequest {

	return &SetURLRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:setURLRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 * param ruleType: 规则类型，取值 'url' (Optional)
 * param config: 规则配置对象 (Optional)
 * param enabled: 是否启用该规则 (Optional)
 */
func NewSetURLRuleRequestWithAllParams(
    domainId int,
    ruleType *string,
    config *interface{},
    enabled *bool,
) *SetURLRuleRequest {

    return &SetURLRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setURLRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        DomainId: domainId,
        RuleType: ruleType,
        Config: config,
        Enabled: enabled,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetURLRuleRequestWithoutParam() *SetURLRuleRequest {

    return &SetURLRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setURLRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *SetURLRuleRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param ruleType: 规则类型，取值 'url'(Optional) */
func (r *SetURLRuleRequest) SetRuleType(ruleType string) {
    r.RuleType = &ruleType
}

/* param config: 规则配置对象(Optional) */
func (r *SetURLRuleRequest) SetConfig(config interface{}) {
    r.Config = &config
}

/* param enabled: 是否启用该规则(Optional) */
func (r *SetURLRuleRequest) SetEnabled(enabled bool) {
    r.Enabled = &enabled
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetURLRuleRequest) GetRegionId() string {
    return ""
}

type SetURLRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetURLRuleResult `json:"result"`
}

type SetURLRuleResult struct {
}