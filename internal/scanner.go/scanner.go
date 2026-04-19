package scanner

import (
	"os"
	"path/filepath"
)

type ProjectFile struct {
	Path    string
	Content string
}

func Scan(root string, extension string) ([]ProjectFile, error) {
	var files []ProjectFile

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || isHidden(path) {
			if info.IsDir() && isHidden(path) {
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(path) == extension {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			files = append(files, ProjectFile{
				Path:    path,
				Content: string(content),
			})
		}
		return nil
	})

	return files, err
}

func isHidden(path string) bool {
	base := filepath.Base(path)
	return base != "." && base[0] == '.'
}
