// This file is auto-generated from graph.idl(0.6.4-33-gcc8d7ca) with jenerator version 0.8.5-6-g5a2c923/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

type GraphClient struct {
	client rpc.Client
	name   string
}

func NewGraphClient(host string, name string) (*GraphClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &GraphClient{*client, name}, nil
}

func (c *GraphClient) CreateNode() string {
	var result string
	c.client.Call("create_node", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *GraphClient) RemoveNode(node_id string) bool {
	var result bool
	c.client.Call("remove_node", codec.MsgpackSpecRpcMultiArgs{c.name, node_id},
		&result)
	return result
}

func (c *GraphClient) UpdateNode(node_id string,
	property map[string]string) bool {
	var result bool
	c.client.Call("update_node", codec.MsgpackSpecRpcMultiArgs{c.name, node_id,
		property}, &result)
	return result
}

func (c *GraphClient) CreateEdge(node_id string, e Edge) int64 {
	var result int64
	c.client.Call("create_edge", codec.MsgpackSpecRpcMultiArgs{c.name, node_id,
		e}, &result)
	return result
}

func (c *GraphClient) UpdateEdge(node_id string, edge_id int64, e Edge) bool {
	var result bool
	c.client.Call("update_edge", codec.MsgpackSpecRpcMultiArgs{c.name, node_id,
		edge_id, e}, &result)
	return result
}

func (c *GraphClient) RemoveEdge(node_id string, edge_id int64) bool {
	var result bool
	c.client.Call("remove_edge", codec.MsgpackSpecRpcMultiArgs{c.name, node_id,
		edge_id}, &result)
	return result
}

func (c *GraphClient) GetCentrality(node_id string, centrality_type int32,
	query PresetQuery) float64 {
	var result float64
	c.client.Call("get_centrality", codec.MsgpackSpecRpcMultiArgs{c.name,
		node_id, centrality_type, query}, &result)
	return result
}

func (c *GraphClient) AddCentralityQuery(query PresetQuery) bool {
	var result bool
	c.client.Call("add_centrality_query", codec.MsgpackSpecRpcMultiArgs{c.name,
		query}, &result)
	return result
}

func (c *GraphClient) AddShortestPathQuery(query PresetQuery) bool {
	var result bool
	c.client.Call("add_shortest_path_query",
		codec.MsgpackSpecRpcMultiArgs{c.name, query}, &result)
	return result
}

func (c *GraphClient) RemoveCentralityQuery(query PresetQuery) bool {
	var result bool
	c.client.Call("remove_centrality_query",
		codec.MsgpackSpecRpcMultiArgs{c.name, query}, &result)
	return result
}

func (c *GraphClient) RemoveShortestPathQuery(query PresetQuery) bool {
	var result bool
	c.client.Call("remove_shortest_path_query",
		codec.MsgpackSpecRpcMultiArgs{c.name, query}, &result)
	return result
}

func (c *GraphClient) GetShortestPath(query ShortestPathQuery) []string {
	var result []string
	c.client.Call("get_shortest_path", codec.MsgpackSpecRpcMultiArgs{c.name,
		query}, &result)
	return result
}

func (c *GraphClient) UpdateIndex() bool {
	var result bool
	c.client.Call("update_index", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *GraphClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *GraphClient) GetNode(node_id string) Node {
	var result Node
	c.client.Call("get_node", codec.MsgpackSpecRpcMultiArgs{c.name, node_id},
		&result)
	return result
}

func (c *GraphClient) GetEdge(node_id string, edge_id int64) Edge {
	var result Edge
	c.client.Call("get_edge", codec.MsgpackSpecRpcMultiArgs{c.name, node_id,
		edge_id}, &result)
	return result
}

func (c *GraphClient) CreateNodeHere(node_id string) bool {
	var result bool
	c.client.Call("create_node_here", codec.MsgpackSpecRpcMultiArgs{c.name,
		node_id}, &result)
	return result
}

func (c *GraphClient) RemoveGlobalNode(node_id string) bool {
	var result bool
	c.client.Call("remove_global_node", codec.MsgpackSpecRpcMultiArgs{c.name,
		node_id}, &result)
	return result
}

func (c *GraphClient) CreateEdgeHere(edge_id int64, e Edge) bool {
	var result bool
	c.client.Call("create_edge_here", codec.MsgpackSpecRpcMultiArgs{c.name,
		edge_id, e}, &result)
	return result
}

func (c *GraphClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *GraphClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *GraphClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *GraphClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *GraphClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *GraphClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *GraphClient) GetName() string {
	return c.name
}

func (c *GraphClient) SetName(new_name string) {
	c.name = new_name
}

func (c *GraphClient) GetClient() rpc.Client {
	return c.client
}
