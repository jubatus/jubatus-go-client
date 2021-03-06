# Jubatus Go Client

RPC client library for [Jubatus](https://github.com/jubatus/jubatus).
All code generated by `jenerator` in Jubatus.

# Install

```
$ go get github.com/jubatus/jubatus-go-client/lib
```

## Usage

First, you have to boot jubatus server.

```
$ jubaclassifier -f some-config.json
```

```go
package main

import (
	"fmt"
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

# License

MIT License
