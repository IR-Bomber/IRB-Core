package utils

import (
	"fmt"
	"strings"
)




func ValidateStatus(status int) bool {
	return status >= 200 && status < 300
}

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