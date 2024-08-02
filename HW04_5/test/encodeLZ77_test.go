package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	name     string
	function interface{}
	inputs   []interface{}
	expected []interface{}
}

func TestFunctions(t *testing.T) {
	testCases := []testCase{
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "ababcbababaa\n"},
			expected: []interface{}{[]Triplet{{0, 0, 97}, {0, 0, 98}, {2, 2, 99}, {4, 3, 97}, {2, 2, 97}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{2, 2, "ABCAABC\n"},
			expected: []interface{}{[]Triplet{{0, 0, 65}, {0, 0, 66}, {0, 0, 67}, {0, 0, 65}, {1, 1, 66}, {0, 0, 67}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{3, 3, "ABCAABC\n"},
			expected: []interface{}{[]Triplet{{0, 0, 65}, {0, 0, 66}, {0, 0, 67}, {3, 1, 65}, {0, 0, 66}, {0, 0, 67}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{5, 5, "cat\ncat\n"},
			expected: []interface{}{[]Triplet{{0, 0, 99}, {0, 0, 97}, {0, 0, 116}, {0, 0, 10}, {4, 4, 0}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "cat\ncat\ncat\n"},
			expected: []interface{}{[]Triplet{{0, 0, 99}, {0, 0, 97}, {0, 0, 116}, {0, 0, 10}, {4, 4, 99}, {4, 3, 0}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 3, "cat\ncat\ncat\n"},
			expected: []interface{}{[]Triplet{{0, 0, 99}, {0, 0, 97}, {0, 0, 116}, {0, 0, 10}, {4, 3, 10}, {4, 3, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "tonnam\n"},
			expected: []interface{}{[]Triplet{{0, 0, 116}, {0, 0, 111}, {0, 0, 110}, {1, 1, 97}, {0, 0, 109}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "tonnam\ncat\n"},
			expected: []interface{}{[]Triplet{{0, 0, 116}, {0, 0, 111}, {0, 0, 110}, {1, 1, 97}, {0, 0, 109}, {0, 0, 10}, {0, 0, 99}, {4, 1, 116}, {4, 1, 0}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "abcdabcde\n"},
			expected: []interface{}{[]Triplet{{0, 0, 97}, {0, 0, 98}, {0, 0, 99}, {0, 0, 100}, {4, 4, 101}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "a\nb\n"},
			expected: []interface{}{[]Triplet{{0, 0, 97}, {0, 0, 10}, {0, 0, 98}, {2, 1, 0}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{4, 4, "aaaaaaaaaaaa\n"},
			expected: []interface{}{[]Triplet{{0, 0, 97}, {1, 1, 97}, {3, 3, 97}, {4, 4, 97}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{10, 6, "The cat in the hat sat on the mat.\n"},
			expected: []interface{}{[]Triplet{{0, 0, 84}, {0, 0, 104}, {0, 0, 101}, {0, 0, 32}, {0, 0, 99}, {0, 0, 97}, {0, 0, 116}, {4, 1, 105}, {0, 0, 110}, {3, 1, 116}, {0, 0, 104}, {0, 0, 101}, {4, 1, 104}, {0, 0, 97}, {6, 1, 32}, {0, 0, 115}, {4, 3, 111}, {0, 0, 110}, {3, 1, 116}, {0, 0, 104}, {0, 0, 101}, {4, 1, 109}, {0, 0, 97}, {6, 1, 46}, {0, 0, 10}}},
		},
		{
			name:     "encodeLZ77",
			function: encodeLZ77,
			inputs:   []interface{}{10, 10, "d at the twili\n"},
			expected: []interface{}{[]Triplet{{0, 0, 100}, {0, 0, 32}, {0, 0, 97}, {0, 0, 116}, {3, 1, 116}, {0, 0, 104}, {0, 0, 101}, {4, 2, 119}, {0, 0, 105}, {0, 0, 108}, {2, 1, 10}}},
		},
	}

	runTests(t, testCases)
}

func runTests(t *testing.T, cases []testCase) {
	fmt.Printf("Total case: %d\n------------------------------\n", len(cases))
	var count int
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
			t.Errorf("\nFunction: %v\ninputs:   %v \nexpected: %v \ngot:      %v\n ",
				c.name, c.inputs, c.expected, resultInterfaces)
		} else {
			count += 1
		}
	}
	fmt.Printf("Total PASS case: %d\nTotal FAIL case: %d\n", count, len(cases)-count)
}
