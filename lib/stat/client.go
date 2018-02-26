// This file is auto-generated from stat.idl(0.6.4-33-gcc8d7ca) with jenerator version 1.0.7-6-g1ae743a/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

type StatClient struct {
	client rpc.Client
	name   string
}

func NewStatClient(host string, name string) (*StatClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &StatClient{*client, name}, nil
}

func (c *StatClient) Push(key string, value float64) bool {
	var result bool
	c.client.Call("push", codec.MsgpackSpecRpcMultiArgs{c.name, key, value},
		&result)
	return result
}

func (c *StatClient) Sum(key string) float64 {
	var result float64
	c.client.Call("sum", codec.MsgpackSpecRpcMultiArgs{c.name, key}, &result)
	return result
}

func (c *StatClient) Stddev(key string) float64 {
	var result float64
	c.client.Call("stddev", codec.MsgpackSpecRpcMultiArgs{c.name, key}, &result)
	return result
}

func (c *StatClient) Max(key string) float64 {
	var result float64
	c.client.Call("max", codec.MsgpackSpecRpcMultiArgs{c.name, key}, &result)
	return result
}

func (c *StatClient) Min(key string) float64 {
	var result float64
	c.client.Call("min", codec.MsgpackSpecRpcMultiArgs{c.name, key}, &result)
	return result
}

func (c *StatClient) Entropy(key string) float64 {
	var result float64
	c.client.Call("entropy", codec.MsgpackSpecRpcMultiArgs{c.name, key},
		&result)
	return result
}

func (c *StatClient) Moment(key string, degree int32, center float64) float64 {
	var result float64
	c.client.Call("moment", codec.MsgpackSpecRpcMultiArgs{c.name, key, degree,
		center}, &result)
	return result
}

func (c *StatClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *StatClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *StatClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *StatClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *StatClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *StatClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *StatClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *StatClient) GetName() string {
	return c.name
}

func (c *StatClient) SetName(new_name string) {
	c.name = new_name
}

func (c *StatClient) GetClient() rpc.Client {
	return c.client
}
