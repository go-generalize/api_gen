// Package go2json ...
package go2json

import (
	"sort"

	tstypes "github.com/go-generalize/go2ts/pkg/types"
)

func (g *Generator) generateArray(typ *tstypes.Array) interface{} {
	return []interface{}{
		g.generateType(typ.Inner),
		g.generateType(typ.Inner),
		g.generateType(typ.Inner),
	}
}

func (g *Generator) generateMap(obj *tstypes.Map) interface{} {
	key := g.generateType(obj.Key)
	val := g.generateType(obj.Value)
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

func (g *Generator) generateObject(obj *tstypes.Object) interface{} {
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
			r[e.Name] = g.generateType(e.Type)
		}
	}

	return r
}
