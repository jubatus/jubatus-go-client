// This file is auto-generated from clustering.idl(0.9.4-18-g4935b2b) with jenerator version 1.0.7-6-g1ae743a/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	common "github.com/jubatus/jubatus-go-client/lib/common"
	"github.com/ugorji/go/codec"
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

func (c *ClusteringClient) Push(points []IndexedPoint) bool {
	var result bool
	c.client.Call("push", codec.MsgpackSpecRpcMultiArgs{c.name, points},
		&result)
	return result
}

func (c *ClusteringClient) GetRevision() uint32 {
	var result uint32
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

func (c *ClusteringClient) GetCoreMembersLight() [][]WeightedIndex {
	var result [][]WeightedIndex
	c.client.Call("get_core_members_light",
		codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
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

func (c *ClusteringClient) GetNearestMembersLight(
	point common.Datum) []WeightedIndex {
	var result []WeightedIndex
	c.client.Call("get_nearest_members_light",
		codec.MsgpackSpecRpcMultiArgs{c.name, point}, &result)
	return result
}

func (c *ClusteringClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *ClusteringClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *ClusteringClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *ClusteringClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *ClusteringClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *ClusteringClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *ClusteringClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *ClusteringClient) GetName() string {
	return c.name
}

func (c *ClusteringClient) SetName(new_name string) {
	c.name = new_name
}

func (c *ClusteringClient) GetClient() rpc.Client {
	return c.client
}
