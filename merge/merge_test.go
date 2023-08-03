package merge_test

import (
	"reflect"
	"testing"

	"github.com/aura-studio/boost/merge"
)

func TestMergeMapStructure(t *testing.T) {
	tests := []struct {
		name       string
		m1         map[string]interface{}
		m2         map[string]interface{}
		mergeType  string
		wantResult map[string]interface{}
	}{
		{
			name: "Test Concatenate",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
			mergeType: "concatenate",
			wantResult: map[string]interface{}{
				"a": []interface{}{1, 2, 5, 6},
				"b": []interface{}{3, 4, 7, 8},
			},
		},
		{
			name: "Test RemoveDuplicates",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2, 3},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{2, 3, 5},
				"b": []interface{}{3, 4},
			},
			mergeType: "remove_duplicates",
			wantResult: map[string]interface{}{
				"a": []interface{}{1, 2, 3, 5},
				"b": []interface{}{3, 4},
			},
		},
		{
			name: "Test Overwrite",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
			mergeType: "overwrite",
			wantResult: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
		},
		{
			name: "Test ReplaceByIndexPreferRight",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
			mergeType: "replace_by_index_prefer_right",
			wantResult: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
		},
		{
			name: "Test ReplaceByIndexPreferLeft",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
			mergeType: "replace_by_index_prefer_left",
			wantResult: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
		},
		{
			name: "Test ReplaceByIndexPreferMax",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": []interface{}{3, 4},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
			mergeType: "replace_by_index_prefer_max",
			wantResult: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": []interface{}{7, 8},
			},
		},
		{
			name: "Test ReplaceByIndexPreferRightRec",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": map[string]interface{}{
					"c": []interface{}{3, 4},
				},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": map[string]interface{}{
					"c": []interface{}{7, 8},
				},
			},
			mergeType: "replace_by_index_prefer_right_rec",
			wantResult: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": map[string]interface{}{
					"c": []interface{}{7, 8},
				},
			},
		},
		{
			name: "Test ReplaceByIndexPreferLeftRec",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": map[string]interface{}{
					"c": []interface{}{3, 4},
				},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": map[string]interface{}{
					"c": []interface{}{7, 8},
				},
			},
			mergeType: "replace_by_index_prefer_left_rec",
			wantResult: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": map[string]interface{}{
					"c": []interface{}{3, 4},
				},
			},
		},
		{
			name: "Test ReplaceByIndexPreferMaxRec",
			m1: map[string]interface{}{
				"a": []interface{}{1, 2},
				"b": map[string]interface{}{
					"c": []interface{}{3, 4},
				},
			},
			m2: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": map[string]interface{}{
					"c": []interface{}{7, 8},
				},
			},
			mergeType: "replace_by_index_prefer_max_rec",
			wantResult: map[string]interface{}{
				"a": []interface{}{5, 6},
				"b": map[string]interface{}{
					"c": []interface{}{7, 8},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := merge.MergeMapStructure(tt.m1, tt.m2, tt.mergeType)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MergeMapStructure() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
