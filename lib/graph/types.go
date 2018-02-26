// This file is auto-generated from graph.idl(0.6.4-33-gcc8d7ca) with jenerator version 1.0.7-6-g1ae743a/master
// *** DO NOT EDIT ***

package jubatus_client

type Node struct {
	Property map[string]string
	InEdges  []uint64
	OutEdges []uint64
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
	MaxHop uint32
	Query  PresetQuery
}
