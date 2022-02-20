# JSON-CAT

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/goserg/jat/Go)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/goserg/jat/golangci-lint?label=linters)

Simple cli json viewer with syntax highlighting

## Installation
    $ go install github.com/goserg/jat@latest

## Usage
with args
 
    $ jat examples/test1.json

or with pipe

    $ cat examples/test1.json | jat