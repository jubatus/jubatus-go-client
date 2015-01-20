package main
import (
	"fmt"
	//jubatus "github.com/jubatus/jubatus-go-client/lib/classifier"
	jubatus "../lib/classifier"
	jubacommon "../lib/common"
)

func main() {
	cli, err := jubatus.NewClassifierClient("localhost:9199", "hoge")
	if err != nil {
		fmt.Println(err)
		return
	}
	datum1 := jubacommon.NewDatum()
	datum1.AddNum("a", 1)
	cli.Train([]jubatus.LabeledDatum{jubatus.LabeledDatum{"fuga", datum1}})

	datum2 := jubacommon.NewDatum()
	datum2.AddNum("b", 1)
	cli.Train([]jubatus.LabeledDatum{jubatus.LabeledDatum{"hoge", datum2}})

	ret := cli.Classify([]jubacommon.Datum{datum1})
	fmt.Println(ret)
}
