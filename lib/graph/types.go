// This file is auto-generated from graph.idl(0.5.4-179-gb59b61b) with jenerator version 0.6.4-39-g6dfab43/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client

type Node struct {
	Property map[string]string
	InEdges  []int64
	OutEdges []int64
}

type Query struct {
	FromID string
	ToID   string
}

type PresetQuery struct {
	EdgeQuery []Query
	NodeQuery []Query
}

type Edge struct {
	Property map[string]string
	Source   string
	Target   string
}

type ShortestPathQuery struct {
	Source string
	Target string
	MaxHop int32
	Query  PresetQuery
}
