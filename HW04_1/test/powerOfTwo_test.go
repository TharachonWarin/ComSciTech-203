package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	function interface{}
	inputs   []interface{}
	expected []interface{}
}

func TestFunctions(t *testing.T) {
	testCases := []testCase{
		{
			function: powerOfTwo,
			inputs:   []interface{}{uint64(8)},
			expected: []interface{}{true},
		},
		{
			function: powerOfTwo,
			inputs:   []interface{}{uint64(18)},
			expected: []interface{}{false},
		},
		{
			function: powerOfTwo,
			inputs:   []interface{}{uint64(1024)},
			expected: []interface{}{true},
		},
		{
			function: powerOfTwo,
			inputs:   []interface{}{uint64(1025)},
			expected: []interface{}{false},
		},
	}

	runTests(t, testCases)
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
