// This file is auto-generated from regression.idl(0.6.4-33-gcc8d7ca) with jenerator version 0.9.4-42-g70f7539/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
)

type RegressionClient struct {
	client rpc.Client
	name   string
}

func NewRegressionClient(host string, name string) (*RegressionClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &RegressionClient{*client, name}, nil
}

func (c *RegressionClient) Train(train_data []ScoredDatum) int32 {
	var result int32
	c.client.Call("train", codec.MsgpackSpecRpcMultiArgs{c.name, train_data},
		&result)
	return result
}

func (c *RegressionClient) Estimate(estimate_data []common.Datum) []float64 {
	var result []float64
	c.client.Call("estimate", codec.MsgpackSpecRpcMultiArgs{c.name,
		estimate_data}, &result)
	return result
}

func (c *RegressionClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RegressionClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *RegressionClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *RegressionClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RegressionClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RegressionClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RegressionClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *RegressionClient) GetName() string {
	return c.name
}

func (c *RegressionClient) SetName(new_name string) {
	c.name = new_name
}

func (c *RegressionClient) GetClient() rpc.Client {
	return c.client
}
