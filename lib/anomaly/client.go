// This file is auto-generated from anomaly.idl(0.5.2-68-g68e898d) with jenerator version 0.6.4-32-g63e3219/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "../common"
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

func (c *AnomalyClient) Add(row common.Datum) IdWithScore {
	var result IdWithScore
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
