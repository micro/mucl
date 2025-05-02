package mucl

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func (m *Server) FileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + "_server.go"
}

func (m *Server) DirectoryName() string {
	if m == nil {
		return ""
	}
	return strings.ToLower(m.Name)
}

func (m *Server) ClientFileName() string {
	if m == nil {
		return ""
	}
	return strcase.ToSnake(m.Name) + "_client" + ".go"
}

func (m *Server) ClientStructName() string {
	if m == nil {
		return ""
	}
	return strcase.ToLowerCamel(m.Name)
}

func (m *Server) Methods() []*Method {
	if m == nil {
		return nil
	}
	var methods []*Method
	for _, entry := range m.Entry {
		if entry.Method != nil {
			methods = append(methods, entry.Method)
		}
	}
	return methods
}
