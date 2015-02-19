package main

import (
	"fmt"
	//jubatus "github.com/jubatus/jubatus-go-client/lib/classifier"
	classifier "github.com/jubatus/jubatus-go-client/lib/classifier"
	common "github.com/jubatus/jubatus-go-client/lib/common"
)

func main() {
	cli, err := classifier.NewClassifierClient("localhost:9199", "hoge")
	if err != nil {
		fmt.Println(err)
		return
	}
	datum1 := common.NewDatum()
	datum1.AddNum("a", 1)
	cli.Train([]classifier.LabeledDatum{classifier.LabeledDatum{"fuga", datum1}})

	datum2 := common.NewDatum()
	datum2.AddNum("b", 1)
	cli.Train([]classifier.LabeledDatum{classifier.LabeledDatum{"hoge", datum2}})

	ret := cli.Classify([]common.Datum{datum1})
	fmt.Println(ret)
}
