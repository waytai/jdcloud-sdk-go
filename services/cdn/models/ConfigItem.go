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


type ConfigItem struct {

    /* 配置项类型 (Optional) */
    ConfigItemType string `json:"configItemType"`

    /* 配置项名称 (Optional) */
    ConfigItemName string `json:"configItemName"`

    /* 配置状态 (Optional) */
    ConfigStatus int `json:"configStatus"`

    /* 配置状态名 (Optional) */
    ConfigStatusName string `json:"configStatusName"`

    /* 配置项细节,类型为Map<String,Object> (Optional) */
    ConfigItemDetails interface{} `json:"configItemDetails"`
}
