package mjson

import (
	"encoding/json"
)

// Merge merges two deep nested JSON objects and returns the merged result.
func Merge(jsonStr1, jsonStr2 string) (string, error) {
	// Parse the first JSON string into a map structure
	var map1 map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr1), &map1)
	if err != nil {
		return "", err
	}

	// Parse the second JSON string into a map structure
	var map2 map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr2), &map2)
	if err != nil {
		return "", err
	}

	// Merge the two map structures
	mergedMap := mergeMaps(map1, map2)

	// Convert the merged map structure back to JSON string
	mergedJSON, err := json.Marshal(mergedMap)
	if err != nil {
		return "", err
	}

	return string(mergedJSON), nil
}

// mergeMaps merges two map structures recursively
func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	mergedMap := make(map[string]interface{})

	// Iterate over the keys of the first map structure
	// and copy the key-value pairs to the merged map
	for key, value := range map1 {
		mergedMap[key] = value
	}

	// Iterate over the keys of the second map structure
	// and copy the key-value pairs to the merged map.
	// If a key already exists in the merged map, recursively merge the values.
	for key, value := range map2 {
		if existingValue, ok := mergedMap[key]; ok {
			// If the key already exists, recursively merge the values
			mergedMap[key] = mergeValues(existingValue, value)
		} else {
			// If the key doesn't exist, simply copy the value
			mergedMap[key] = value
		}
	}

	return mergedMap
}

// mergeValues merges two values recursively
func mergeValues(value1, value2 interface{}) interface{} {
	switch value1 := value1.(type) {
	case map[string]interface{}:
		// If the value is a map structure, recursively merge the two map structures
		if value2, ok := value2.(map[string]interface{}); ok {
			return mergeMaps(value1, value2)
		}
	case []interface{}:
		// If the value is a slice, merge the two slices
		if value2, ok := value2.([]interface{}); ok {
			return append(value1, value2...)
		}
	}

	// For other cases, simply use the value from the second object to overwrite the value from the first object
	return value2
}
