package main

import (
	"strings"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	var cases = []struct {
		input    string // input
		expected string // expescted result
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, tt := range cases {
		actual := simplifyPath(tt.input)
		if actual != tt.expected {
			t.Errorf("simplifyPath(%#v) = %v; expected %v", tt.input, actual, tt.expected)
		}
	}
}

func simplifyPath(path string) string {
	ret := []string{}
	for _, v := range strings.Split(path, "/") {
		switch v {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(ret); l > 0 {
				ret = ret[0 : l-1]
			}
			break
		default:
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}
