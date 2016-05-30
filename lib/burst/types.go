// This file is auto-generated from burst.idl(0.6.4-96-g66ed74d) with jenerator version 0.8.5-6-g5a2c923/develop
// *** DO NOT EDIT ***

package jubatus_client

type KeywordWithParams struct {
	Keyword      string
	ScalingParam float64
	Gamma        float64
}

type Batch struct {
	AllDataCount      int32
	RelevantDataCount int32
	BurstWeight       float64
}

type Window struct {
	StartPos float64
	Batches  []Batch
}

type Document struct {
	Pos  float64
	Text string
}
