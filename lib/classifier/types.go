// This file is auto-generated from classifier.idl with jenerator version 0.6.4-35-gd438373/feature/go_client
// *** DO NOT EDIT ***

package jubatus_client
import (
    common "../common"
)

type EstimateResult struct {
    Label string
    Score float64
}

type LabeledDatum struct {
    Label string
    Data common.Datum
}

