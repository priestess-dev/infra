package cache

import (
	"github.com/priestess-dev/infra/utils/fs"
	"regexp"
)

func GetYAMLConfigPaths(path string) ([]string, error) {
	re := regexp.MustCompile(`\w+\.(yaml|yml)`)
	return fs.ReadFilesWithPattern(path, re)
}
