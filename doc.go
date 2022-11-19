package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func WriteDocument(dst string, v any) error {
	ext := filepath.Ext(dst)
	switch ext {
	case ".json":
		data, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(dst, data, 0644)
	case ".yaml", ".yml":
		data, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		return os.WriteFile(dst, data, 0644)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func ReadDocument(src string, v any) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	ext := filepath.Ext(src)
	switch ext {
	case ".json":
		return json.Unmarshal(data, v)
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, v)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func IsDocument(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".json", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func YamlToJson(in []byte) ([]byte, error) {
	out := make(map[string]any)
	err := yaml.Unmarshal(in, &out)
	if err != nil {
		return nil, fmt.Errorf("error un marshalling yaml: %w", err)
	}
	return json.Marshal(out)
}
