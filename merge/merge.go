package merge

import (
	"reflect"
)

const (
	Concatenate                  = "concatenate"
	RemoveDuplicates             = "remove_duplicates"
	Overwrite                    = "overwrite"
	ReplaceByIndexPreferRight    = "replace_by_index_prefer_right"
	ReplaceByIndexPreferLeft     = "replace_by_index_prefer_left"
	ReplaceByIndexPreferMax      = "replace_by_index_prefer_max"
	ReplaceByIndexPreferRightRec = "replace_by_index_prefer_right_rec"
	ReplaceByIndexPreferLeftRec  = "replace_by_index_prefer_left_rec"
	ReplaceByIndexPreferMaxRec   = "replace_by_index_prefer_max_rec"
)

// MergeMapStructure 合并两个map[string]interface{}
func MergeMapStructure(m1, m2 map[string]interface{}, mergeType string) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		switch mergeType {
		case ReplaceByIndexPreferRightRec, ReplaceByIndexPreferLeftRec, ReplaceByIndexPreferMaxRec:
			if reflect.TypeOf(v).Kind() == reflect.Slice && reflect.TypeOf(result[k]).Kind() == reflect.Slice {
				result[k] = mergeSlice(result[k].([]interface{}), v.([]interface{}), mergeType)
			} else if reflect.TypeOf(v).Kind() == reflect.Map && reflect.TypeOf(result[k]).Kind() == reflect.Map {
				result[k] = MergeMapStructure(result[k].(map[string]interface{}), v.(map[string]interface{}), mergeType)
			} else {
				result[k] = v
			}
		default:
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				if reflect.TypeOf(result[k]).Kind() != reflect.Map {
					result[k] = make(map[string]interface{})
				}
				result[k] = MergeMapStructure(result[k].(map[string]interface{}), v.(map[string]interface{}), mergeType)
			case reflect.Slice:
				if reflect.TypeOf(result[k]).Kind() != reflect.Slice {
					result[k] = make([]interface{}, 0)
				}
				result[k] = mergeSlice(result[k].([]interface{}), v.([]interface{}), mergeType)
			default:
				result[k] = v
			}
		}
	}

	return result
}

// mergeSlice 合并两个Slice
func mergeSlice(s1, s2 []interface{}, mergeType string) []interface{} {
	var result []interface{}

	switch mergeType {
	case Concatenate:
		result = append(result, s1...)
		result = append(result, s2...)
	case RemoveDuplicates:
		result = removeDuplicates(s1, s2)
	case Overwrite:
		result = s2
	case ReplaceByIndexPreferRight, ReplaceByIndexPreferRightRec:
		result = replaceByIndex(s1, s2, len(s2))
	case ReplaceByIndexPreferLeft, ReplaceByIndexPreferLeftRec:
		result = replaceByIndex(s1, s2, len(s1))
	case ReplaceByIndexPreferMax, ReplaceByIndexPreferMaxRec:
		maxLen := len(s1)
		if len(s2) > maxLen {
			maxLen = len(s2)
		}
		result = replaceByIndex(s1, s2, maxLen)
	}

	return result
}

// removeDuplicates 去重
func removeDuplicates(s1, s2 []interface{}) []interface{} {
	result := make([]interface{}, 0)
	seen := make(map[interface{}]bool)

	for _, v := range s1 {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	for _, v := range s2 {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}

	return result
}

// replaceByIndex 按索引位替换Slice元素
func replaceByIndex(s1, s2 []interface{}, length int) []interface{} {
	result := make([]interface{}, length)

	copy(result, s1)
	copy(result, s2)

	return result
}
