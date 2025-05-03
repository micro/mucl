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
type {{.Service.Name}} struct {
}
	
func New{{.Service.Name}}() *{{.Service.Name}} {
	return &{{.Service.Name}}{}
}

{{ $server := .Service.Name }}{{$module := .Module}}{{range .Service.Methods}}func (s *{{$server}}) {{.Name}}(ctx context.Context, req *{{$module}}.{{.Request.String}}, rsp *{{$module}}.{{.Response.String}}) error {

	return nil
}{{end}}
	
	`)
}
func ServiceTemplate() []byte {
	return []byte(`package main

import (
	"{{.Module}}/handlers"
	"go-micro.dev/v5"
)

func main() {

	handler := handlers.New{{.Def.Name}}()
	// create service

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
// Client API for {{.Def.Name}}

type {{.Def.Name}}Client interface {
	{{ $server := .Service.Name }}{{range .Def.Methods}}{{.Name}}(ctx context.Context, req *{{.Request.String}},opts ...client.CallOption) (*{{.Response.String}},error)
{{end}}
}

type {{.Def.ClientStructName}}Client struct {
	c    client.Client
	name string
}

func New{{.Def.Name}}Client(name string, c client.Client) {{.Def.Name}}Client {
	return &{{.Def.ClientStructName}}Client{
		c:    c,
		name: name,
	}
}

{{ $server := .Service.Name }}{{$def := .Def}}{{range .Def.Methods}}
func (c *{{$def.ClientStructName}}Client) {{.Name}}(ctx context.Context, in *{{.Request.String}}, opts ...client.CallOption) (*{{.Response.String}}, error) {
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
