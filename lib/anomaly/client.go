// This file is auto-generated from anomaly.idl(0.7.2-50-gbcc1e21) with jenerator version 0.8.5-6-g5a2c923/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
)

type AnomalyClient struct {
	client rpc.Client
	name   string
}

func NewAnomalyClient(host string, name string) (*AnomalyClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &AnomalyClient{*client, name}, nil
}

func (c *AnomalyClient) ClearRow(id string) bool {
	var result bool
	c.client.Call("clear_row", codec.MsgpackSpecRpcMultiArgs{c.name, id},
		&result)
	return result
}

func (c *AnomalyClient) Add(row common.Datum) IDWithScore {
	var result IDWithScore
	c.client.Call("add", codec.MsgpackSpecRpcMultiArgs{c.name, row}, &result)
	return result
}

func (c *AnomalyClient) Update(id string, row common.Datum) float64 {
	var result float64
	c.client.Call("update", codec.MsgpackSpecRpcMultiArgs{c.name, id, row},
		&result)
	return result
}

func (c *AnomalyClient) Overwrite(id string, row common.Datum) float64 {
	var result float64
	c.client.Call("overwrite", codec.MsgpackSpecRpcMultiArgs{c.name, id, row},
		&result)
	return result
}

func (c *AnomalyClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *AnomalyClient) CalcScore(row common.Datum) float64 {
	var result float64
	c.client.Call("calc_score", codec.MsgpackSpecRpcMultiArgs{c.name, row},
		&result)
	return result
}

func (c *AnomalyClient) GetAllRows() []string {
	var result []string
	c.client.Call("get_all_rows", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *AnomalyClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *AnomalyClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *AnomalyClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *AnomalyClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *AnomalyClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *AnomalyClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *AnomalyClient) GetName() string {
	return c.name
}

func (c *AnomalyClient) SetName(new_name string) {
	c.name = new_name
}

func (c *AnomalyClient) GetClient() rpc.Client {
	return c.client
}
