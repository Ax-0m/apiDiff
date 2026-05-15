// way to handle unknown json structure is: map[string]interface{}

package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

func ParseFile(filename string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not read the file %s: %w", filename, err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("Could not parse the file %s: %w", filename, err)
	}

	return result, nil
}
