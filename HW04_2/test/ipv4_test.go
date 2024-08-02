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
		// DecodeIP
		{
			function: ipv4Decode,
			inputs:   []interface{}{uint32(3232237069)},
			expected: []interface{}{"192.168.6.13"},
		},
		// {
		// 	function: ipv4Decode,
		// 	inputs:   []interface{}{uint32(1)},
		// 	expected: []interface{}{"0.0.0.1"},
		// },
		// {
		// 	function: ipv4Decode,
		// 	inputs:   []interface{}{"1101.111", uint8(2)},
		// 	expected: []interface{}{"1110.00"},
		// },
		// {
		// 	function: ipv4Decode,
		// 	inputs:   []interface{}{"1101.111", uint8(5)},
		// 	expected: []interface{}{"1101.11100"},
		// },
		// {
		// 	function: ipv4Decode,
		// 	inputs:   []interface{}{"10.00011", uint8(2)},
		// 	expected: []interface{}{"10.00"},
		// },
		// EncodeIP
		{
			function: ipv4Encode,
			inputs:   []interface{}{"10.28.4.2"},
			expected: []interface{}{uint32(169608194)},
		},
		// {
		// 	function: ipv4Encode,
		// 	inputs:   []interface{}{"10.11100", uint8(2)},
		// 	expected: []interface{}{"11.00"},
		// },
		// {
		// 	function: ipv4Encode,
		// 	inputs:   []interface{}{"10.10100", uint8(2)},
		// 	expected: []interface{}{"10.10"},
		// },
		// {
		// 	function: ipv4Encode,
		// 	inputs:   []interface{}{"11.11", uint8(0)},
		// 	expected: []interface{}{"100"},
		// },
		// {
		// 	function: ipv4Encode,
		// 	inputs:   []interface{}{"1001.1", uint8(0)},
		// 	expected: []interface{}{"1010"},
		// },
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
