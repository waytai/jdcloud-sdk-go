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

package models


type AttackDetail struct {

    /* 请求id (Optional) */
    RequestId string `json:"requestId"`

    /* 攻击来源ip (Optional) */
    Ip string `json:"ip"`

    /* 攻击来源地域 (Optional) */
    Area string `json:"area"`

    /* 攻击流量(MB) (Optional) */
    Flow string `json:"flow"`

    /* 攻击时间 (Optional) */
    TimeUtc string `json:"timeUtc"`

    /* 攻击方法（post，get等） (Optional) */
    Method string `json:"method"`

    /* 攻击url (Optional) */
    Url string `json:"url"`

    /* 攻击类型 (Optional) */
    AttackType string `json:"attackType"`

    /* 针对该攻击做出的动作 (Optional) */
    Action string `json:"action"`
}