// This file is auto-generated from classifier.idl(0.6.1-18-gbb16715) with jenerator version 0.6.4-39-g6dfab43/feature/go_client
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
