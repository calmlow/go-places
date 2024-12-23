package config

import (
	"log"
	"testing"
)

func Test_should_read_local_config_file(t *testing.T) {
	file, err := ReadYamlConfigFile("../../assets/test-files/test-local-config.yaml")
	if err != nil {
		log.Panicf("Failed unit test %v", err)
	}

	expected := []string{"repo-selector", "other-repo"}
	for i, place := range file.Places {
		actualPlaceName := place.Name
		actualShortcut := place.ShortcutAsRune()
		actualDescription := place.Description
		actualPath := place.Path
		actualDocsUrl := place.DocsUrl

		if actualPlaceName != expected[i] {
			t.Errorf("Test failed, expected: '%s' but got: '%s'", expected[i], actualPlaceName)
		}
		if actualDescription != "This tool is what you are viewing right now" {
			t.Errorf("Test failed, 'description' not expected: '%s'", actualDescription)
		}
		if actualShortcut != '_' {
			t.Errorf("Test failed, 'shortcut' not expected: '%v'", actualShortcut)
		}
		if actualPath != "/home/some-repo" {
			t.Errorf("Test failed, 'path' not expected: '%s'", actualPath)
		}
		if actualDocsUrl != "test docs url" {
			t.Errorf("Test failed, 'docs-url' not expected: '%s'", actualDocsUrl)
		}
	}
}
