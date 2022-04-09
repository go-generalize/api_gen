// Package go2json ...
package go2json

import (
	"sort"
	"strings"

	"github.com/go-generalize/go-easyparser/types"
	"github.com/go-utils/count"
)

func (g *Generator) generateArray(typ *types.Array, packageStack []string, exitRecursion bool) interface{} {
	item := g.generateType(typ.Inner, packageStack, exitRecursion)
	return []interface{}{item, item, item}
}

func (g *Generator) generateMap(obj *types.Map, packageStack []string, exitRecursion bool) interface{} {
	key := g.generateType(obj.Key, packageStack, exitRecursion)
	val := g.generateType(obj.Value, packageStack, exitRecursion)
	switch obj.Key.(type) {
	case *types.String:
		return map[string]interface{}{
			key.(string): val,
		}
	case *types.Number:
		return map[int]interface{}{
			key.(int): val,
		}
	}
	panic("map key is `string` or `number`")
}

func (g *Generator) generateObject(obj *types.Object, packageStack []string, exitRecursion bool) interface{} {
	sp := strings.Split(obj.Name, ".")
	name := sp[len(sp)-1]

	cnt, err := count.Do(packageStack, name)
	if err != nil {
		panic("unexpected error")
	}

	if cnt > 2 {
		exitRecursion = true
	}

	packageStack = append(packageStack, name)

	type entry struct {
		Name     string
		Type     types.Type
		Optional bool
	}

	entries := make([]*entry, 0, len(obj.Entries))

	for k, v := range obj.Entries {
		if exitRecursion && v.Optional {
			continue
		}

		entries = append(entries, &entry{
			Name:     k,
			Type:     v.Type,
			Optional: v.Optional,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	if len(entries) == 0 {
		return map[string]interface{}{}
	}

	r := make(map[string]interface{})
	for _, e := range entries {
		r[e.Name] = g.generateType(e.Type, packageStack, exitRecursion)
	}

	return r
}
