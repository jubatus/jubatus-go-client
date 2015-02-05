// This file is auto-generated from classifier.idl with jenerator version 0.6.4-35-gd438373/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

import (
    common "../common"
    "github.com/ugorji/go/codec"
    "net"
    "net/rpc"
)

type ClassifierClient struct {
    client rpc.Client
    name   string
}

func NewClassifierClient(host string, name string) (*ClassifierClient, error) {
    conn, err := net.Dial("tcp", host)
    if err != nil {
        return nil, err
    }
    mh := new(codec.MsgpackHandle)
    mh.StructToArray = true
    rpcCodec := codec.MsgpackSpecRpc.ClientCodec(conn, mh)
    client := rpc.NewClientWithCodec(rpcCodec)
    return &ClassifierClient{*client, name}, nil
}

func (c *ClassifierClient) Train(data []LabeledDatum) int32 {
    var result int32
    c.client.Call("train", codec.MsgpackSpecRpcMultiArgs{c.name, data}, &result)
    return result
}

func (c *ClassifierClient) Classify(data []common.Datum) [][]EstimateResult {
    var result [][]EstimateResult
    c.client.Call("classify", codec.MsgpackSpecRpcMultiArgs{c.name, data},
        &result)
    return result
}

func (c *ClassifierClient) GetLabels() []string {
    var result []string
    c.client.Call("get_labels", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
    return result
}

func (c *ClassifierClient) SetLabel(new_label string) bool {
    var result bool
    c.client.Call("set_label", codec.MsgpackSpecRpcMultiArgs{c.name, new_label},
        &result)
    return result
}

func (c *ClassifierClient) Clear() bool {
    var result bool
    c.client.Call("clear", codec.MsgpackSpecRpcMultiArgs{c.name}, &result)
    return result
}

func (c *ClassifierClient) DeleteLabel(target_label string) bool {
    var result bool
    c.client.Call("delete_label", codec.MsgpackSpecRpcMultiArgs{c.name,
        target_label}, &result)
    return result
}
