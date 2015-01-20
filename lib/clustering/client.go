// This file is auto-generated from clustering.idl(0.6.1-17-g1bca359) with jenerator version 0.6.4-32-g63e3219/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
	common "../common"
	"github.com/ugorji/go/codec"
	"net"
	"net/rpc"
)

type ClusteringClient struct {
	client rpc.Client
	name   string
}

func NewClusteringClient(host string, name string) (*ClusteringClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &ClusteringClient{*client, name}, nil
}

func (c *ClusteringClient) Push(points []common.Datum) bool {
	var result bool
	c.client.Call("push", codec.MsgpackSpecRpcMultiArgs{c.name, points},
		&result)
	return result
}

func (c *ClusteringClient) GetRevision() int32 {
	var result int32
	c.client.Call("get_revision", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *ClusteringClient) GetCoreMembers() [][]WeightedDatum {
	var result [][]WeightedDatum
	c.client.Call("get_core_members", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *ClusteringClient) GetKCenter() []common.Datum {
	var result []common.Datum
	c.client.Call("get_k_center", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *ClusteringClient) GetNearestCenter(point common.Datum) common.Datum {
	var result common.Datum
	c.client.Call("get_nearest_center", codec.MsgpackSpecRpcMultiArgs{c.name,
		point}, &result)
	return result
}

func (c *ClusteringClient) GetNearestMembers(
	point common.Datum) []WeightedDatum {
	var result []WeightedDatum
	c.client.Call("get_nearest_members", codec.MsgpackSpecRpcMultiArgs{c.name,
		point}, &result)
	return result
}

func (c *ClusteringClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}
