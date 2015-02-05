// This file is auto-generated from nearest_neighbor.idl with jenerator version 0.6.4-35-gd438373/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
    common "../common"
    "github.com/ugorji/go/codec"
    "net"
    "net/rpc"
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
