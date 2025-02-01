package utils

import (
	"fmt"
	"strings"
)




func ValidateStatus(status int) bool {
	return status >= 200 && status < 300
}

func FormatJSON(data any, args ...any) any {
	switch v := data.(type) {
	case string:
		if strings.Contains(v, "%s") && len(args) > 0 {
			return fmt.Sprintf(v, args...)
		}
		return v
	case map[string]any:
		for key, value := range v {
			v[key] = FormatJSON(value, args...)
		}
	case []any:
		for i, value := range v {
			v[i] = FormatJSON(value, args...)
		}
	}
	return data
}
