package anagramms_test

import (
	"reflect"
	"testing"

	"github.com/beganov/anagramms/internal/anagramms"
)

func TestGroupAnagramms(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []string
		want  map[string][]string
	}{
		{
			name:  "empty input",
			input: []string{},
			want:  map[string][]string{},
		},
		{
			name:  "single word",
			input: []string{"hello"},
			want:  map[string][]string{},
		},
		{
			name:  "no anagrams",
			input: []string{"cat", "dog", "bird"},
			want:  map[string][]string{},
		},
		{
			name:  "simple anagrams",
			input: []string{"eat", "tea", "ate"},
			want: map[string][]string{
				"eat": {"ate", "eat", "tea"},
			},
		},
		{
			name:  "multiple groups",
			input: []string{"eat", "tea", "ate", "tan", "nat", "bat"},
			want: map[string][]string{
				"eat": {"ate", "eat", "tea"},
				"tan": {"nat", "tan"},
			},
		},
		{
			name:  "words with different lengths",
			input: []string{"a", "ab", "ba", "abc", "cab", "bca"},
			want: map[string][]string{
				"ab":  {"ab", "ba"},
				"abc": {"abc", "cab", "bca"},
			},
		},
		{
			name:  "case sensitive",
			input: []string{"Eat", "eat", "Tea", "tea"},
			want: map[string][]string{
				"eat": {"eat", "eat", "tea", "tea"},
			},
		},
		{
			name:  "duplicate words",
			input: []string{"eat", "tea", "eat", "ate", "tea"},
			want: map[string][]string{
				"eat": {"ate", "eat", "eat", "tea", "tea"},
			},
		},
		{
			name:  "empty strings",
			input: []string{"", "", "a", ""},
			want: map[string][]string{
				"": {"", "", ""},
			},
		},
		{
			name:  "mixed with single words",
			input: []string{"cat", "act", "tac", "dog", "god", "fish"},
			want: map[string][]string{
				"cat": {"act", "cat", "tac"},
				"dog": {"dog", "god"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := anagramms.GroupAnagramms(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagramms() = %v, want %v", got, tt.want)
			}
		})
	}
}
