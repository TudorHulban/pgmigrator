package migration

import (
	"fmt"
	"io/fs"
	"strings"
)

func Load(directory fs.FS, getID GetIDFromFilename) (Migrations, error) {
	var result []Migration

	if err := fs.WalkDir(
		directory,
		".",
		func(path string, d fs.DirEntry, errReadFile error) error {
			if errReadFile != nil {
				return errReadFile
			}

			if !strings.HasSuffix(path, ".sql") {
				return nil
			}

			content, errReadContent := fs.ReadFile(directory, path)
			if errReadContent != nil {
				return errReadContent
			}

			result = append(
				result,
				Migration{
					ID:  getID(d.Name()),
					SQL: string(content),
				},
			)

			return nil
		},
	); err != nil {
		return nil,
			fmt.Errorf("load: %w", err)
	}

	return result, nil
}
