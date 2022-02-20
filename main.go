package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/goserg/jat/pkg/colorify/colorjson"
	"github.com/goserg/jat/pkg/inputer"
	"github.com/goserg/jat/pkg/inputer/ifile"
	"github.com/goserg/jat/pkg/inputer/ipipe"
	"github.com/goserg/jat/pkg/inputer/istdin"
)

func main() {
	args := os.Args

	var input inputer.Inputer
	if len(args) >= 2 {
		input = ifile.New(os.Args[1:])
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			fatalError(err)
		}

		if info.Mode()&os.ModeCharDevice != 0 {
			input = ipipe.New()
		} else {
			input = istdin.New()
		}
	}

	dataChan := make(chan []byte)
	input.GetInput(dataChan)

	for data := range dataChan {
		result := bytes.Buffer{}
		err := json.Indent(&result, data, "", "    ")
		if err != nil {
			fatalError(err)
		}
		cj := colorjson.CJ{}
		colored, err := cj.Colorify(result.Bytes())
		if err != nil {
			fatalError(err)
		}

		os.Stdout.WriteString(string(colored) + "\n")
	}
	if err := input.GetError(); err != nil {
		fatalError(err)
	}
}

func fatalError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
