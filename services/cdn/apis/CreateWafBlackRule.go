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

type CreateWafBlackRuleRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 黑名单类型， uri ip geo (Optional) */
    RuleType *string `json:"ruleType"`

    /* 匹配模式,uri类型有效，0=完全匹配  1=前缀匹配 2=包含 3=正则 4=大于 5=后缀 (Optional) */
    MatchOp *int `json:"matchOp"`

    /* 匹配值 (Optional) */
    Val *string `json:"val"`

    /* 1：forbidden，493封禁并返回自定义页面 2：redirect，302跳转 3： verify@captcha 4： verify@jscookie (Optional) */
    AtOp *int `json:"atOp"`

    /* action为1时为自定义页面名称,空值或缺省值default为默认页面，2时为跳转url，其他时无效 (Optional) */
    AtVal *string `json:"atVal"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateWafBlackRuleRequest(
    domain string,
) *CreateWafBlackRuleRequest {

	return &CreateWafBlackRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/wafBlackRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param ruleType: 黑名单类型， uri ip geo (Optional)
 * param matchOp: 匹配模式,uri类型有效，0=完全匹配  1=前缀匹配 2=包含 3=正则 4=大于 5=后缀 (Optional)
 * param val: 匹配值 (Optional)
 * param atOp: 1：forbidden，493封禁并返回自定义页面 2：redirect，302跳转 3： verify@captcha 4： verify@jscookie (Optional)
 * param atVal: action为1时为自定义页面名称,空值或缺省值default为默认页面，2时为跳转url，其他时无效 (Optional)
 */
func NewCreateWafBlackRuleRequestWithAllParams(
    domain string,
    ruleType *string,
    matchOp *int,
    val *string,
    atOp *int,
    atVal *string,
) *CreateWafBlackRuleRequest {

    return &CreateWafBlackRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        RuleType: ruleType,
        MatchOp: matchOp,
        Val: val,
        AtOp: atOp,
        AtVal: atVal,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateWafBlackRuleRequestWithoutParam() *CreateWafBlackRuleRequest {

    return &CreateWafBlackRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafBlackRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *CreateWafBlackRuleRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param ruleType: 黑名单类型， uri ip geo(Optional) */
func (r *CreateWafBlackRuleRequest) SetRuleType(ruleType string) {
    r.RuleType = &ruleType
}

/* param matchOp: 匹配模式,uri类型有效，0=完全匹配  1=前缀匹配 2=包含 3=正则 4=大于 5=后缀(Optional) */
func (r *CreateWafBlackRuleRequest) SetMatchOp(matchOp int) {
    r.MatchOp = &matchOp
}

/* param val: 匹配值(Optional) */
func (r *CreateWafBlackRuleRequest) SetVal(val string) {
    r.Val = &val
}

/* param atOp: 1：forbidden，493封禁并返回自定义页面 2：redirect，302跳转 3： verify@captcha 4： verify@jscookie(Optional) */
func (r *CreateWafBlackRuleRequest) SetAtOp(atOp int) {
    r.AtOp = &atOp
}

/* param atVal: action为1时为自定义页面名称,空值或缺省值default为默认页面，2时为跳转url，其他时无效(Optional) */
func (r *CreateWafBlackRuleRequest) SetAtVal(atVal string) {
    r.AtVal = &atVal
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateWafBlackRuleRequest) GetRegionId() string {
    return ""
}

type CreateWafBlackRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateWafBlackRuleResult `json:"result"`
}

type CreateWafBlackRuleResult struct {
}