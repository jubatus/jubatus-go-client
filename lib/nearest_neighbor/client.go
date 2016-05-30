// This file is auto-generated from nearest_neighbor.idl(0.8.2-20-g8e4dc3b) with jenerator version 0.8.5-6-g5a2c923/develop
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
)

type NearestNeighborClient struct {
	client rpc.Client
	name   string
}

func NewNearestNeighborClient(host string, name string) (*NearestNeighborClient,
	error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &NearestNeighborClient{*client, name}, nil
}

func (c *NearestNeighborClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *NearestNeighborClient) SetRow(id string, d common.Datum) bool {
	var result bool
	c.client.Call("set_row", codec.MsgpackSpecRpcMultiArgs{c.name, id, d},
		&result)
	return result
}

func (c *NearestNeighborClient) NeighborRowFromID(id string,
	size int32) []IDWithScore {
	var result []IDWithScore
	c.client.Call("neighbor_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
		id, size}, &result)
	return result
}

func (c *NearestNeighborClient) NeighborRowFromDatum(query common.Datum,
	size int32) []IDWithScore {
	var result []IDWithScore
	c.client.Call("neighbor_row_from_datum",
		codec.MsgpackSpecRpcMultiArgs{c.name, query, size}, &result)
	return result
}

func (c *NearestNeighborClient) SimilarRowFromID(id string,
	ret_num int32) []IDWithScore {
	var result []IDWithScore
	c.client.Call("similar_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
		id, ret_num}, &result)
	return result
}

func (c *NearestNeighborClient) SimilarRowFromDatum(query common.Datum,
	ret_num int32) []IDWithScore {
	var result []IDWithScore
	c.client.Call("similar_row_from_datum",
		codec.MsgpackSpecRpcMultiArgs{c.name, query, ret_num}, &result)
	return result
}

func (c *NearestNeighborClient) GetAllRows() []string {
	var result []string
	c.client.Call("get_all_rows", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *NearestNeighborClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *NearestNeighborClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *NearestNeighborClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *NearestNeighborClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *NearestNeighborClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *NearestNeighborClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *NearestNeighborClient) GetName() string {
	return c.name
}

func (c *NearestNeighborClient) SetName(new_name string) {
	c.name = new_name
}

func (c *NearestNeighborClient) GetClient() rpc.Client {
	return c.client
}
