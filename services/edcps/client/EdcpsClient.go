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
    edcps "github.com/jdcloud-api/jdcloud-sdk-go/services/edcps/apis"
    "encoding/json"
    "errors"
)

type EdcpsClient struct {
    core.JDCloudClient
}

func NewEdcpsClient(credential *core.Credential) *EdcpsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("edcps.jdcloud-api.com")

    return &EdcpsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "edcps",
            Revision:    "1.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *EdcpsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *EdcpsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *EdcpsClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 修改分布式云物理服务器部分信息，包括名称、描述 */
func (c *EdcpsClient) ModifyInstance(request *edcps.ModifyInstanceRequest) (*edcps.ModifyInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ModifyInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 解绑弹性公网IP
 */
func (c *EdcpsClient) DisassociateElasticIp(request *edcps.DisassociateElasticIpRequest) (*edcps.DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建密钥对 */
func (c *EdcpsClient) CreateKeypairs(request *edcps.CreateKeypairsRequest) (*edcps.CreateKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.CreateKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询私有网络列表 */
func (c *EdcpsClient) DescribeVpcs(request *edcps.DescribeVpcsRequest) (*edcps.DescribeVpcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeVpcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询私有网络详情 */
func (c *EdcpsClient) DescribeVpc(request *edcps.DescribeVpcRequest) (*edcps.DescribeVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 批量查询分布式云物理服务器详细信息<br/>
支持分页查询，默认每页20条<br/>
 */
func (c *EdcpsClient) DescribeInstances(request *edcps.DescribeInstancesRequest) (*edcps.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改子网 */
func (c *EdcpsClient) ModifySubnet(request *edcps.ModifySubnetRequest) (*edcps.ModifySubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ModifySubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性公网IP列表<br/>
支持分页查询，默认每页20条<br/>
 */
func (c *EdcpsClient) DescribeElasticIps(request *edcps.DescribeElasticIpsRequest) (*edcps.DescribeElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询分布式云物理服务器支持的操作系统 */
func (c *EdcpsClient) DescribeOS(request *edcps.DescribeOSRequest) (*edcps.DescribeOSResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeOSResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个分布式云物理服务器硬件监控信息 */
func (c *EdcpsClient) DescribeInstanceStatus(request *edcps.DescribeInstanceStatusRequest) (*edcps.DescribeInstanceStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeInstanceStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询分布式云物理服务器实例类型 */
func (c *EdcpsClient) DescribeDeviceTypes(request *edcps.DescribeDeviceTypesRequest) (*edcps.DescribeDeviceTypesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeDeviceTypesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改弹性公网IP带宽
 */
func (c *EdcpsClient) ModifyElasticIpBandwidth(request *edcps.ModifyElasticIpBandwidthRequest) (*edcps.ModifyElasticIpBandwidthResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ModifyElasticIpBandwidthResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单个分布式云物理服务器已安装的RAID信息，包括系统盘RAID信息和数据盘RAID信息 */
func (c *EdcpsClient) DescribeInstanceRaid(request *edcps.DescribeInstanceRaidRequest) (*edcps.DescribeInstanceRaidResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeInstanceRaidResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重启单台分布式云物理服务器，只能重启running状态的服务器 [MFA enabled] */
func (c *EdcpsClient) RestartInstance(request *edcps.RestartInstanceRequest) (*edcps.RestartInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.RestartInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询子网列表 */
func (c *EdcpsClient) DescribeSubnets(request *edcps.DescribeSubnetsRequest) (*edcps.DescribeSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除私有网络
 */
func (c *EdcpsClient) DeleteVpc(request *edcps.DeleteVpcRequest) (*edcps.DeleteVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DeleteVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性公网IP库存 */
func (c *EdcpsClient) DescribeElasticIpStock(request *edcps.DescribeElasticIpStockRequest) (*edcps.DescribeElasticIpStockResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeElasticIpStockResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 导入密钥对 */
func (c *EdcpsClient) ImportKeypairs(request *edcps.ImportKeypairsRequest) (*edcps.ImportKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ImportKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询子网详情 */
func (c *EdcpsClient) DescribeSubnet(request *edcps.DescribeSubnetRequest) (*edcps.DescribeSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询链路类型列表 */
func (c *EdcpsClient) DescribeLineTypes(request *edcps.DescribeLineTypesRequest) (*edcps.DescribeLineTypesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeLineTypesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 对单台分布式云物理服务器执行开机操作，只能启动stopped状态的服务器 */
func (c *EdcpsClient) StartInstance(request *edcps.StartInstanceRequest) (*edcps.StartInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.StartInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询密钥对列表 */
func (c *EdcpsClient) DescribeKeypairs(request *edcps.DescribeKeypairsRequest) (*edcps.DescribeKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询密钥对详情 */
func (c *EdcpsClient) DescribeKeypair(request *edcps.DescribeKeypairRequest) (*edcps.DescribeKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 绑定弹性公网IP
 */
func (c *EdcpsClient) AssociateElasticIp(request *edcps.AssociateElasticIpRequest) (*edcps.AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询分布式云物理服务器库存 */
func (c *EdcpsClient) DescribeDeviceStock(request *edcps.DescribeDeviceStockRequest) (*edcps.DescribeDeviceStockResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeDeviceStockResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询单台分布式云物理服务器详细信息 */
func (c *EdcpsClient) DescribeInstance(request *edcps.DescribeInstanceRequest) (*edcps.DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一台或多台指定配置的分布式云物理服务器<br/>
- 地域与可用区<br/>
  - 调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域与可用区<br/>
- 实例类型<br/>
  - 调用接口（describeDeviceTypes）获取物理实例类型列表<br/>
  - 不能使用已下线、或已售馨的实例类型<br/>
- 操作系统<br/>
  - 可调用接口（describeOS）获取分布式云物理服务器支持的操作系统列表<br/>
- 存储<br/>
  - 数据盘多种RAID可选，可调用接口（describeDeviceRaids）获取服务器支持的RAID列表<br/>
- 网络<br/>
  - 网络类型目前支持vpc<br/>
  - 线路目前支持联通un、电信ct、移动cm<br/>
  - 支持不启用外网，如果启用外网，带宽范围[1,200] 单位Mbps<br/>
- 其他<br/>
  - 购买时长，可按年或月购买：月取值范围[1,9], 年取值范围[1,3]<br/>
  - 密码设置参考公共参数规范<br/>
 */
func (c *EdcpsClient) CreateInstances(request *edcps.CreateInstancesRequest) (*edcps.CreateInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.CreateInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可用的私有IP列表 */
func (c *EdcpsClient) DescribeAvailablePrivateIp(request *edcps.DescribeAvailablePrivateIpRequest) (*edcps.DescribeAvailablePrivateIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeAvailablePrivateIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询分布式云物理服务器名称 */
func (c *EdcpsClient) DescribeInstanceName(request *edcps.DescribeInstanceNameRequest) (*edcps.DescribeInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重置分布式云物理服务器密码
 */
func (c *EdcpsClient) ResetPassword(request *edcps.ResetPasswordRequest) (*edcps.ResetPasswordResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ResetPasswordResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建子网 */
func (c *EdcpsClient) CreateSubnet(request *edcps.CreateSubnetRequest) (*edcps.CreateSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.CreateSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性公网IP详情 */
func (c *EdcpsClient) DescribeElasticIp(request *edcps.DescribeElasticIpRequest) (*edcps.DescribeElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除密钥对 */
func (c *EdcpsClient) DeleteKeypairs(request *edcps.DeleteKeypairsRequest) (*edcps.DeleteKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DeleteKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改私有网络
 */
func (c *EdcpsClient) ModifyVpc(request *edcps.ModifyVpcRequest) (*edcps.ModifyVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ModifyVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询某种实例类型的分布式云物理服务器支持的RAID类型，可查询系统盘RAID类型和数据盘RAID类型 */
func (c *EdcpsClient) DescribeDeviceRaids(request *edcps.DescribeDeviceRaidsRequest) (*edcps.DescribeDeviceRaidsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeDeviceRaidsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建私有网络 */
func (c *EdcpsClient) CreateVpc(request *edcps.CreateVpcRequest) (*edcps.CreateVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.CreateVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重装分布式云物理服务器，只能重装stopped状态的服务器<br/>
- 可调用接口（describeOS）获取分布式云物理服务器支持的操作系统列表
 [MFA enabled] */
func (c *EdcpsClient) ReinstallInstance(request *edcps.ReinstallInstanceRequest) (*edcps.ReinstallInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ReinstallInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除子网 */
func (c *EdcpsClient) DeleteSubnet(request *edcps.DeleteSubnetRequest) (*edcps.DeleteSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DeleteSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 对单台分布式云物理服务器执行关机操作，只能停止running状态的服务器 [MFA enabled] */
func (c *EdcpsClient) StopInstance(request *edcps.StopInstanceRequest) (*edcps.StopInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.StopInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 申请弹性公网IP
 */
func (c *EdcpsClient) ApplyElasticIps(request *edcps.ApplyElasticIpsRequest) (*edcps.ApplyElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.ApplyElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询分布式分布式云物理服务器地域列表 */
func (c *EdcpsClient) DescribeEdCPSRegions(request *edcps.DescribeEdCPSRegionsRequest) (*edcps.DescribeEdCPSRegionsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &edcps.DescribeEdCPSRegionsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

