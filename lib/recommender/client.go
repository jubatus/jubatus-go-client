// This file is auto-generated from recommender.idl(0.5.2-68-g68e898d) with jenerator version 0.5.4-224-g49229fa/develop
// *** DO NOT EDIT ***

package jubatus_client

import (
    "net"
    "net/rpc"
    "github.com/ugorji/go/codec"
    common "../common"
)

type RecommenderClient struct {
    client rpc.Client
    name string
}

func NewRecommenderClient(host string, name string) (*RecommenderClient,
    error) {
    conn, err := net.Dial("tcp", host)
    if err != nil {
    return nil, err
    }
    mh := new(codec.MsgpackHandle)
    rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
    client := rpc.NewClientWithCodec(rpcCodec)
    return &RecommenderClient{*client, name}, nil
}

func (c *RecommenderClient) ClearRow(id string) bool {
    var result bool
    c.client.Call("clear_row", codec.MsgpackSpecRpcMultiArgs{c.name, id},
        &result)
    return result
}

func (c *RecommenderClient) UpdateRow(id string, row common.Datum) bool {
    var result bool
    c.client.Call("update_row", codec.MsgpackSpecRpcMultiArgs{c.name, id, row},
        &result)
    return result
}

func (c *RecommenderClient) Clear() bool {
    var result bool
    c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
    return result
}

func (c *RecommenderClient) CompleteRowFromId(id string) common.Datum {
    var result common.Datum
    c.client.Call("complete_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
        id}, &result)
    return result
}

func (c *RecommenderClient) CompleteRowFromDatum(
    row common.Datum) common.Datum {
    var result common.Datum
    c.client.Call("complete_row_from_datum",
        codec.MsgpackSpecRpcMultiArgs{c.name, row}, &result)
    return result
}

func (c *RecommenderClient) SimilarRowFromId(id string,
    size int32) []IdWithScore {
    var result []IdWithScore
    c.client.Call("similar_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
        id, size}, &result)
    return result
}

func (c *RecommenderClient) SimilarRowFromDatum(row common.Datum,
    size int32) []IdWithScore {
    var result []IdWithScore
    c.client.Call("similar_row_from_datum",
        codec.MsgpackSpecRpcMultiArgs{c.name, row, size}, &result)
    return result
}

func (c *RecommenderClient) DecodeRow(id string) common.Datum {
    var result common.Datum
    c.client.Call("decode_row", codec.MsgpackSpecRpcMultiArgs{c.name, id},
        &result)
    return result
}

func (c *RecommenderClient) GetAllRows() []string {
    var result []string
    c.client.Call("get_all_rows", codec.MsgpackSpecRpcMultiArgs{c.name},
        &result)
    return result
}

func (c *RecommenderClient) CalcSimilarity(lhs common.Datum,
    rhs common.Datum) float64 {
    var result float64
    c.client.Call("calc_similarity", codec.MsgpackSpecRpcMultiArgs{c.name, lhs,
        rhs}, &result)
    return result
}

func (c *RecommenderClient) CalcL2norm(row common.Datum) float64 {
    var result float64
    c.client.Call("calc_l2norm", codec.MsgpackSpecRpcMultiArgs{c.name, row},
        &result)
    return result
}
