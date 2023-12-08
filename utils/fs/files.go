package fs

import (
	"os"
	"regexp"
)

// ReadFilesWithPattern reads all files in path that match the given pattern
func ReadFilesWithPattern(path string, re *regexp.Regexp) ([]string, error) {
	files, err := os.ReadDir(path)
	result := make([]string, 0)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if re.MatchString(f.Name()) {
			result = append(result, f.Name())
		}
	}
	return result, nil
}

// ReadFilesWithSuffix reads all files in path that have the given suffix
func ReadFilesWithSuffix(path string, suffix string) ([]string, error) {
	re := regexp.MustCompile(`\w+\.` + suffix)
	return ReadFilesWithPattern(path, re)
}

// ReadFilesRecursiveWithPattern reads all files in path and its subdirectories that match the given pattern
func ReadFilesRecursiveWithPattern(path string, re *regexp.Regexp) ([]string, error) {
	files, err := os.ReadDir(path)
	result := make([]string, 0)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			subFiles, err := ReadFilesRecursiveWithPattern(path+"/"+f.Name(), re)
			if err != nil {
				return nil, err
			}
			result = append(result, subFiles...)
		} else if re.MatchString(f.Name()) {
			result = append(result, f.Name())
		}
	}
	return result, nil
}
