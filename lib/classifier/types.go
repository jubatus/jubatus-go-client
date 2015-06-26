// This file is auto-generated from classifier.idl(0.7.2-49-g5a6436d) with jenerator version 0.7.2-85-g1b6087f/fix-go-client
// *** DO NOT EDIT ***

package jubatus_client

import (
	common "github.com/jubatus/jubatus-go-client/lib/common"
)

type EstimateResult struct {
	Label string
	Score float64
}

type LabeledDatum struct {
	Label string
	Data  common.Datum
}
