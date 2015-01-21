# Jubatus Go Client

RPC client library for [Jubatus](https://github.com/jubatus/jubatus).

## Usage

First, you have to boot jubatus server.

```
$ jubaclassifier -f some-config.json
```

```go
package main

import (
	"fmt"
	jubatus "github.com/jubatus/jubatus-go-client/lib/classifier"
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
```

see more in [examples](https://github.com/jubatus/jubatus-go-client/tree/master/examples).

## Generate

You can regenerate library for latest jubatus.

You need `goimports` and `gofmt` to regenerate.
`gofmt` is automatically installed your environment if you have Golang SDK.
`goimports` have to be installed manually.

```
$ go get golang.org/x/tools/cmd/goimports
```

And you can regenerate with

```
$ ./generate.sh
```

