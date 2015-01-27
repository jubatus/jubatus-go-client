// This file is auto-generated from graph.idl(0.5.4-179-gb59b61b) with jenerator version 0.6.4-32-g63e3219/feature/go_client
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
