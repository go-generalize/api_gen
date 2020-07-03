package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/go-generalize/api_gen/common"

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

const rootPackageName = "root"

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Specify the target directory.")
		os.Exit(1)
	}

	versionFlag := flag.Bool("v", false, "print version")

	flag.Parse()

	if *versionFlag {
		fmt.Printf(common.AppVersion)
		return
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
		fmt.Println("Make sure go.mod exists")

		return err
	}
	bootstrapTemplates := make([]*BootstrapTemplates, 0)
	packageName := ""
	endpointReplaceMatchRule := regexp.MustCompile(`:(.*?)/`)

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

		importPackageName := ""
		for i, p := range strings.Split(relPath, "/") {
			if i < 2 {
				continue
			}

			if importPackageName == "" {
				importPackageName += p
			} else {
				importPackageName += strings.ToUpper(p[:1]) + p[1:]
			}
		}

		endpointPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return err
		}

		endpointPath = filepath.Join(endpointPath, "/")
		endpoint := endpointPath
		if endpointPath == "." {
			endpoint = ""
			endpointPath = "/"
		} else {
			endpoint += "/"
			endpointPath = "/" + endpointPath
		}

		endpointParams := make([]string, 0)
		endpointPath = strings.ReplaceAll(endpointPath, "/_", "/:")
		endpointParam := strings.ReplaceAll(endpoint, "/_", "/:")
		endPointParamsFromRegexp := endpointReplaceMatchRule.FindAllStringSubmatch(endpointPath+"/", -1)

		for _, e := range endPointParamsFromRegexp {
			endpointParams = append(endpointParams, e[1])
		}

		cs, err := parsePackages(path, endpointParam, endpointParams)
		if err != nil {
			return err
		}
		if len(cs) == 0 {
			return nil
		}
		if endpointPath == "/" {
			packageName = cs[0].Package
		}

		bootstrapTemplates = append(bootstrapTemplates, &BootstrapTemplates{
			PackagePath:       packagePath,
			ImportPackageName: strcase.ToLowerCamel(importPackageName),
			EndpointPath:      endpointPath,
			Endpoint:          endpoint,
			Controller:        cs[0],
		})

		return nil
	})
	if err != nil {
		return err
	}

	for _, b := range bootstrapTemplates {
		ep := b.EndpointPath
		if ep == "/" {
			b.ParentPackageName = ""
			b.HasParent = false
			b.RouteGroupName = rootPackageName
			continue
		}

		for _, b2 := range bootstrapTemplates {
			rel := ""
			rel, err = filepath.Rel(b2.EndpointPath, ep)
			if err != nil {
				continue
			}
			if rel == "." || strings.HasPrefix(rel, "..") {
				continue
			}

			b.ParentPackageName = b2.ImportPackageName
			if b.ParentPackageName == "" {
				b.ParentPackageName = rootPackageName
			}

			b.RouteGroupName = b.ImportPackageName
			b.HasParent = true
			b.Endpoint = rel + "/"

			if !strings.HasSuffix(b.EndpointPath, "/") {
				b.EndpointPath = b.EndpointPath + "/"
			}
		}
	}

	for _, b := range bootstrapTemplates {
		if b.ImportPackageName == filepath.Base(b.PackagePath) {
			b.ImportPackageName = ""
		}
	}

	bootstrapFilePath := filepath.Join(rootPath+"/", "bootstrap_gen.go")
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

func parsePackages(path, endpointBase string, endpointParams []string) ([]*ControllerTemplate, error) {
	replaceRule := regexp.MustCompile(`:(.*?)(/|$)`)
	routes := make(map[string][]*ControllerTemplate)
	structPair, err := findStructPairList(path, endpointParams)
	if err != nil {
		return nil, err
	}
	for cn, s := range structPair {
		req := s.Request
		res := s.Response
		structFilePath := s.FileName
		fileName := filepath.Base(structFilePath)
		fileName = fileName[:len(fileName)-len(filepath.Ext(structFilePath))]

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
		createFileName := strcase.ToSnake(cn) + "_controller_gen.go"
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

		if strings.HasPrefix(fileName, "0_") {
			param := s.LastParam
			endpoint = fmt.Sprintf(":%s", param)
		}

		fullEndpoint := "/" + filepath.Join(endpointBase, endpoint)
		fullEndpoint = replaceRule.ReplaceAllString(fullEndpoint, "{$1}$2")

		ct := &ControllerTemplate{
			Package:               req.PackageName,
			ControllerName:        fmt.Sprintf("%sController", cn),
			ControllerNameInitial: strings.ToLower(string([]rune(cn)[0])),
			HandlerName:           cn,
			RawEndpointPath:       fullEndpoint,
			Endpoint:              endpoint,
			HTTPMethod:            strings.ToUpper(httpMethod),
			RequestStructName:     req.StructName,
			ResponseStructName:    res.StructName,
			RequestParams:         req.RequestParams,
		}

		routes[createDir] = append(routes[createDir], ct)

		err = createFromTemplate("/controller_template.go.tmpl", createPath, ct, false)
		if err != nil {
			return nil, err
		}
	}

	controllers := make([]*ControllerTemplate, 0)
	for dir, cs := range routes {
		routePath := filepath.Join(dir+"/", "routes_gen.go")
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
