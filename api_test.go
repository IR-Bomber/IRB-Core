package main

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/IR-Bomber/IRB-Core/utils"
)


// TestAPI tests the fetchAPI function and validates the API structure
func TestAPI(t *testing.T) {
	// Fetch the data from the API
	err, result := utils.FetchAPI()
	if err != nil {
		t.Errorf("Error with API: %v", err)
		return
	}

	// Check that we received a valid response
	if len(result) == 0 {
		t.Errorf("API response is empty")
		return
	}

	// Now check the keys of each API
	for i, api := range result {
		apiMap, ok := utils.FormatJSON(api, "9123456789").(map[string]interface{})
		if !ok {
			t.Errorf("API %d: Expected map for API data, but got %T", i, api)
			continue
		}

		// Check if "method" exists and is a valid HTTP method
		if method, exists := apiMap["method"]; exists {
			if !isValidHTTPMethod(method) {
				t.Errorf("API %d: '%s' is not a valid HTTP method", i, method)
			}
		} else {
			t.Errorf("API %d: 'method' is missing", i)
		}

		// Check if "url" exists and is a valid URL
		if urlStr, exists := apiMap["url"]; exists {
			if !isValidURL(urlStr) {
				t.Errorf("API %d: '%s' is not a valid URL", i, urlStr)
			}
		} else {
			t.Errorf("API %d: 'url' is missing", i)
		}

		// Check if "payload" exists and is of the correct type (if not nil)
		if payload, exists := apiMap["payload"]; exists {
			checkKeyType(t, "payload", payload, "map[string]interface {}", i)
		}

		// Check if "data" exists and is of the correct type (if not nil)
		if data, exists := apiMap["data"]; exists {
			checkKeyType(t, "data", data, "string", i)
		}
	}
}

// checkKeyType checks if the key's value is of the expected type
func checkKeyType(t *testing.T, key string, value interface{}, expectedType string, apiIndex int) {
	if value == nil {
		// No error if the value is nil
		return
	}

	actualType := reflect.TypeOf(value).String()

	if actualType != expectedType {
		t.Errorf("API %d: '%s' should be of type '%s' but got '%s'.", apiIndex, key, expectedType, actualType)
	}
}

// isValidHTTPMethod checks if the given method is a valid HTTP method
func isValidHTTPMethod(method interface{}) bool {
	validMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
	for _, validMethod := range validMethods {
		if method == validMethod {
			return true
		}
	}
	return false
}

// isValidURL checks if the given string is a valid URL
func isValidURL(urlStr interface{}) bool {
	str, ok := urlStr.(string)
	if !ok {
		return false
	}

	_, err := url.ParseRequestURI(str)
	return err == nil
}