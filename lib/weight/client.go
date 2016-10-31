// This file is auto-generated from weight.idl(0.9.0-24-gda61383) with jenerator version 0.9.4-42-g70f7539/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
)

type WeightClient struct {
	client rpc.Client
	name   string
}

func NewWeightClient(host string, name string) (*WeightClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &WeightClient{*client, name}, nil
}

func (c *WeightClient) Update(d common.Datum) []Feature {
	var result []Feature
	c.client.Call("update", codec.MsgpackSpecRpcMultiArgs{c.name, d}, &result)
	return result
}

func (c *WeightClient) CalcWeight(d common.Datum) []Feature {
	var result []Feature
	c.client.Call("calc_weight", codec.MsgpackSpecRpcMultiArgs{c.name, d},
		&result)
	return result
}

func (c *WeightClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *WeightClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *WeightClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *WeightClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *WeightClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *WeightClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *WeightClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *WeightClient) GetName() string {
	return c.name
}

func (c *WeightClient) SetName(new_name string) {
	c.name = new_name
}

func (c *WeightClient) GetClient() rpc.Client {
	return c.client
}
