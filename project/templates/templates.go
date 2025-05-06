// Package templates provides templates for generating code
package templates

func TypeTemplate() []byte {
	return []byte(`package {{.Module}}

{{range $m,$v := .Service.MessageMap }}// {{ $m }} is a struct for the {{ $m }} type
type {{ $m }} struct { {{range $f, $fv := $v.FieldMap}}
  {{$fv.DeclarationName}} {{$fv.DeclarationType}}{{end}}
}
{{end}}
{{range $m,$v := .Service.EnumMap}}// {{ $m }} is a type for the {{ $m }} enum
type {{$v.Name}} int

const ({{range $v.Values}}
	{{.Key}} {{$v.Name}} = {{.Value }}{{end}}
)
{{end}}
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
	return []byte(`// Package handlers contains the implementation of the {{.Service.Name}} service
package handlers
	
import (
	"context"
	"{{.Module}}"
)

// {{.Endpoint.Name}} is a struct for the {{.Endpoint.Name}} endpoint
// It is the server implementation of the {{.Endpoint.Name}}Server interface
// TODO: Add fields to the struct if needed for server dependencies and state
type {{.Endpoint.Name}} struct {
}

{{ $server := .Endpoint.Name }}{{$module := .Module}}{{range .Endpoint.GetAllMethods}}// {{.Name}} is the implementation of the {{$server}}.{{.Name}} endpoint
func (s *{{$server}}) {{.Name}}(ctx context.Context, req *{{$module}}.{{.RequestTypeName}}, rsp *{{$module}}.{{.ResponseTypeName}}) error {
	// TODO: implement the endpoint logic
	return nil
}{{end}}

// New{{.Endpoint.Name}} creates a new {{.Endpoint.Name}} struct
// TODO: Add parameters to the the function if needed to set server dependencies and state
func New{{.Endpoint.Name}}() *{{.Endpoint.Name}} {
	return &{{.Endpoint.Name}}{}
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

	
	// create go-micro service
	service := micro.New("{{.Service.Name}}")

  {{range .Service.GetAllEndpoints}}
  // {{.Name}} handler
  {{.Name}}Handler := handlers.New{{.Name}}()
	// register {{.Name}}Handler
	service.Handle({{.Name}}Handler)
{{end}}

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

{{ $server := .Endpoint.Name }}func (s *{{$server}}) {{.Def.Name}}(ctx context.Context, req *{{.Module}}.{{.Def.Request.String}}, rsp *{{.Module}}.{{.Def.Response.String}}) error {

	return nil
}
`)
}

func ServiceClientTemplate() []byte {
	return []byte(`// Package {{.Module}} defines the types and interfaces for the {{.Def.Name}} service
package {{.Module}}

import (
	"context"

	client "go-micro.dev/v5/client"
)

{{$service := .Service}}{{range .Service.GetAllEndpoints}}
// Endpoint {{.Name}}

// Interface for {{$service.Name}}/{{.Name}} endpoint
type {{.Name}}Server interface { 
  {{range .GetAllMethods}}{{.Name}}(ctx context.Context, req *{{.RequestTypeName}},opts ...client.CallOption) (*{{.ResponseTypeName}},error)
{{end}} }

// {{.ClientStructName}}Server implements the {{.Name}}Server interface
// It is used to call the {{.Name}} service
type {{.ClientStructName}}Server struct {
	c    client.Client
	name string
}

// New{{.Name}}Server creates a new {{.Name}}Server
// It functions as a client for the {{.Name}} service
func New{{.Name}}Server(name string, c client.Client) {{.Name}}Server{
	return &{{.ClientStructName}}Server{
		c:    c,
		name: name,
	}
}

{{$def := .}}{{range .GetAllMethods}}
func (c *{{$def.ClientStructName}}Server) {{.Name}}(ctx context.Context, in *{{.RequestTypeName}}, opts ...client.CallOption) (*{{.ResponseTypeName}}, error) {
	req := c.c.NewRequest(c.name, "{{$def}}.{{.Name}}", in)
	out := new({{.ResponseTypeName}})
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
{{end}}
{{end}}
`)
}

func ConfigTemplate() []byte {
	return []byte(`service {{.Service}} {}

type {{.Method}}Request {
  input string
}

type {{.Method}}Response {
  output string
}

endpoint {{.Endpoint}} {
  rpc {{.Method}}({{.Method}}Request) returns ({{.Method}}Response)
}
`)
}

func GitIgnoreTemplate() []byte {
	return []byte(`.DS_Store  
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env

# The binary output of the build tool
/{{.SERVICE_NAME}}

# Taskfile
/.task

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
      - ./{{.SERVICE_NAME}}
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
      - ./{{.SERVICE_NAME}}

  clean:
    desc: Clean the project	
    cmds:
      - rm ./{{.SERVICE_NAME}}

`)
}
