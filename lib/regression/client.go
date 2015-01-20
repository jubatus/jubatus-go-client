// This file is auto-generated from regression.idl(0.5.2-68-g68e898d) with jenerator version 0.5.4-224-g49229fa/develop
// *** DO NOT EDIT ***

package jubatus_client

import (
    "net"
    "net/rpc"
    "github.com/ugorji/go/codec"
    common "../common"
)

type RegressionClient struct {
    client rpc.Client
    name string
}

func NewRegressionClient(host string, name string) (*RegressionClient, error) {
    conn, err := net.Dial("tcp", host)
    if err != nil {
    return nil, err
    }
    mh := new(codec.MsgpackHandle)
    rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
    client := rpc.NewClientWithCodec(rpcCodec)
    return &RegressionClient{*client, name}, nil
}

func (c *RegressionClient) Train(train_data []ScoredDatum) int32 {
    var result int32
    c.client.Call("train", codec.MsgpackSpecRpcMultiArgs{c.name, train_data},
        &result)
    return result
}

func (c *RegressionClient) Estimate(estimate_data []common.Datum) []float64 {
    var result []float64
    c.client.Call("estimate", codec.MsgpackSpecRpcMultiArgs{c.name,
        estimate_data}, &result)
    return result
}

func (c *RegressionClient) Clear() bool {
    var result bool
    c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
    return result
}
