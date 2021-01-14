// Package main ...
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
	"github.com/go-generalize/api_gen/server_generator/go2json"
	_ "github.com/go-generalize/api_gen/server_generator/statik"
	go2tsparser "github.com/go-generalize/go2ts/pkg/parser"
	"github.com/iancoleman/strcase"
	"github.com/rakyll/statik/fs"
	"golang.org/x/xerrors"
)

var (
	supportHTTPMethod = []string{
		"get",
		"post",
		"put",
		"delete",
		"patch",
	}
	withMock bool
)

var (
	replaceRule              = regexp.MustCompile(`:(.*?)(/|$)`)
	endpointReplaceMatchRule = regexp.MustCompile(`:(.*?)/`)
)

const rootPackageName = "root"

func main() {
	l := len(os.Args)
	if l < 2 {
		fmt.Println("Specify the target directory.")
		os.Exit(1)
	}

	versionFlag := flag.Bool("v", false, "print version")
	updateCheckFlag := flag.Bool("check-update", false, "check for updates")
	mockFlag := flag.Bool("mock", false, "generate mock")

	flag.Parse()

	if *versionFlag {
		fmt.Println(common.AppVersion)
		return
	}

	if *updateCheckFlag {
		common.CheckUpdate()
		return
	}

	if *mockFlag {
		withMock = true
	}

	args := os.Args
	err := run(args[len(args)-1])
	if err != nil {
		log.Fatal(err)
	}
}

