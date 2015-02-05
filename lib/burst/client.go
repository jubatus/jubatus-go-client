// This file is auto-generated from burst.idl with jenerator version 0.6.4-35-gd438373/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
    common "../common"
    "github.com/ugorji/go/codec"
    "net"
    "net/rpc"
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
