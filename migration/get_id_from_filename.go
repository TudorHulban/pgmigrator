package migration

import (
	"path/filepath"
	"strings"
)

type GetIDFromFilename func(filename string) string

// "0001__initial" == IDFromFilename("V0001__initial.sql")
func IDFromFilename(filename string) string {
	return strings.TrimPrefix(
		strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename)),
		"V",
	)
}
