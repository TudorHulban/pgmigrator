package migration

import (
	"path/filepath"
	"strings"
)

type GetIDFromFilename func(filename string) string

// "0001__initial" == IDFromFilename("V0001__initial.sql")
func IDFromFilename(filename string) string {
	base := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))

	if strings.HasPrefix(base, "V") {
		base = base[1:]
	}

	return base
}
