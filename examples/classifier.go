package main
import (
	"fmt"
	jubatus "github.com/jubatus/jubatus-go-client/lib/jubatus/classifier"
)

func main() {
	fmt.Println("hello")
	cli := jubatus.NewClassifierClient("localhost:9199", "hoge")
	datum := jubatus.NewDatum()
	datum.addNumValue("a", 1)
	labeled_datum := jubatus.LabeledDatum("a", datum)
	arg := []jubatus.LabeledDatum
	arg[0] = labeled_datum
	cli.Train(arg)
}
