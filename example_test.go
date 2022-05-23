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

	flattened := opt.Flatten(c)

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

	unflattened := opt.Unflatten(m)

	fmt.Sprintln(unflattened)

	// Output:
}
