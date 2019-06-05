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


type JobRequest struct {

    /* 项目名称  */
    Name string `json:"name"`

    /* 源提供商，代码存储类型代码存储类型， 目前只支持github  */
    CodeType string `json:"codeType"`

    /* 仓库地址，代码clone路径  */
    CodeRepoUrl string `json:"codeRepoUrl"`

    /* 代码分支  */
    CodeRepoBranch string `json:"codeRepoBranch"`

    /* 获取用户OSS库用，用户云存储路径用户云存储路径，如果为空，使用公用的云存储 (Optional) */
    OssPath *string `json:"ossPath"`

    /* 上传区域，用户云存储主机，实际为用户云存储所在地域  */
    OssHost string `json:"ossHost"`

    /* 获取用户OSS库用，用户云存储bucket，如果为空，使用公用的云存储 (Optional) */
    OssBucket *string `json:"ossBucket"`

    /* 选择类型 和 运行版本 共同拼出此项，编译镜像地址  */
    BuildImage string `json:"buildImage"`

    /* 构建规范，选择在源代码供目录中使用build.yml则为false，选择插入构建命令则为true，这项为true，则buildSetConfig需要有内容，如果这项为false，即使buildSetConfig有内容，也不生效 (Optional) */
    IsUserBuildSetConfig *bool `json:"isUserBuildSetConfig"`

    /* 插入构建命令，isUserBuildSetConfig选择true时，这项让用户填写内容，内容从接口/regions/{regionId}/jobs/default/buildSet 获取 (Optional) */
    BuildSetConfig *string `json:"buildSetConfig"`

    /* 超时时间，单位秒 (Optional) */
    BuildTimeOut *int `json:"buildTimeOut"`

    /* 计算类型中 cpu分配核数 (Optional) */
    BuildResourceCpu *int `json:"buildResourceCpu"`

    /* 计算类型中 内存分配大小，单位MB (Optional) */
    BuildResourceMem *int `json:"buildResourceMem"`

    /* 通知邮件 (Optional) */
    NoticeMail *string `json:"noticeMail"`

    /* 通知频率， MAIL_FAILED失败时通知，MAIL_EVERY每次构建就通知 (Optional) */
    NoticeType *string `json:"noticeType"`

    /* 构建类型 (Optional) */
    CompilerType *string `json:"compilerType"`

    /* 镜像注册表名 (Optional) */
    DockerRegistry *string `json:"dockerRegistry"`

    /* 镜像仓库名 (Optional) */
    DockerRepository *string `json:"dockerRepository"`

    /* 注册表的URI (Optional) */
    DockerRegistryUri *string `json:"dockerRegistryUri"`
}