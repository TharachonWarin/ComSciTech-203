package test

import (
	"reflect"
	"testing"
)

type testCase struct {
	function interface{}
	inputs   []interface{}
	expected []interface{}
}

func runTests(t *testing.T, cases []testCase) {
	for _, c := range cases {
		inputs := make([]reflect.Value, len(c.inputs))
		for i, input := range c.inputs {
			inputs[i] = reflect.ValueOf(input)
		}

		// Call the function with reflect
		results := reflect.ValueOf(c.function).Call(inputs)
		resultInterfaces := make([]interface{}, len(results))
		for i, result := range results {
			resultInterfaces[i] = result.Interface()
		}

		if !reflect.DeepEqual(resultInterfaces, c.expected) {
			t.Errorf("For function %v with inputs %v, expected %v, got %v",
				reflect.ValueOf(c.function), c.inputs, c.expected, resultInterfaces)
		}
	}
}

func TestFunctions(t *testing.T) {
	testCases := []testCase{
		{
			function: addNSubtract,
			inputs:   []interface{}{"1101", "1", uint8(5)},
			expected: []interface{}{int64(-4), int64(-2)},
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"01101", "01000", uint8(4)},
			expected: []interface{}{int64(5), int64(5)},
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"0111", "011111", uint8(7)},
			expected: []interface{}{int64(38), int64(-24)},  //false
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"0001", "010", uint8(5)},
			expected: []interface{}{int64(3), int64(-1)}, //false
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"0101", "1011", uint8(4)},
			expected: []interface{}{int64(0), int64(-6)},
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"0101", "1", uint8(5)},
			expected: []interface{}{int64(4), int64(6)},
		},
		{
			function: addNSubtract,
			inputs:   []interface{}{"0111", "01", uint8(3)},
			expected: []interface{}{int64(0), int64(-2)},
		},
	}

	runTests(t, testCases)
}
