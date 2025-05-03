// Package templates provides templates for generating code
package templates

func TypeTemplate() []byte {
	return []byte(`package {{.Module}}

{{range .Def}}type {{.Name}} struct { {{range .Fields}}
	{{.ExportedName}} {{ if .Repeated}}[]{{end}}{{.Type.String}}{{end}}
}
{{end}}
`)
}

func EnumTemplate() []byte {
	return []byte(`package {{.Module}}

type {{.Def.Name}} int

const ({{ $name := .Def.Name }}{{range .Def.Values}}
	{{$name}}{{.Value.ExportedName}} {{ $name }} = {{.Value.Value}}{{end}}
)
`)
}

func InfraTemplate() []byte {
	return []byte(`package infra

import (

	{{range .Plugins}}_ "{{.}}"
{{end}}

)
`)
}

func PluginsTemplate() []byte {
	return []byte(`package main

import (
	_ "{{.Module}}/infra"
)
`)
}

func HandlerTemplate() []byte {
	return []byte(`package handlers
	
import (
	"context"
	"{{.Module}}"
)

// {{.Service.Name}} is a struct for the {{.Service.Name}} endpoint
// It is the server implementation of the {{.Service.Name}}Server interface
// TODO: Add fields to the struct if needed for server dependencies and state
type {{.Service.Name}} struct {
}

{{ $server := .Service.Name }}{{$module := .Module}}{{range .Service.Methods}}// {{.Name}} is the implementation of the {{$server}}.{{.Name}} endpoint
func (s *{{$server}}) {{.Name}}(ctx context.Context, req *{{$module}}.{{.Request.String}}, rsp *{{$module}}.{{.Response.String}}) error {
	// TODO: implement the endpoint logic
	return nil
}{{end}}

// New{{.Service.Name}} creates a new {{.Service.Name}} struct
// TODO: Add parameters to the the function if needed to set server dependencies and state
func New{{.Service.Name}}() *{{.Service.Name}} {
	return &{{.Service.Name}}{}
}
	`)
}
func ServiceTemplate() []byte {
	return []byte(`package main

import (
	"{{.Module}}/handlers"
	"go-micro.dev/v5"
)

func main() {
	// create endpoint handler
	handler := handlers.New{{.Def.Name}}()
	
	// create go-micro service
	service := micro.New("{{.ServiceName}}")

	// register handler
	service.Handle(handler)

	// init service
	service.Init()

	// run service
	service.Run()
}
`)
}

func ServiceHandlerTemplate() []byte {
	return []byte(`package handlers

import (
	"context"

	"{{.Module}}"
)

{{ $server := .Service.Name }}func (s *{{$server}}) {{.Def.Name}}(ctx context.Context, req *{{.Module}}.{{.Def.Request.String}}, rsp *{{.Module}}.{{.Def.Response.String}}) error {

	return nil
}
`)
}

func ServiceClientTemplate() []byte {
	return []byte(`package {{.Module}}

import (
	"context"

	client "go-micro.dev/v5/client"
)

// Interface for {{.Def.Name}} service
type {{.Def.Name}}Server interface { {{ $server := .Service.Name }}
  {{range .Def.Methods}}{{.Name}}(ctx context.Context, req *{{.Request.String}},opts ...client.CallOption) (*{{.Response.String}},error)
{{end}} }

// {{.Def.ClientStructName}}Server implements the {{.Def.Name}}Server interface
// It is used to call the {{.Def.Name}} service
type {{.Def.ClientStructName}}Server struct {
	c    client.Client
	name string
}

// New{{.Def.Name}}Server creates a new {{.Def.Name}}Server
// It functions as a client for the {{.Def.Name}} service
func New{{.Def.Name}}Server(name string, c client.Client) {{.Def.Name}}Server{
	return &{{.Def.ClientStructName}}Server{
		c:    c,
		name: name,
	}
}

{{ $server := .Service.Name }}{{$def := .Def}}{{range .Def.Methods}}
func (c *{{$def.ClientStructName}}Server) {{.Name}}(ctx context.Context, in *{{.Request.String}}, opts ...client.CallOption) (*{{.Response.String}}, error) {
	req := c.c.NewRequest(c.name, "{{$server}}.{{.Name}}", in)
	out := new({{.Response.String}})
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
{{end}}
`)
}
func ConfigTemplate() []byte {
	return []byte(`project "{{.Module}}"
service {{.Service}} {}

type {{.Method}}Request {
  input string
}

type {{.Method}}Response {
  output string
}

server {{.Endpoint}} {
  rpc {{.Method}}({{.Method}}Request) returns ({{.Method}}Response)
}
`)
}
func TaskfileTemplate() []byte {
	return []byte(`# https://taskfile.dev

version: "3"

env:
  GO111MODULE: on
  GOPROXY: https://proxy.golang.org,direct

tasks:

  setup:
    desc: Install dependencies
    cmds:
      - go mod tidy

  build:
    desc: Build the binary
    sources:
      - ./**/*.go
    generates:
      - ./{{.BINARY_NAME}}
    cmds:
      - go build ./cmd/{{.SERVICE_NAME}}

  install:
    desc: Install the binary locally
    sources:
      - ./**/*.go
    cmds:
      - go install ./cmd/{{.SERVICE_NAME}} 

  test:
    desc: Run tests
    cmds:
      - go test -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./...  -timeout=15m

  cover:
    desc: Open the cover tool
    cmds:
      - go tool cover -html=coverage.txt

  ci:
    desc: Run all CI steps
    cmds:
      - task: build
      - task: test

  default:
    desc: Runs the default tasks
    cmds:
      - task: ci

  run:
    desc: Run the service
    deps:
      - build
    cmds:
      - go run ./cmd/{{.SERVICE_NAME}}

  clean:
    desc: Clean the project	
    cmds:
      - rm  ./{{.BINARY_NAME}}

`)
}
