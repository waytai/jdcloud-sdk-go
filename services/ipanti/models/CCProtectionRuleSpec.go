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


type CCProtectionRuleSpec struct {

    /* CC 防护规则名称, 不允许为空, 长度不超过 32 个字符, 支持中文, 大小写字母, 数字及字符'-'、'/'、'.'、'_'  */
    Name string `json:"name"`

    /* uri, 不允许为空, 以 / 开头, 长度不超过 2048 个字符  */
    Uri string `json:"uri"`

    /* 匹配 uri 类型, 0: 精确匹配, 1: 前缀匹配  */
    MatchType int `json:"matchType"`

    /* 检测周期, 单位为秒, 取值范围[5, 10800]  */
    DetectPeriod int64 `json:"detectPeriod"`

    /* ip 访问次数, 取值范围[2, 2000]  */
    SingleIpLimit int64 `json:"singleIpLimit"`

    /* 阻断类型, 1: 封禁, 2: 人机交互  */
    BlockType int `json:"blockType"`

    /* 阻断持续时间, 单位为分钟, 取值范围[1, 1440]  */
    BlockTime int64 `json:"blockTime"`
}
