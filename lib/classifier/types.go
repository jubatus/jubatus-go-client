// This file is auto-generated from classifier.idl(0.8.9-17-gd4c007f) with jenerator version 0.8.5-6-g5a2c923/develop
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
