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

type SetIPRuleRequest struct {

    core.JDCloudRequest

    /* 域名ID  */
    DomainId int `json:"domainId"`

    /* 规则类型，取值 'ip' (Optional) */
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
func NewSetIPRuleRequest(
    domainId int,
) *SetIPRuleRequest {

	return &SetIPRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domains/{domainId}:setIPRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        DomainId: domainId,
	}
}

/*
 * param domainId: 域名ID (Required)
 * param ruleType: 规则类型，取值 'ip' (Optional)
 * param config: 规则配置对象 (Optional)
 * param enabled: 是否启用该规则 (Optional)
 */
func NewSetIPRuleRequestWithAllParams(
    domainId int,
    ruleType *string,
    config *interface{},
    enabled *bool,
) *SetIPRuleRequest {

    return &SetIPRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setIPRule",
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
func NewSetIPRuleRequestWithoutParam() *SetIPRuleRequest {

    return &SetIPRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domains/{domainId}:setIPRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domainId: 域名ID(Required) */
func (r *SetIPRuleRequest) SetDomainId(domainId int) {
    r.DomainId = domainId
}

/* param ruleType: 规则类型，取值 'ip'(Optional) */
func (r *SetIPRuleRequest) SetRuleType(ruleType string) {
    r.RuleType = &ruleType
}

/* param config: 规则配置对象(Optional) */
func (r *SetIPRuleRequest) SetConfig(config interface{}) {
    r.Config = &config
}

/* param enabled: 是否启用该规则(Optional) */
func (r *SetIPRuleRequest) SetEnabled(enabled bool) {
    r.Enabled = &enabled
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetIPRuleRequest) GetRegionId() string {
    return ""
}

type SetIPRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetIPRuleResult `json:"result"`
}

type SetIPRuleResult struct {
}