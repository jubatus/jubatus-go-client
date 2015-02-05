// This file is auto-generated from graph.idl with jenerator version 0.6.4-35-gd438373/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client
import (
    common "../common"
)

type Node struct {
    Property map[string]string
    InEdges []int64
    OutEdges []int64
}

type Query struct {
    FromID string
    ToID string
}

type PresetQuery struct {
    EdgeQuery []Query
    NodeQuery []Query
}

type Edge struct {
    Property map[string]string
    Source string
    Target string
}

type ShortestPathQuery struct {
    Source string
    Target string
    MaxHop int32
    Query PresetQuery
}

