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


type MonitorIotCoreInstance struct {

    /* 实例Id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例名称 (Optional) */
    InstanceName string `json:"instanceName"`

    /* 创建时间 (Optional) */
    CreatedTime int64 `json:"createdTime"`

    /* 规格 (Optional) */
    Specifications string `json:"specifications"`

    /* 实例状态[0-创建中，1-运行中，2-停止] (Optional) */
    InstanceStatus string `json:"instanceStatus"`

    /* 描述 (Optional) */
    InstanceDescription string `json:"instanceDescription"`

    /* 资源到期时间 (Optional) */
    ExpiredTime int64 `json:"expiredTime"`

    /* 私有网络 (Optional) */
    VpcName string `json:"vpcName"`

    /* 子网 (Optional) */
    SubNetName string `json:"subNetName"`

    /* 内网域名 (Optional) */
    PriDomain string `json:"priDomain"`

    /* 公网域名 (Optional) */
    PubDomain string `json:"pubDomain"`

    /* 最大同时在线设备数 (Optional) */
    MaxOnlineDevices int64 `json:"maxOnlineDevices"`

    /* 每月最大消息数 (Optional) */
    MaxMonthMessages int64 `json:"maxMonthMessages"`
}
