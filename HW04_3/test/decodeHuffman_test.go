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
			function: decodeHuffman,
			inputs:   []interface{}{"01011101010110000011110001101111", map[string]string{"000":" " , "001":"r",
									 "010":"H", "011":"W", "10":"l", "110":"o", "1110":"e", "1111":"d"}},
			expected: []interface{}{"Hello World"},
		},
		{
			function: decodeHuffman,
			inputs:   []interface{}{"000001010101111011100111010110110000010111000010010101110101010100110011101101100001100111111110100111",
									 map[string]string{"0000":"L", "0001":"u", "0010":"d", "00110":"a",
										"00111":".", "010":"o", "011":"m", "100":" " , "1010":"r",
										"10110":"p", "10111":"l", "1100":"s", "1101":"t",
										"1110":"i", "1111":"e"}},
			expected: []interface{}{"Lorem ipsum dolor sit amet."},
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
