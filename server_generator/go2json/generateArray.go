// Package go2json ...
package go2json

import (
	"fmt"
	"sort"
	"strings"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
	"github.com/go-utils/count"
)

func (g *Generator) generateArray(typ *tstypes.Array, packageStack []string) interface{} {
	return []interface{}{
		g.generateType(typ.Inner, packageStack),
		g.generateType(typ.Inner, packageStack),
		g.generateType(typ.Inner, packageStack),
	}
}

func (g *Generator) generateMap(obj *tstypes.Map, packageStack []string) interface{} {
	key := g.generateType(obj.Key, packageStack)
	val := g.generateType(obj.Value, packageStack)
	switch obj.Key.(type) {
	case *tstypes.String:
		return map[string]interface{}{
			key.(string): val,
		}
	case *tstypes.Number:
		return map[int]interface{}{
			key.(int): val,
		}
	}
	panic("map key is `string` or `number`")
}

func (g *Generator) generateObject(obj *tstypes.Object, packageStack []string) interface{} {
	sp := strings.Split(obj.Name, ".")
	name := sp[len(sp)-1]

	cnt, err := count.Do(packageStack, name)
	if err != nil {
		panic("unexpected error")
	}

	if cnt > 2 {
		return fmt.Sprintf("%s INFINITY LOOP", name)
	}

	packageStack = append(packageStack, name)

	type entry struct {
		Name     string
		Type     tstypes.Type
		Optional bool
	}

	entries := make([]*entry, 0, len(obj.Entries))

	for k, v := range obj.Entries {
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
		if e.Optional {
			r[e.Name] = nil
		} else {
			r[e.Name] = g.generateType(e.Type, packageStack)
		}
	}

	return r
}
