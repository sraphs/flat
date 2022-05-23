# flat

[![CI](https://github.com/sraphs/flat/actions/workflows/ci.yml/badge.svg)](https://github.com/sraphs/flat/actions/workflows/ci.yml)

>  flatten/unflatten nested map or struct(only flatten support  struct).

## Usage

```go
package flat_test

import (
	"fmt"
	"time"

	"github.com/sraphs/flat"
)

func ExampleFlatten() {
	type ServerHttp struct {
		Network string
		Addr    string
		Timeout time.Duration
	}

	type ServerGrpc struct {
		Addr    string
		Timeout time.Duration
	}

	type Server struct {
		Http ServerHttp
	}

	type Bootstrap struct {
		Server Server
	}

	c := Bootstrap{
		Server: Server{
			Http: ServerHttp{
				Addr:    "0.0.0.0:8000",
				Timeout: 2 * time.Second,
			},
		},
	}

	opt := flat.Option{
		Separator: ".",
		Case:      flat.CaseLower,
	}

	flattened := flat.Flatten(c, opt)

	fmt.Println(flattened["server.http.addr"])
	fmt.Println(flattened["server.http.timeout"])

	// Output:
	// 0.0.0.0:8000
	// 2s
}

func ExampleUnflatten() {
	m := map[string]interface{}{
		"server.http.addr":    "0.0.0.0:8000",
		"server.http.timeout": "2s",
	}

	opt := flat.Option{
		Separator: ".",
		Case:      flat.CaseLower,
	}

	unflattened := flat.Unflatten(m, opt)

	fmt.Sprintln(unflattened)

	// Output:
}

```

## Contributing

We alway welcome your contributions :clap:

1.  Fork the repository
2.  Create Feat_xxx branch
3.  Commit your code
4.  Create Pull Request


## CHANGELOG
See [Releases](https://github.com/sraphs/flat/releases)

## License
[MIT © sraph.com](./LICENSE)
