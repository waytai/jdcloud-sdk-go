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


type SnowLeopardIpSourceInfo struct {

    /* 1：主；2：备 (Optional) */
    Master int `json:"master"`

    /* 回源IP (Optional) */
    Ip string `json:"ip"`

    /* 占比 (Optional) */
    Ratio float64 `json:"ratio"`

    /* 运营商 (Optional) */
    Isp string `json:"isp"`
}
