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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    iotlink "github.com/jdcloud-api/jdcloud-sdk-go/services/iotlink/apis"
    "encoding/json"
    "errors"
)

type IotlinkClient struct {
    core.JDCloudClient
}

func NewIotlinkClient(credential *core.Credential) *IotlinkClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("iotlink.jdcloud-api.com")

    return &IotlinkClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "iotlink",
            Revision:    "1.0.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *IotlinkClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *IotlinkClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *IotlinkClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 根据物联网卡iccid查询该卡的开关机状态信息 */
func (c *IotlinkClient) OnOffStatus(request *iotlink.OnOffStatusRequest) (*iotlink.OnOffStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.OnOffStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 物联网卡开机操作 */
func (c *IotlinkClient) OpenIotCard(request *iotlink.OpenIotCardRequest) (*iotlink.OpenIotCardResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.OpenIotCardResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 物联网卡停机操作 */
func (c *IotlinkClient) CloseIotCard(request *iotlink.CloseIotCardRequest) (*iotlink.CloseIotCardResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.CloseIotCardResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据物联网卡iccid查询该卡的生命周期信息 */
func (c *IotlinkClient) LifeStatus(request *iotlink.LifeStatusRequest) (*iotlink.LifeStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.LifeStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 物联网卡停流量操作 */
func (c *IotlinkClient) CloseIotFlow(request *iotlink.CloseIotFlowRequest) (*iotlink.CloseIotFlowResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.CloseIotFlowResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据物联网卡iccid查询该卡的gprs状态信息 */
func (c *IotlinkClient) GprsStatus(request *iotlink.GprsStatusRequest) (*iotlink.GprsStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.GprsStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 物联网卡开启流量操作 */
func (c *IotlinkClient) OpenIotFlow(request *iotlink.OpenIotFlowRequest) (*iotlink.OpenIotFlowResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.OpenIotFlowResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据物联网卡iccid查询该卡的当月套餐内的GPRS实时使用量 */
func (c *IotlinkClient) GprsRealtimeInfo(request *iotlink.GprsRealtimeInfoRequest) (*iotlink.GprsRealtimeInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotlink.GprsRealtimeInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

