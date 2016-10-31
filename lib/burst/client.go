// This file is auto-generated from burst.idl(0.6.4-96-g66ed74d) with jenerator version 0.9.4-42-g70f7539/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

type BurstClient struct {
	client rpc.Client
	name   string
}

func NewBurstClient(host string, name string) (*BurstClient, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	mh := new(codec.MsgpackHandle)
	mh.StructToArray = true
	rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
	client := rpc.NewClientWithCodec(rpcCodec)
	return &BurstClient{*client, name}, nil
}

func (c *BurstClient) AddDocuments(data []Document) int32 {
	var result int32
	c.client.Call("add_documents", codec.MsgpackSpecRpcMultiArgs{c.name, data},
		&result)
	return result
}

func (c *BurstClient) GetResult(keyword string) Window {
	var result Window
	c.client.Call("get_result", codec.MsgpackSpecRpcMultiArgs{c.name, keyword},
		&result)
	return result
}

func (c *BurstClient) GetResultAt(keyword string, pos float64) Window {
	var result Window
	c.client.Call("get_result_at", codec.MsgpackSpecRpcMultiArgs{c.name,
		keyword, pos}, &result)
	return result
}

func (c *BurstClient) GetAllBurstedResults() map[string]Window {
	var result map[string]Window
	c.client.Call("get_all_bursted_results",
		codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BurstClient) GetAllBurstedResultsAt(pos float64) map[string]Window {
	var result map[string]Window
	c.client.Call("get_all_bursted_results_at",
		codec.MsgpackSpecRpcMultiArgs{c.name, pos}, &result)
	return result
}

func (c *BurstClient) GetAllKeywords() []KeywordWithParams {
	var result []KeywordWithParams
	c.client.Call("get_all_keywords", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *BurstClient) AddKeyword(keyword KeywordWithParams) bool {
	var result bool
	c.client.Call("add_keyword", codec.MsgpackSpecRpcMultiArgs{c.name, keyword},
		&result)
	return result
}

func (c *BurstClient) RemoveKeyword(keyword string) bool {
	var result bool
	c.client.Call("remove_keyword", codec.MsgpackSpecRpcMultiArgs{c.name,
		keyword}, &result)
	return result
}

func (c *BurstClient) RemoveAllKeywords() bool {
	var result bool
	c.client.Call("remove_all_keywords", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *BurstClient) Clear() bool {
	var result bool
	c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BurstClient) Save(id string) map[string]string {
	var result map[string]string
	c.client.Call("save", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *BurstClient) Load(id string) bool {
	var result bool
	c.client.Call("load", codec.MsgpackSpecRpcMultiArgs{c.name, id}, &result)
	return result
}

func (c *BurstClient) GetConfig() string {
	var result string
	c.client.Call("get_config", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BurstClient) GetStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_status", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BurstClient) DoMix() bool {
	var result bool
	c.client.Call("do_mix", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
	return result
}

func (c *BurstClient) GetProxyStatus() map[string]map[string]string {
	var result map[string]map[string]string
	c.client.Call("get_proxy_status", codec.MsgpackSpecRpcMultiArgs{c.name},
		&result)
	return result
}

func (c *BurstClient) GetName() string {
	return c.name
}

func (c *BurstClient) SetName(new_name string) {
	c.name = new_name
}

func (c *BurstClient) GetClient() rpc.Client {
	return c.client
}
