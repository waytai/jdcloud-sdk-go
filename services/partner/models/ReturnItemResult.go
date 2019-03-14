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


type ReturnItemResult struct {

    /* ID (Optional) */
    Id int `json:"id"`

    /* 项目编码 (Optional) */
    ItemId int `json:"itemId"`

    /* 项目名称 (Optional) */
    ItemName string `json:"itemName"`

    /* 部门ID (Optional) */
    DeptId int `json:"deptId"`

    /* 部门名称 (Optional) */
    DeptName string `json:"deptName"`

    /* 渠道商类型 (Optional) */
    DistributorType int `json:"distributorType"`

    /* 渠道商类型名称 (Optional) */
    DistributorTypeName string `json:"distributorTypeName"`

    /* 返还类型 (Optional) */
    ReturnType int `json:"returnType"`

    /* 项目条件 (Optional) */
    ConditionItem string `json:"conditionItem"`

    /* 项目条件名称 (Optional) */
    ConditionItemName string `json:"conditionItemName"`

    /* 周期类型 (Optional) */
    CircleType int `json:"circleType"`

    /* 周期名称 (Optional) */
    CircleName string `json:"circleName"`

    /* 说明 (Optional) */
    Remark string `json:"remark"`

    /* 使用标识（1在用，0未使用） (Optional) */
    UseFlag string `json:"useFlag"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 创建人 (Optional) */
    CreateUser string `json:"createUser"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 修改人 (Optional) */
    UpdateUser string `json:"updateUser"`

    /* 是否删除0未删除,1已删除 (Optional) */
    Yn int `json:"yn"`
}
