package migration

import "sort"

type Migrations []Migration

func (migrations Migrations) SortByID() {
	sort.Slice(
		migrations,
		func(i, j int) bool {
			return migrations[i].ID < migrations[j].ID
		},
	)
}
