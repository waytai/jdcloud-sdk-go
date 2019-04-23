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


type ContainerState struct {

    /* 容器running的详细信息 (Optional) */
    Running ContainerStateRunning `json:"running"`

    /* 容器Terminated的详细信息 (Optional) */
    Terminated ContainerStateTerminated `json:"terminated"`

    /* 容器waiting的详细信息 (Optional) */
    Waiting ContainerStateWaiting `json:"waiting"`
}