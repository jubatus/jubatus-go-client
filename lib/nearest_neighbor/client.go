// This file is auto-generated from nearest_neighbor.idl(0.5.4-186-g163c6bd) with jenerator version 0.6.4-32-g63e3219/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "../common"
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

func (c *NearestNeighborClient) NeighborRowFromId(id string,
	size int32) []IdWithScore {
	var result []IdWithScore
	c.client.Call("neighbor_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
		id, size}, &result)
	return result
}

func (c *NearestNeighborClient) NeighborRowFromDatum(query common.Datum,
	size int32) []IdWithScore {
	var result []IdWithScore
	c.client.Call("neighbor_row_from_datum",
		codec.MsgpackSpecRpcMultiArgs{c.name, query, size}, &result)
	return result
}

func (c *NearestNeighborClient) SimilarRowFromId(id string,
	ret_num int32) []IdWithScore {
	var result []IdWithScore
	c.client.Call("similar_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
		id, ret_num}, &result)
	return result
}

func (c *NearestNeighborClient) SimilarRowFromDatum(query common.Datum,
	ret_num int32) []IdWithScore {
	var result []IdWithScore
	c.client.Call("similar_row_from_datum",
		codec.MsgpackSpecRpcMultiArgs{c.name, query, ret_num}, &result)
	return result
}
