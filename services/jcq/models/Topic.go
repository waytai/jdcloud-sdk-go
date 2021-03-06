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


type Topic struct {

    /* topic Id (Optional) */
    TopicId string `json:"topicId"`

    /* topic名称 (Optional) */
    TopicName string `json:"topicName"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新时间 (Optional) */
    LastUpdateTime string `json:"lastUpdateTime"`

    /* topicStatus (Optional) */
    TopicStatus string `json:"topicStatus"`

    /* 自己创建的订阅数 (Optional) */
    SubscriptionCount int `json:"subscriptionCount"`

    /* 消息生命周期时长小时 (Optional) */
    MessageLifeTimeInHours int `json:"messageLifeTimeInHours"`

    /* 配置信息 (Optional) */
    TopicConfig TopicConfig `json:"topicConfig"`

    /* 是否是自己的topic (Optional) */
    Own bool `json:"own"`

    /* 被授权的权限[PUB,SUB,PUBSUB,READ_ONLY,ADMIN] (Optional) */
    AuthorizedPermission string `json:"authorizedPermission"`

    /* 标签信息 (Optional) */
    Tags []Tag `json:"tags"`
}
