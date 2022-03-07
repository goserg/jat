package main

import (
	"bytes"
	"encoding/json"
	"log"
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
			log.Fatal(err.Error())
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
			log.Fatal(err.Error())
		}
		cj := colorjson.CJ{}
		colored, err := cj.Colorify(result.Bytes())
		if err != nil {
			log.Fatal(err.Error())
		}

		os.Stdout.WriteString(string(colored) + "\n")
	}
	if err := input.GetError(); err != nil {
		log.Fatal(err.Error())
	}
}
