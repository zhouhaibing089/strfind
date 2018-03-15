package main

import "testing"

func match(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for index, value := range left {
		if value != right[index] {
			return false
		}
	}
	return true
}

func TestTableFor(t *testing.T) {
	tests := []struct {
		pattern string
		table   []int
	}{
		{
			pattern: "abcdabd",
			table:   []int{-1, 0, 0, 0, 0, 1, 2},
		},
		{
			pattern: "cabx",
			table:   []int{-1, 0, 0, 0},
		},
	}

	for _, test := range tests {
		table := TableFor(test.pattern)
		if !match(table, test.table) {
			t.Errorf("expect to get %v, but see %v", test.table, table)
		}
	}
}

func TestStrFind(t *testing.T) {
	tests := []struct {
		text    string
		pattern string
		index   int
	}{
		{
			text:    "abcabdcabxcabccdef",
			pattern: "cabx",
			index:   6,
		},
		{
			text:    "ABC ABCDAB ABCDABCDABDE",
			pattern: "ABCDABD",
			index:   15,
		},
		{
			text:    "ABC",
			pattern: "ABCDABD",
			index:   -1,
		},
		{
			text:    "ABCDABC",
			pattern: "ABCDABD",
			index:   -1,
		},
	}

	for _, test := range tests {
		index := StrFind(test.text, test.pattern)
		if index != test.index {
			t.Errorf("expect to get %d, but see %d", test.index, index)
		}
	}
}
