package common

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func GetSetDifference(key string, d *schema.ResourceData, f schema.SchemaSetFunc, supportUpdate bool) (add *schema.Set, remove *schema.Set, modify *schema.Set, cache map[int]interface{}) {
	if d.HasChange(key) {
		ov, nv := d.GetChange(key)
		if ov == nil {
			ov = new(schema.Set)
		}
		if nv == nil {
			nv = new(schema.Set)

		}
		os := ov.(*schema.Set)
		ns := nv.(*schema.Set)
		cache = make(map[int]interface{})
		addProbably := schema.NewSet(f, ns.Difference(os).List())
		for _, entry := range addProbably.List() {
			index := f(entry)
			cache[index] = entry
		}
		removeProbably := schema.NewSet(f, os.Difference(ns).List())

		add = addProbably.Difference(removeProbably)
		remove = removeProbably.Difference(addProbably)

		if supportUpdate {
			modify = removeProbably.Difference(add)
		}

		return add, remove, modify, cache
	}
	return add, remove, modify, cache
}
