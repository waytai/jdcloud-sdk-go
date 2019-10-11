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


type SharedUsers struct {

    /* 用户Pin信息 (Optional) */
    UserPin int `json:"userPin"`

    /* 用户对应的资源编号 (Optional) */
    SourceId int `json:"sourceId"`

    /* 用户所属region (Optional) */
    RegionName int `json:"regionName"`

    /* 用户开通时间 (Optional) */
    CreateTime int `json:"createTime"`

    /* 租户状态[1-正常使用，2-欠费停服，3-软删除保护期] (Optional) */
    TenantStatus int `json:"tenantStatus"`

    /* 在线设备数 (Optional) */
    OnDevices int `json:"onDevices"`

    /* 日消息条数 (Optional) */
    DailyMessages int `json:"dailyMessages"`

    /* 总消息条数 (Optional) */
    TotalMessages int `json:"totalMessages"`
}
