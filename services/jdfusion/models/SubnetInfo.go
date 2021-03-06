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


type SubnetInfo struct {

    /* Subnet的Id (Optional) */
    Id string `json:"id"`

    /* 子网所属VPC的Id (Optional) */
    VpcId string `json:"vpcId"`

    /* 子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 子网网段 (Optional) */
    CidrBlock string `json:"cidrBlock"`

    /* 子网可用ip数量 (Optional) */
    AvailableIpCount int `json:"availableIpCount"`

    /* 子网描述信息 (Optional) */
    Description string `json:"description"`

    /* 子网的结束地址 (Optional) */
    EndIp string `json:"endIp"`

    /* 子网关联的路由表Id (Optional) */
    RouteTableId string `json:"routeTableId"`

    /* 子网的起始地址 (Optional) */
    StartIp string `json:"startIp"`

    /* 所属云提供商ID (Optional) */
    CloudID string `json:"cloudID"`

    /* 可用区 (Optional) */
    Az string `json:"az"`

    /* 子网创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
