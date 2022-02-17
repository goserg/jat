package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("need file name")
		os.Exit(1)
	}
	for i, fileName := range os.Args {
		if i == 0 {
			continue
		}
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		result := bytes.Buffer{}
		err = json.Indent(&result, b, "", "    ")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Stdout.WriteString(result.String() + "\n")
	}
}