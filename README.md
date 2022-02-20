# JSON-CAT

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/goserg/jat/Go)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/goserg/jat/golangci-lint?label=linters)

Simple cli json prettifier with syntax highlighting

## Installation
    $ go install github.com/goserg/jat@latest

## Usage
with args
 
    $ jat examples/test1.json

or with pipe

    $ cat examples/test1.json | jat

or in interactive mode

    $ jat

![interactive mode](https://github.com/goserg/jat/blob/master/img/v0.2.0_interactive_mode.gif?raw=true)
