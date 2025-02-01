package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)




func ValidateStatus(status int) bool {
	return status >= 200 && status < 300
}

const (
	APIURL = "https://raw.githubusercontent.com/M-logique/Iran-Bomber-Core/refs/heads/main/data/API.json"
)

// FormatJSON formats the given data by replacing "$num" in strings with provided arguments
func FormatJSON(data any, args ...any) any {
	switch v := data.(type) {
	case string:
		// Check if the string contains "$num" and replace it with the first argument
		if strings.Contains(v, "$num") && len(args) > 0 {
			return strings.Replace(v, "$num", fmt.Sprintf("%v", args[0]), -1)
		}
		return v
	case map[string]any:
		// Recursively process each key-value pair in the map
		for key, value := range v {
			v[key] = FormatJSON(value, args...)
		}
	case []any:
		// Recursively process each element in the array
		for i, value := range v {
			v[i] = FormatJSON(value, args...)
		}
	}
	return data
}


func FetchAPI() (error, []interface{}) {
	resp, err := http.Get(APIURL)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err, nil
	}

	// Check if the "APIs" field is in the correct format
	apis, ok := result["APIs"].([]interface{})
	if !ok {
		return fmt.Errorf("Expected 'APIs' to be an array, but got %T", result["APIs"]), nil
	}

	return nil, apis
}

func LoadAPI(filePath string) (error, []interface{}) {
	fi, err := os.Open(filePath)

	if err != nil {
		return err, nil
	}

	defer fi.Close()

	var result map[string]interface{}

	if err = json.NewDecoder(fi).Decode(&result); err != nil {
		return err, nil
	}

	return nil, result["APIs"].([]interface{})

}