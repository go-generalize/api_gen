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
	return map[interface{}]interface{}{
		g.generateType(obj.Key): g.generateType(obj.Value),
	}
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
