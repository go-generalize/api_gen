package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	_ "github.com/go-generalize/api_gen/server_generator/statik"
	"github.com/iancoleman/strcase"
	"github.com/rakyll/statik/fs"
)

var supportHTTPMethod = []string{
	"get",
	"post",
	"put",
	"delete",
	"patch",
}

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Specify the target directory.")
		os.Exit(1)
	}
	err := run(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func run(arg string) error {
	rootPath, err := filepath.Abs(arg)
	if err != nil {
		return err
	}

	packageRootPath := GetGoRootPath()
	basePackagePath, err := GetGoRootPackageName()
	if err != nil {
		return err
	}
	bootstrapTemplates := make([]*BootstrapTemplates, 0)
	packageName := ""

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(packageRootPath, path)
		if err != nil {
			return err
		}
		packagePath := filepath.Join(basePackagePath+"/", relPath)
		if rootPath == path {
			packagePath = ""
		}

		cs, err := parsePackages(path)
		if err != nil {
			return err
		}

		endpointPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}
		if len(cs) == 0 {
			return nil
		}

		endpointPath = filepath.Join(endpointPath, "/")
		endpoint := endpointPath
		if endpointPath == "." {
			packageName = cs[0].Package
			endpoint = ""
			endpointPath = "/"
		} else {
			endpoint += "/"
			endpointPath = "/" + endpointPath
		}

		bootstrapTemplates = append(bootstrapTemplates, &BootstrapTemplates{
			PackagePath:  packagePath,
			EndpointPath: endpointPath,
			Endpoint:     endpoint,
			Controller:   cs[0],
		})

		return nil
	})
	if err != nil {
		return err
	}

	for _, b := range bootstrapTemplates {
		ep := b.EndpointPath
		if ep == "/" {
			b.ParentIndex = -1
			continue
		}

		for i, b2 := range bootstrapTemplates {
			rel := ""
			rel, err = filepath.Rel(b2.EndpointPath, ep)
			if err != nil {
				continue
			}
			if rel == "." || strings.HasPrefix(rel, "..") {
				continue
			}

			b.ParentIndex = i
			b.Endpoint = rel + "/"

			if !strings.HasSuffix(b.EndpointPath, "/") {
				b.EndpointPath = b.EndpointPath + "/"
			}
		}
	}

	bootstrapFilePath := filepath.Join(rootPath+"/", "bootstrap.go")
	err = createFromTemplate(
		"/bootstrap_template.go.tmpl",
		bootstrapFilePath, &BootstrapTemplate{
			PackageName: packageName,
			Bootstraps:  bootstrapTemplates,
		},
		true)
	if err != nil {
		return err
	}
	return nil
}

func parsePackages(path string) ([]*ControllerTemplate, error) {
	routes := make(map[string][]*ControllerTemplate)
	structPair, err := findStructPairList(path)
	if err != nil {
		return nil, err
	}
	for cn, s := range structPair {
		req := s.Request
		res := s.Response

		if req == nil || res == nil {
			continue
		}

		structName := req.StructName

		lowerStructName := strings.ToLower(structName)

		httpMethod := ""
		for _, m := range supportHTTPMethod {
			if !strings.HasPrefix(lowerStructName, m) {
				continue
			}
			httpMethod = m
			break
		}

		createDir := filepath.Clean(path + "/")
		createFileName := strcase.ToSnake(cn) + "_controller.go"
		createPath := filepath.Join(createDir, createFileName)

		if _, ok := routes[createDir]; !ok {
			routes[createDir] = make([]*ControllerTemplate, 0)
		}

		endpoint := string([]rune(cn)[len(httpMethod):])
		if ep, ok := epMap[structName]; ok && ep != "" {
			endpoint = ep
		} else {
			endpoint = strcase.ToSnake(endpoint)
		}

		ct := &ControllerTemplate{
			Package:               req.PackageName,
			ControllerName:        fmt.Sprintf("%sController", cn),
			ControllerNameInitial: strings.ToLower(string([]rune(cn)[0])),
			HandlerName:           cn,
			Endpoint:              endpoint,
			HTTPMethod:            strings.ToUpper(httpMethod),
			RequestStructName:     req.StructName,
			ResponseStructName:    res.StructName,
		}

		routes[createDir] = append(routes[createDir], ct)

		err = createFromTemplate("/controller_template.go.tmpl", createPath, ct, false)
		if err != nil {
			return nil, err
		}
	}

	controllers := make([]*ControllerTemplate, 0)
	for dir, cs := range routes {
		routePath := filepath.Join(dir+"/", "routes.go")
		packageName := ""
		if len(cs) > 0 {
			packageName = cs[0].Package
		}

		sort.Slice(cs, func(i, j int) bool {
			// nolint:scopelint
			return cs[i].Endpoint < cs[j].Endpoint
		})

		sort.SliceStable(cs, func(i, j int) bool {
			// nolint:scopelint
			return cs[i].HTTPMethod < cs[j].HTTPMethod
		})

		controllers = append(controllers, cs...)

		err := createFromTemplate("/routes_template.go.tmpl", routePath, &RoutesTemplate{
			Package:     packageName,
			Controllers: cs,
		}, true)
		if err != nil {
			return nil, err
		}
	}

	return controllers, nil
}

func createFromTemplate(templatePath, path string, m interface{}, isOverRide bool) error {
	_, err := os.Stat(path)
	if err == nil {
		if !isOverRide {
			return nil
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	statikFs, err := fs.New()
	if err != nil {
		return err
	}
	f, err := statikFs.Open(templatePath)
	if err != nil {
		return err
	}

	t, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	tpl := template.Must(template.New("tmpl").Parse(string(t)))

	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer w.Close()

	err = tpl.Execute(w, m)
	if err != nil {
		return err
	}

	_, err = ExecCommand("goimports", "-w", path)
	if err != nil {
		return err
	}

	return nil
}
