package pgmigrator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type ParamsNewPGMigrator struct {
	Directories []fs.FS
	FilePaths   []string
	T           *testing.T
}

func NewPGMigrator(dirs []fs.FS, filePaths []string, opts ...Option) (*PGMigrator, error) {
	var migrations migration.Migrations

	for _, dir := range dirs {
		buf, errLoad := pgmigrate.Load(dir)
		if errLoad != nil {
			return nil, errLoad
		}

		migrations = append(migrations, buf...)
	}

	for _, filePath := range filePaths {
		content, errRead := os.ReadFile(filePath)
		if errRead != nil {
			return nil, fmt.Errorf("failed to read file: %w", errRead)
		}

		filename := filepath.Base(filePath)

		migrations = append(migrations,
			pgmigrate.Migration{
				ID:  strings.TrimSuffix(filename, filepath.Ext(filename)),
				SQL: string(content),
			},
		)
	}

	pgm := PGMigrator{
		m: pgmigrate.NewMigrator(migrations),
	}

	for _, opt := range opts {
		opt(&pgm)
	}

	return &pgm, nil
}
