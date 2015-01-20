// This file is auto-generated from graph.idl(0.5.4-179-gb59b61b) with jenerator version 0.5.4-224-g49229fa/develop
// *** DO NOT EDIT ***

package jubatus_client

type Node struct {
    Property TMap.new(string, string)
    InEdges []int64
    OutEdges []int64
}

type Query struct {
    FromId string
    ToId string
}

type PresetQuery struct {
    EdgeQuery []Query
    NodeQuery []Query
}

type Edge struct {
    Property TMap.new(string, string)
    Source string
    Target string
}

type ShortestPathQuery struct {
    Source string
    Target string
    MaxHop int32
    Query PresetQuery
}