func run(arg string) error {
	const (
		commonPropsDir    = "props"
		commonWrapperDir  = "wrapper"
		commonInternalDir = "internal"
	)

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

	var apiRootPackage string
	var apiRootPathRel string
	{
		r, err := filepath.Rel(packageRootPath, rootPath)
		if err != nil {
			return err
		}
		apiRootPathRel = r
		apiRootPackage = filepath.Join(basePackagePath+"/", r)
	}

	var controllerPropsPackage string
	{
		dir := filepath.Join(rootPath, commonPropsDir)
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}

		r, err := filepath.Rel(packageRootPath, dir)

		if err != nil {
			return err
		}

		controllerPropsPackage = filepath.Join(basePackagePath+"/", r)

		err = createFromTemplate(
			"/templates/controller_props.go.tmpl",
			filepath.Join(dir, "controller_props.go"),
			nil, false, nil,
		)

		if err != nil {
			return err
		}
	}

	var wrapperInternalPackage string
	{
		dir := filepath.Join(rootPath, commonWrapperDir, commonInternalDir)
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}

		r, err := filepath.Rel(packageRootPath, dir)

		if err != nil {
			return err
		}

		wrapperInternalPackage = filepath.Join(basePackagePath+"/", r)

		err = createFromTemplate(
			"/templates/fmt_template.go.tmpl",
			filepath.Join(dir, "fmt.go"),
			map[string]string{
				"AppVersion": common.AppVersion,
			}, true, nil,
		)

		if err != nil {
			return err
		}
	}

	var wrapperPackage string
	{
		dir := filepath.Join(rootPath, commonWrapperDir)
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}

		r, err := filepath.Rel(packageRootPath, dir)

		if err != nil {
			return err
		}

		wrapperPackage = filepath.Join(basePackagePath+"/", r)

		err = createFromTemplate(
			"/templates/wrapper_template.go.tmpl",
			filepath.Join(dir, "wrapper.go"),
			map[string]string{
				"AppVersion":      common.AppVersion,
				"InternalPackage": wrapperInternalPackage,
			}, true, nil,
		)

		if err != nil {
			return err
		}
	}

	isExistRoot := false
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		if p, _ := filepath.Rel(rootPath, path); p == commonPropsDir {
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
			endpoint = "/"
			endpointPath = "/"
		} else {
			endpointPath = "/" + endpointPath
		}

		endpointParams := make([]string, 0)
		rawEndpointFilePath := endpointPath
		endpointPath = strings.ReplaceAll(endpointPath, "/_", "/:")
		endpointParam := strings.ReplaceAll(endpoint, "/_", "/:")
		endPointParamsFromRegexp := endpointReplaceMatchRule.FindAllStringSubmatch(endpointPath+"/", -1)

		for _, e := range endPointParamsFromRegexp {
			endpointParams = append(endpointParams, e[1])
		}

		cs, err := parsePackages(
			path, endpointParam, endpointParams, controllerPropsPackage, wrapperPackage,
		)
		if err != nil {
			return err
		}
		if len(cs) == 0 {
			return nil
		}
		if endpointPath == "/" {
			packageName = cs[0].Package
			isExistRoot = true

			packagePath = ""
			importPackageName = ""
		}

		bootstrapTemplates = append(bootstrapTemplates, &BootstrapTemplates{
			PackagePath:         packagePath,
			ImportPackageName:   strcase.ToLowerCamel(importPackageName),
			EndpointPath:        endpointPath,
			RawEndpointFilePath: rawEndpointFilePath,
			Endpoint:            endpoint,
			Controller:          cs[0],
		})

		if withMock {
			// create mock json
			if err = createMockJSON(rootPath, path); err != nil {
				return xerrors.Errorf("error in createMockJSON method: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if !isExistRoot {
		packageName = filepath.Base(rootPath)
		bootstrapTemplates = append([]*BootstrapTemplates{
			{
				PackagePath:       "",
				ImportPackageName: "",
				EndpointPath:      "/",
				Endpoint:          "/",
				Controller:        nil,
			},
		}, bootstrapTemplates...)
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
	bootstraptemplate := &BootstrapTemplate{
		AppVersion:             common.AppVersion,
		PackageName:            packageName,
		Bootstraps:             bootstrapTemplates,
		ControllerPropsPackage: controllerPropsPackage,
	}
	err = createFromTemplate(
		"/templates/bootstrap_template.go.tmpl",
		bootstrapFilePath, bootstraptemplate,
		true, template.FuncMap{
			"GetGroupName": getGroupName,
			"GetNewRoute":  getNewRoute,
		})
	if err != nil {
		return err
	}

	if withMock {
		if err = createMock(&CreateMockRequest{
			RootPath:               rootPath,
			ControllerPropsPackage: controllerPropsPackage,
			APIRootPackage:         apiRootPackage,
			BootstrapTemplate:      bootstraptemplate,
			APIRootPathRel:         apiRootPathRel,
		}); err != nil {
			return xerrors.Errorf("error in createMock method: %w", err)
		}
	}

	return nil
}

func parsePackages(
	path, endpointBase string,
	endpointParams []string,
	controllerPropsPackage string,
	wrapperPackage string,
) ([]*ControllerTemplate, error) {
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

		fullEndpoint := filepath.Join("/", endpointBase, endpoint)
		fullEndpoint = replaceRule.ReplaceAllString(fullEndpoint, "{$1}$2")

		ct := &ControllerTemplate{
			AppVersion:             common.AppVersion,
			Package:                req.PackageName,
			ControllerName:         fmt.Sprintf("%sController", cn),
			ControllerNameInitial:  strings.ToLower(string([]rune(cn)[0])),
			HandlerName:            cn,
			RawEndpointPath:        fullEndpoint,
			Endpoint:               endpoint,
			HTTPMethod:             strings.ToUpper(httpMethod),
			RequestStructName:      req.StructName,
			ResponseStructName:     res.StructName,
			RequestParams:          req.RequestParams,
			ControllerPropsPackage: controllerPropsPackage,
			WrapperPackage:         wrapperPackage,
		}

		routes[createDir] = append(routes[createDir], ct)

		err = createFromTemplate("/templates/controller_template.go.tmpl",
			createPath, ct, false, template.FuncMap{})
		if err != nil {
			return nil, err
		}
	}

	controllers := make([]*ControllerTemplate, 0)
	for dir, cs := range routes {
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

		// routes_gen.go
		{
			routePath := filepath.Join(dir+"/", "routes_gen.go")
			err := createFromTemplate("/templates/routes_template.go.tmpl", routePath, &RoutesTemplate{
				AppVersion:             common.AppVersion,
				Package:                packageName,
				Controllers:            cs,
				ControllerPropsPackage: controllerPropsPackage,
				WrapperPackage:         wrapperPackage,
			}, true, template.FuncMap{
				"GetJSONDir": getJSONDir,
			})
			if err != nil {
				return nil, err
			}
		}

		// mock_routes_gen.go
		if withMock {
			routePath := filepath.Join(dir+"/", "mock_routes_gen.go")
			err := createFromTemplate("/templates/mock_routes_template.go.tmpl", routePath, &RoutesTemplate{
				AppVersion:             common.AppVersion,
				Package:                packageName,
				Controllers:            cs,
				ControllerPropsPackage: controllerPropsPackage,
			}, true, template.FuncMap{
				"GetJSONDir": getJSONDir,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return controllers, nil
}

func createMock(req *CreateMockRequest) error {
	// cmd/mock/main.go
	{
		mockPath := filepath.Join(req.RootPath+"/", "/cmd/mock/")
		if i, err := os.Stat(mockPath); err == nil {
			if !i.IsDir() {
				return xerrors.Errorf("%s is must be directory: %w", mockPath, err)
			}
		} else {
			err = os.MkdirAll(mockPath, 0775)
			if err != nil {
				return xerrors.Errorf("Failed create an %s: %w", mockPath, err)
			}
		}

		mockMainPath := filepath.Join(mockPath, "main.go")
		err := createFromTemplate("/templates/mock_main.go.tmpl", mockMainPath, &MockMainTemplate{
			AppVersion:             common.AppVersion,
			APIPackageRoot:         req.APIRootPackage,
			APIRootPackageName:     filepath.Base(req.APIRootPackage),
			ControllerPropsPackage: req.ControllerPropsPackage,
			DefaultJSONDirPath:     filepath.Join(req.APIRootPathRel+"/", "mock_jsons/"),
		}, true, template.FuncMap{})
		if err != nil {
			return xerrors.Errorf("Failed create an %s: %w", mockMainPath, err)
		}
	}

	// mock_bootstrap_gen.go
	{
		mockBootstrapPath := filepath.Join(req.RootPath+"/", "mock_bootstrap_gen.go")
		err := createFromTemplate("/templates/mock_bootstrap_template.go.tmpl", mockBootstrapPath,
			req.BootstrapTemplate,
			true, template.FuncMap{
				"GetGroupName":    getGroupName,
				"GetNewMockRoute": getNewMockRoute,
			},
		)
		if err != nil {
			return xerrors.Errorf("Failed create an %s: %w", mockBootstrapPath, err)
		}
	}

	return nil
}

func createMockJSON(rootPath, path string) error {
	f, err := go2tsparser.NewParser(path, func(opt *go2tsparser.FilterOpt) bool {
		if opt.Dependency {
			return true
		}
		if !opt.BasePackage {
			return false
		}
		if !opt.Exported {
			return false
		}

		return strings.HasSuffix(opt.Name, "Request") || strings.HasSuffix(opt.Name, "Response")
	})
	if err != nil {
		return err
	}

	r, err := filepath.Rel(rootPath, path)
	if err != nil {
		return err
	}

	jsonRoot := filepath.Join(rootPath+"/", "mock_jsons/")
	jsonDir := filepath.Join(jsonRoot, r)
	i, err := os.Stat(jsonDir)
	if err == nil {
		if !i.IsDir() {
			return xerrors.Errorf("%s must be dir", jsonDir)
		}
	} else {
		err = os.MkdirAll(jsonDir, 0775)
		if err != nil {
			return xerrors.Errorf("failed to create %s: %w", jsonDir, err)
		}
	}

	ts, err := f.Parse()
	if err != nil {
		return err
	}
	err = go2json.NewGenerator(ts).Generate(jsonDir)
	return err
}

func createFromTemplate(templatePath, path string, m interface{}, isOverRide bool, funcMap template.FuncMap) error {
	_, err := os.Stat(path)
	if err == nil {
		if !isOverRide {
			return nil
		}

		err = os.Remove(path)
		if err != nil {
			return xerrors.Errorf("%s remove error: %w", path, err)
		}
	}

	statikFs, err := fs.New()
	if err != nil {
		return xerrors.Errorf("statikFs init error: %w", err)
	}
	f, err := statikFs.Open(templatePath)
	if err != nil {
		return xerrors.Errorf("%s open error in statikFs: %w", templatePath, err)
	}

	t, err := ioutil.ReadAll(f)
	if err != nil {
		return xerrors.Errorf("%s read error in ioutil: %w", templatePath, err)
	}

	tpl := template.Must(template.New("tmpl").Funcs(funcMap).Parse(string(t)))

	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return xerrors.Errorf("%s create error: %w", path, err)
	}
	defer w.Close()

	err = tpl.Execute(w, m)
	if err != nil {
		return xerrors.Errorf("template exec error in %s: %w", path, err)
	}

	_, err = ExecCommand("goimports", "-w", path)
	if err != nil {
		return xerrors.Errorf("goimports exec error (%s): %w", path, err)
	}

	return nil
}
