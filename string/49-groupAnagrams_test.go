package main

import "testing"

// func TestGroupAnagrams(t *testing.T) {
// 	var cases = []struct {
// 		input    []string   // input
// 		expected [][]string // expescted result
// 	}{
// 		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
// 			[]string{"ate", "eat", "tea"},
// 			[]string{"nat", "tan"},
// 			[]string{"bat"},
// 		}},
// 	}
// 	for _, tt := range cases {
// 		actual := groupAnagrams(tt.input)
// 		if actual != tt.expected {
// 			t.Errorf("groupAnagrams(%#v) = %v; expected %v", tt.input, actual, tt.expected)
// 		}
// 	}
// }

// func groupAnagrams(strs []string) [][]string {

// }

func BenchmarkEqual(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringSliceEqual(sa, sb)
	}
}

func BenchmarkEqualBCE(b *testing.B) {
	sa := []string{"q", "w", "e", "r", "t"}
	sb := []string{"q", "w", "a", "s", "z", "x"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		StringSliceEqualBCE(sa, sb)
	}
}

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
