// This file is auto-generated from recommender.idl(0.5.2-68-g68e898d) with jenerator version 0.6.4-39-g6dfab43/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
)

type RecommenderClient struct {
	client rpc.Client
	name   string
}

func NewRecommenderClient(host string, name string) (*RecommenderClient,
	error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
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

func (c *RecommenderClient) CompleteRowFromID(id string) common.Datum {
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

func (c *RecommenderClient) SimilarRowFromID(id string,
	size int32) []IDWithScore {
	var result []IDWithScore
	c.client.Call("similar_row_from_id", codec.MsgpackSpecRpcMultiArgs{c.name,
		id, size}, &result)
	return result
}

func (c *RecommenderClient) SimilarRowFromDatum(row common.Datum,
	size int32) []IDWithScore {
	var result []IDWithScore
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

func (c *RecommenderClient) Save(id string) bool {
	var result bool
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *RecommenderClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *RecommenderClient) GetConfig() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RecommenderClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *RecommenderClient) GetName() string {
	return c.name
}

func (c *RecommenderClient) SetName(new_name string) {
	c.name = new_name
}

func (c *RecommenderClient) GetClient() rpc.Client {
	return c.client
}