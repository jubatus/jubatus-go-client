// This file is auto-generated from graph.idl(0.6.4-33-gcc8d7ca) with jenerator version 0.9.4-42-g70f7539/master
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
