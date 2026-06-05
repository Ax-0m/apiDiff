package snapshot

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Save(projectName string, endpoint string, data interface{}) error {
	dir := fmt.Sprintf("snapshots/%s", projectName)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("could not create snapshot directory: %w", err)
	}

	filename := endpointToFileName(endpoint)
	filepath := fmt.Sprintf("%s/%s.json", dir, filename)

	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("could not marshal snapshot data: %w", err)
	}

	err = os.WriteFile(filepath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("could not write snapshot file: %w", err)
	}

	fmt.Println("saved snapshot:", filepath)
	return nil
}

func endpointToFileName(endpoint string) string {
	name := strings.TrimPrefix(endpoint, "/")
	name = strings.ReplaceAll(name, "/", "_")
	return name
}

func Load(projectName string, endpoint string) (interface{}, error) {
	filename := endpointToFileName(endpoint)
	filepath := fmt.Sprintf("snapshots/%s/%s.json", projectName, filename)

	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("could not read snapshot file %s: %w", filepath, err)
	}

	var result interface{}
	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, fmt.Errorf("could not parse snapshot file %s: %w", filepath, err)
	}

	return result, nil
}
