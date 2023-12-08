package cache

import (
	"regexp"
	"testing"
)

func TestYAMLRegex(t *testing.T) {
	// regex for yaml config file
	re := regexp.MustCompile(`\w+\.(yaml|yml)`)
	trueSet := []string{
		"config.yaml",
		"config.yml",
		"test.yaml",
		"Config.yml",
	}
	for _, s := range trueSet {
		if !re.MatchString(s) {
			t.Errorf("yaml regex failed")
		}
	}
	falseSet := []string{
		"config.json",
		"config.toml",
	}
	for _, s := range falseSet {
		if re.MatchString(s) {
			t.Errorf("yaml regex failed")
		}
	}
}

func TestGetYAMLConfigPaths(t *testing.T) {
	path := "./tmp"
	confs, err := GetYAMLConfigPaths(path)
	if err != nil {
		t.Errorf("get yaml config paths failed")
	}
	for _, conf := range confs {
		t.Logf("%s\n", conf)
	}
}
