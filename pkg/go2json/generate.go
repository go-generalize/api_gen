// Package go2json ...
package go2json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-generalize/go-easyparser/types"
	"github.com/go-generalize/go-easyparser/util"
	"github.com/iancoleman/strcase"
	"golang.org/x/xerrors"
)

var (
	fixedTime = time.Date(
		2006, time.January, 02, 15, 04, 05, 0, time.Local,
	).Format(time.RFC3339)
)

// Generator go struct->json
type Generator struct {
	types   map[string]types.Type
	altPkgs map[string]string

	BasePackage     string
	CustomGenerator func(t types.Type, packageStack []string, exitRecursion bool) (bool, interface{})
}

// NewGenerator Generator constructor
func NewGenerator(types map[string]types.Type) *Generator {
	return &Generator{
		types:   types,
		altPkgs: make(map[string]string),
	}
}

// Generate go struct -> json
func (g *Generator) Generate(dir string) error {
	type entry struct {
		key string
		typ types.Type
	}

	entries := make([]*entry, 0, len(g.types))
	for k, v := range g.types {
		entries = append(entries, &entry{
			key: k,
			typ: v,
		})
	}
	sort.Slice(entries, func(i, j int) bool {
		left := strings.HasPrefix(entries[i].key, g.BasePackage+".")
		right := strings.HasPrefix(entries[j].key, g.BasePackage+".")

		if left && !right {
			return true
		}
		if !left && right {
			return false
		}

		return entries[i].key < entries[j].key
	})

	used := map[string]struct{}{}
	for i, e := range entries {
		obj, ok := e.typ.(*types.Object)

		if !ok {
			continue
		}

		pkg, strct := util.SplitPackegeStruct(obj.Name)
		if _, ok := used[strct]; !ok {
			g.altPkgs[obj.Name] = strct
			used[strct] = struct{}{}
			continue
		}

		p := util.GetPackageNameFromPath(pkg)
		name := strcase.ToCamel(p + "_" + strct)

		if _, ok := used[name]; !ok {
			g.altPkgs[obj.Name] = name
			used[name] = struct{}{}
			continue
		}

		name = fmt.Sprintf("%s%03d", name, i)
		g.altPkgs[obj.Name] = name
		used[name] = struct{}{}
	}

	typeSets := make(map[string]*GenerateSet)
	for n, t := range g.types {
		_, structName := util.SplitPackegeStruct(n)

		if strings.HasSuffix(structName, "Request") {
			requestName := strings.TrimSuffix(structName, "Request")
			if _, ok := typeSets[requestName]; ok {
				typeSets[requestName].Request = t
				continue
			}
			typeSets[requestName] = &GenerateSet{
				Request:  t,
				Response: nil,
			}
		} else if strings.HasSuffix(structName, "Response") {
			requestName := strings.TrimSuffix(structName, "Response")
			if _, ok := typeSets[requestName]; ok {
				typeSets[requestName].Response = t
				continue
			}
			typeSets[requestName] = &GenerateSet{
				Request:  nil,
				Response: t,
			}
		}
	}

	for c, gs := range typeSets {
		packageStack := make([]string, 0)
		req := g.generateType(gs.Request, packageStack, false)
		packageStack = make([]string, 0)
		res := g.generateType(gs.Response, packageStack, false)

		gt := &GenerateType{
			Meta: &MockMeta{
				Status:       200,
				MatchRequest: req,
			},
			Payload: res,
		}
		jsonByte, err := json.MarshalIndent(gt, "", "    ")
		if err != nil {
			return xerrors.Errorf("GenerateType Marshal error at %s: %w", c, err)
		}
		// append \n
		jsonByte = append(jsonByte, 0xa)

		httpMethod := ""
		httpMethods := []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		}
		uc := strings.ToUpper(c)
		for _, h := range httpMethods {
			if !strings.HasPrefix(uc, h) {
				continue
			}
			httpMethod = strings.ToLower(h)
			break
		}

		rc := []rune(c)
		opName := strcase.ToSnake(string(rc[len(httpMethod):]))
		if len(opName) > 0 {
			opName = "_" + opName
		}
		opJSONDir := filepath.Join(dir+"/", fmt.Sprintf("%s%s/", httpMethod, opName))
		{
			if i, err := os.Stat(opJSONDir); err == nil {
				if !i.IsDir() {
					return xerrors.Errorf("%s must be directory.", opJSONDir)
				}
			} else {
				err = os.MkdirAll(opJSONDir, 0775)
				if err != nil {
					return xerrors.Errorf("Failed to create %s: %w", opJSONDir, err)
				}
			}
		}

		jsonPath := filepath.Join(opJSONDir, "/default.json")
		err = os.WriteFile(jsonPath, jsonByte, 0664)
		if err != nil {
			return xerrors.Errorf("Write mock json error in %s: %w", jsonPath, err)
		}
	}

	return nil
}

func (g *Generator) generateType(t types.Type, packageStack []string, exitRecursion bool) interface{} {
	if t == nil {
		return nil
	}

	if g.CustomGenerator != nil {
		handled, r := g.CustomGenerator(t, packageStack, exitRecursion)

		if handled {
			return r
		}
	}

	switch v := t.(type) {
	case *types.Array:
		return g.generateArray(v, packageStack, exitRecursion)
	case *types.Object:
		return g.generateObject(v, packageStack, exitRecursion)
	case *types.String:
		return "string"
	case *types.Number:
		return 1234
	case *types.Boolean:
		return true
	case *types.Date:
		return fixedTime
	case *types.Nullable:
		if exitRecursion {
			return nil
		}

		return g.generateType(v.Inner, packageStack, false)
	case *types.Any:
		return "any"
	case *types.Map:
		return g.generateMap(v, packageStack, exitRecursion)
	default:
		panic("unsupported")
	}
}
