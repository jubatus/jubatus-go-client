// This file is auto-generated from clustering.idl(0.9.4-18-g4935b2b) with jenerator version 1.0.7-6-g1ae743a/master
// *** DO NOT EDIT ***

package jubatus_client

import (
	common "github.com/jubatus/jubatus-go-client/lib/common"
)

type WeightedDatum struct {
	Weight float64
	Point  common.Datum
}

type IndexedPoint struct {
	ID    string
	Point common.Datum
}

type WeightedIndex struct {
	Weight float64
	ID     string
}
