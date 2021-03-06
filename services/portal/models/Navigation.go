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


type Navigation struct {

    /* 主键id (Optional) */
    Id int `json:"id"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 图标地址 (Optional) */
    IconUrl string `json:"iconUrl"`

    /* ICON 样式 (Optional) */
    IconClass string `json:"iconClass"`

    /* 链接地址 (Optional) */
    WebUrl string `json:"webUrl"`

    /* url：用于查询产品详情 (Optional) */
    Url string `json:"url"`

    /* 是否为京东云产品；0:是京东云产品；1:不是京东云产品 (Optional) */
    ProductStatus int `json:"productStatus"`

    /* 排序 (Optional) */
    Sort int `json:"sort"`

    /* 修改时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 父ID (Optional) */
    ParentId int `json:"parentId"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 导航层级 (Optional) */
    Level int `json:"level"`

    /* 标签 (Optional) */
    Label string `json:"label"`

    /* 帮助文档地址 (Optional) */
    HelpUrl string `json:"helpUrl"`

    /* 自营标签 (Optional) */
    SelfRun string `json:"selfRun"`

    /* 语言：中文cn；英文en (Optional) */
    Lang string `json:"lang"`

    /* 子结构 (Optional) */
    ExtChildren []Navigation `json:"extChildren"`
}
