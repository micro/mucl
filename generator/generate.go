package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/micro/mu/generator/templates"
	"github.com/micro/mu/mucl"
)

type Generator struct {
	d            *mucl.Definition
	AbsolutePath string
	onlyTypes    bool
}

func NewGenerator(d *mucl.Definition, onlyTypes bool) (*Generator, error) {
	// Get the absolute path of the current directory
	absPath, err := filepath.Abs(".")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}
	return &Generator{
		d:            d,
		AbsolutePath: absPath,
		onlyTypes:    onlyTypes,
	}, nil
}

func (g *Generator) Generate() error {
	if !g.onlyTypes {
		err := g.GenerateGoMod()
		if err != nil {
			return err
		}
	}
	// if !g.onlyTypes {
	// 	err := g.GenerateInfra()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	err := g.GenerateTypes()
	if err != nil {
		return err
	}
	if !g.onlyTypes {
		err := g.GenerateServers()
		if err != nil {
			return err
		}
	}
	if !g.onlyTypes {
		err := g.GenerateHandlers()
		if err != nil {
			return err
		}
	}
	if !g.onlyTypes {
		err = g.Tidy()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) GenerateGoMod() error {
	cmd := exec.Command("go", "mod", "init", g.d.Project())
	cmd.Dir = g.AbsolutePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go mod init: %v", err)
	}

	return nil
}

func (g *Generator) GenerateInfra() error {
	infraPath := filepath.Join(g.AbsolutePath, "infra")
	// Check if the directory exists, if not create it
	if _, err := os.Stat(infraPath); os.IsNotExist(err) {
		err := os.MkdirAll(infraPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	infraFile, err := os.Create(fmt.Sprintf("%s/%s", infraPath, "plugins.go"))
	if err != nil {
		return err
	}
	defer infraFile.Close()
	svc := g.d.Service()

	infraTemplate := template.Must(template.New("type").Parse(string(templates.InfraTemplate())))
	err = infraTemplate.Execute(infraFile,
		map[string]interface{}{
			"Plugins": GetPluginList(svc),
			"Def":     g.d,
		})
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) Tidy() error {
	cmd := exec.Command("go", "get", "go-micro.dev/v5")
	cmd.Dir = g.AbsolutePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy: %v", err)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = g.AbsolutePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy: %v", err)
	}

	return nil
}

func (g *Generator) GenerateTypes() error {
	tt := g.d.Messages()
	for _, t := range tt {
		typePath := g.AbsolutePath
		// Check if the directory exists, if not create it
		if _, err := os.Stat(typePath); os.IsNotExist(err) {
			err := os.MkdirAll(typePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
		typeFile, err := os.Create(fmt.Sprintf("%s/%s", typePath, t.FileName()))
		if err != nil {
			return err
		}
		defer typeFile.Close()

		typeTemplate := template.Must(template.New("type").Parse(string(templates.TypeTemplate())))
		err = typeTemplate.Execute(typeFile, map[string]interface{}{
			"Module": g.d.Project(),
			"Def":    t,
		})
		if err != nil {
			return err
		}
	}
	ee := g.d.Enums()
	for _, e := range ee {
		typePath := g.AbsolutePath
		// Check if the directory exists, if not create it
		if _, err := os.Stat(typePath); os.IsNotExist(err) {
			err := os.MkdirAll(typePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
		typeFile, err := os.Create(fmt.Sprintf("%s/%s", typePath, e.FileName()))
		if err != nil {
			return err
		}
		defer typeFile.Close()

		typeTemplate := template.Must(template.New("type").Parse(string(templates.EnumTemplate())))
		err = typeTemplate.Execute(typeFile,
			map[string]interface{}{
				"Module": g.d.Project(),
				"Def":    e,
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) GenerateServers() error {
	tt := g.d.Servers()
	serverPath := filepath.Join(g.AbsolutePath, "cmd", g.d.Project())
	// Check if the directory exists, if not create it
	if _, err := os.Stat(serverPath); os.IsNotExist(err) {
		err := os.MkdirAll(serverPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	for _, t := range tt {

		serverFile, err := os.Create(fmt.Sprintf("%s/%s", serverPath, t.FileName()))
		if err != nil {
			return err
		}
		defer serverFile.Close()

		serverTemplate := template.Must(template.New("type").Parse(string(templates.ServiceTemplate())))
		err = serverTemplate.Execute(serverFile, map[string]interface{}{
			"ServiceName": g.d.ServiceName(),
			"Module":      g.d.Project(),
			"Def":         t,
		})
		if err != nil {
			return err
		}
		// pluginsFile, err := os.Create(fmt.Sprintf("%s/%s", serverPath, "plugins.go"))
		// if err != nil {
		// 	return err
		// }
		// defer pluginsFile.Close()
		// pluginsTemplate := template.Must(template.New("type").Parse(string(templates.PluginsTemplate())))
		// err = pluginsTemplate.Execute(pluginsFile, map[string]interface{}{
		// 	"Module": g.d.Project(),
		// })
		// if err != nil {
		// 	return err
		// }

		clientPath := g.AbsolutePath

		clientFile, err := os.Create(fmt.Sprintf("%s/%s", clientPath, t.ClientFileName()))
		if err != nil {
			return err
		}
		defer clientFile.Close()

		clientTemplate := template.Must(template.New("client").Parse(string(templates.ServiceClientTemplate())))
		err = clientTemplate.Execute(clientFile, map[string]interface{}{
			"Service": t,
			"Module":  g.d.Project(),
			"Project": g.d.Project(),
			"Def":     t,
		})
		if err != nil {
			return err
		}

	}
	return nil
}

func (g *Generator) GenerateHandlers() error {
	tt := g.d.Servers()
	for _, t := range tt {
		handlerPath := filepath.Join(g.AbsolutePath, "handlers")
		// Check if the directory exists, if not create it
		if _, err := os.Stat(handlerPath); os.IsNotExist(err) {
			err := os.MkdirAll(handlerPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
		}
		handlerFile, err := os.Create(fmt.Sprintf("%s/%s", handlerPath, strings.ToLower(t.Name)+".go"))
		if err != nil {
			return err
		}
		defer handlerFile.Close()

		handlerTemplate := template.Must(template.New("type").Parse(string(templates.HandlerTemplate())))
		err = handlerTemplate.Execute(handlerFile, map[string]interface{}{
			"Service": t,
			"Module":  g.d.Project(),
		})
		if err != nil {
			return err
		}

		// for _, m := range t.Methods() {
		// 	methodFile, err := os.Create(fmt.Sprintf("%s/%s", handlerPath, m.FileName()))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	defer methodFile.Close()

		// 	methodTemplate := template.Must(template.New("type").Parse(string(templates.ServiceHandlerTemplate())))
		// 	err = methodTemplate.Execute(methodFile, map[string]interface{}{
		// 		"Service": t,
		// 		"Def":     m,
		// 		"Module":  g.d.Project(),
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}
	return nil

}
