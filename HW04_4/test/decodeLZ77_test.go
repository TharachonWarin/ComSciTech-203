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
		// {
		// 	function: decodeLZ77,
		// 	inputs: []interface{}{[]Triplet{{0, 0, 97}, {0, 0, 98},
		// 		{2, 2, 99}, {4, 3, 97},
		// 		{2, 2, 97}, {0, 0, 10}}}, //(2,2,99), (4,3,97), (2,2,97), (0,0,10)
		// 	expected: []interface{}{"ababcbababaa\n"},
		// },
		// {
		// 	function: decodeLZ77,
		// 	inputs: []interface{}{[]Triplet{{0, 0, 97}, {0, 0, 98},
		// 		{0, 0, 114}, {3, 1, 99},
		// 		{2, 1, 100}, {7, 4, 10}}}, //(2,2,99), (4,3,97), (2,2,97), (0,0,10)
		// 	expected: []interface{}{"abracadabra\n"},
		// },
		// {
		// 	function: decodeLZ77,
		// 	inputs: []interface{}{[]Triplet{{0, 0, 104}, {0, 0, 97}, {0, 0, 110}, {1, 1, 97}, {5, 1, 32},
		// 		{0, 0, 98}, {7, 2, 97}, {2, 2, 10}}}, //(2,2,99), (4,3,97), (2,2,97), (0,0,10)
		// 	expected: []interface{}{"hannah banana\n"},
		// },
		{
			function: decodeLZ77,
			inputs:   []interface{}{[]Triplet{{0, 0, 99}, {0, 0, 97}, {0, 0, 116}, {0, 0, 10}, {4, 4, 0}}}, //(2,2,99), (4,3,97), (2,2,97), (0,0,10)
			expected: []interface{}{"cat\ncat\n"},
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
