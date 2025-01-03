package pgmigrator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/TudorHulban/pgmigrator/migration"
	"github.com/stretchr/testify/require"
)

type PGMigrator struct {
	migration.Migrations
}

type ParamsNewPGMigrator struct {
	Directories       []fs.FS
	FilePaths         []string
	GetIDFromFilename migration.GetIDFromFilename
	T                 *testing.T
}

func NewPGMigrator(params *ParamsNewPGMigrator) *PGMigrator {
	var migrations migration.Migrations

	getID := migration.IDFromFilename

	if params.GetIDFromFilename != nil {
		getID = params.GetIDFromFilename
	}

	for _, directory := range params.Directories {
		buf, errLoad := migration.Load(
			directory,
			getID,
		)
		require.NoError(params.T, errLoad)

		migrations = append(migrations, buf...)
	}

	for _, filePath := range params.FilePaths {
		content, errRead := os.ReadFile(filePath)
		require.NoError(
			params.T,
			errRead,
			fmt.Errorf("failed to read file: %w", errRead),
		)

		filename := filepath.Base(filePath)

		migrations = append(migrations,
			migration.Migration{
				ID:  strings.TrimSuffix(filename, filepath.Ext(filename)),
				SQL: string(content),
			},
		)
	}

	return &PGMigrator{
		Migrations: migrations,
	}
}
