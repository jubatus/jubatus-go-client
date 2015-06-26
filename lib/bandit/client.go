// This file is auto-generated from bandit.idl(0.7.2-79-g2db27d7) with jenerator version 0.7.2-85-g1b6087f/fix-go-client
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

type BanditClient struct {
	client rpc.Client
	name   string
}

func NewBanditClient(host string, name string) (*BanditClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &BanditClient{*client, name}, nil
}

func (c *BanditClient) RegisterArm(arm_id string) bool {
	var result bool
	c.client.Call("register_arm", codec.MsgpackSpecRpcMultiArgs{c.name, arm_id},
		&result)
	return result
}

func (c *BanditClient) DeleteArm(arm_id string) bool {
	var result bool
	c.client.Call("delete_arm", codec.MsgpackSpecRpcMultiArgs{c.name, arm_id},
		&result)
	return result
}

func (c *BanditClient) SelectArm(player_id string) string {
	var result string
	c.client.Call("select_arm", codec.MsgpackSpecRpcMultiArgs{c.name,
		player_id}, &result)
	return result
}

func (c *BanditClient) RegisterReward(player_id string, arm_id string,
	reward float64) bool {
	var result bool
	c.client.Call("register_reward", codec.MsgpackSpecRpcMultiArgs{c.name,
		player_id, arm_id, reward}, &result)
	return result
}

func (c *BanditClient) GetArmInfo(player_id string) map[string]ArmInfo {
	var result map[string]ArmInfo
	c.client.Call("get_arm_info", codec.MsgpackSpecRpcMultiArgs{c.name,
		player_id}, &result)
	return result
}

func (c *BanditClient) Reset(player_id string) bool {
	var result bool
	c.client.Call("reset", codec.MsgpackSpecRpcMultiArgs{c.name, player_id},
		&result)
	return result
}

func (c *BanditClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BanditClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *BanditClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *BanditClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BanditClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BanditClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BanditClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *BanditClient) GetName() string {
	return c.name
}

func (c *BanditClient) SetName(new_name string) {
	c.name = new_name
}

func (c *BanditClient) GetClient() rpc.Client {
	return c.client
}
