package flatteners

import "testing"

var getFlattenArrayTests = []struct {
	name string
	inputArray interface{}
	flattenArray []interface{}
}{
	{"simple", []interface{}{1, 2, "32"}, []interface{}{1, 2, "32"}},
	{"one depth", []interface{}{1, 2, "32", []interface{}{324, 234, 4}}, []interface{}{1, 2, "32", 324, 234, 4}},
}

func TestGetFlattenArray(t *testing.T) {
	for _, tt := range getFlattenArrayTests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFlattenArray(tt.inputArray)
			if !equalsArrays(got, tt.flattenArray) {
				t.Errorf("Expected: %v, got: %v", tt.flattenArray, got)
			}
		})
	}
}

func equalsArrays(input, output []interface{}) bool{
	for key, value := range input{
		if value != output[key]{
			return false
		}
	}
	return true
}