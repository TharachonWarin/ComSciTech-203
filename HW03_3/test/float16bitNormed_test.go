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
			function: float16bitNormed,
			inputs:   []interface{}{float32(204203)},
			expected: []interface{}{"0 10010000 1000111"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.002)},
			expected: []interface{}{"0 01110110 0000011"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(-23)},
			expected: []interface{}{"1 10000011 0111000"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.0000000000000000000000000000000001)},
			expected: []interface{}{"0 00001110 0000100"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.00000000000000000000000000000000000005)},
			expected: []interface{}{"0 00000011 0001000"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.00000000000000000000000000000000000002)},
			expected: []interface{}{"0 00000001 1011001"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.0000000000000000000000000000000000000118)},
			expected: []interface{}{"0 00000001 0000000"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(0.0000000000000000000000000000000000000000001)},
			expected: []interface{}{""},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(338953138925153547590470800371487866881)},
			expected: []interface{}{"0 11111110 1111111"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(18446744073709551616)},
			expected: []interface{}{"0 10111111 0000000"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(9223372036854775808)},
			expected: []interface{}{"0 10111110 0000000"},
		},
		{
			function: float16bitNormed,
			inputs:   []interface{}{float32(10.25)},
			expected: []interface{}{"0 10111110 0000000"},
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
