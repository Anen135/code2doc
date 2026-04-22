package scanner

import (
	"fmt"
	"os"
	"path/filepath"

	"code2doc/internal/logger"
)

type ProjectFile struct {
	Path    string
	Content string
}

func Scan(root string, extension string) ([]ProjectFile, error) {
	var files []ProjectFile

	logger.Info(fmt.Sprintf("Scanning root directory: %s for .%s files", root, extension))

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Warning(fmt.Sprintf("Error accessing path %s: %v", path, err))
			return err
		}

		if info.IsDir() || isHidden(path) {
			if info.IsDir() && isHidden(path) {
				logger.Debug(fmt.Sprintf("Skipping hidden directory: %s", path))
				return filepath.SkipDir
			}
			return nil
		}

		if filepath.Ext(path) == extension {
			content, err := os.ReadFile(path)
			if err != nil {
				logger.Error(fmt.Sprintf("Failed to read file %s: %v", path, err))
				return nil
			}
			files = append(files, ProjectFile{
				Path:    path,
				Content: string(content),
			})
			logger.Debug(fmt.Sprintf("Found file: %s", path))
		}
		return nil
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Error during directory walk: %v", err))
	} else {
		logger.Info(fmt.Sprintf("Scan complete. Found %d files", len(files)))
	}

	return files, err
}

func isHidden(path string) bool {
	base := filepath.Base(path)
	return base != "." && base[0] == '.'
}
