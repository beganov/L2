package custom_sort

import (
	"slices"
	"strings"
	"testing"
)

var UnSortedMessage []string = []string{""}
var SortedMessage []string = []string{"Строки не отсортированы"}

func cmpII(a []IndexedInput, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].value != b[i] {
			return false
		}
	}
	return true
}

func TestCustomSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{

		{
			name:  "Empty slice",
			input: []IndexedInput{},
			want:  []string{},
		},
		{
			name:  "Single element",
			input: []IndexedInput{{value: "hello", valueModified: "hello"}},
			want:  []string{"hello"},
		},
		{
			name: "Multiple elements simple",
			input: []IndexedInput{
				{value: "zebra", valueModified: "zebra"},
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"apple", "banana", "zebra"},
		},
		{
			name: "Case sensitivity",
			input: []IndexedInput{
				{value: "Apple", valueModified: "Apple"},
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"Apple", "apple", "banana"},
		},
		{
			name: "Numbers as strings",
			input: []IndexedInput{
				{value: "100", valueModified: "100"},
				{value: "20", valueModified: "20"},
				{value: "3", valueModified: "3"},
			},
			want: []string{"100", "20", "3"},
		},
		{
			name: "Special characters",
			input: []IndexedInput{
				{value: "!test", valueModified: "!test"},
				{value: "#test", valueModified: "#test"},
				{value: "$test", valueModified: "$test"},
				{value: "test", valueModified: "test"},
			},
			want: []string{"!test", "#test", "$test", "test"},
		},
		{
			name: "Emoji and special symbols",
			input: []IndexedInput{
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
			},
			want: []string{"apple", "© copyright", "✓ check", "⭐ star", "🚀 rocket"},
		},

		{
			name: "Very long strings",
			input: []IndexedInput{
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
			},
			want: []string{
				"a" + strings.Repeat("b", 1000) + "c",
				"a" + strings.Repeat("b", 500) + "c",
				"a" + strings.Repeat("b", 100) + "c",
				"abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomSort(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !cmpII(got, tt.want) {
				t.Errorf("CustomSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomReverseSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{
		{
			name:  "Simple reverse",
			input: []IndexedInput{{value: "baba", valueModified: "baba"}, {value: "abba", valueModified: "abba"}},
			want:  []string{"baba", "abba"},
		},
		{
			name:  "Empty slice",
			input: []IndexedInput{},
			want:  []string{},
		},
		{
			name:  "Single element",
			input: []IndexedInput{{value: "hello", valueModified: "hello"}},
			want:  []string{"hello"},
		},
		{
			name: "Multiple elements",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "cherry", valueModified: "cherry"},
				{value: "banana", valueModified: "banana"},
				{value: "date", valueModified: "date"},
			},
			want: []string{"date", "cherry", "banana", "apple"},
		},
		{
			name: "With duplicates",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"banana", "banana", "apple", "apple"},
		},
		{
			name: "Case sensitivity",
			input: []IndexedInput{
				{value: "Apple", valueModified: "Apple"},
				{value: "apple", valueModified: "apple"},
				{value: "Banana", valueModified: "Banana"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"banana", "apple", "Banana", "Apple"},
		},
		{
			name: "Numbers as strings",
			input: []IndexedInput{
				{value: "100", valueModified: "100"},
				{value: "20", valueModified: "20"},
				{value: "3", valueModified: "3"},
			},
			want: []string{"3", "20", "100"},
		},
		{
			name: "Special characters",
			input: []IndexedInput{
				{value: "!test", valueModified: "!test"},
				{value: "#test", valueModified: "#test"},
				{value: "$test", valueModified: "$test"},
				{value: "test", valueModified: "test"},
			},
			want: []string{"test", "$test", "#test", "!test"},
		},
		{
			name: "Emoji and special symbols",
			input: []IndexedInput{
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
			},
			want: []string{"🚀 rocket", "⭐ star", "✓ check", "© copyright", "apple"},
		},
		{
			name: "Very long strings",
			input: []IndexedInput{
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
			},
			want: []string{
				"abc",
				"a" + strings.Repeat("b", 100) + "c",
				"a" + strings.Repeat("b", 500) + "c",
				"a" + strings.Repeat("b", 1000) + "c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomReverseSort(CustomSort(tt.input))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomReverseSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomUnicSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{
		{
			name: "Simple duplicates",
			input: []IndexedInput{
				{value: "baba", valueModified: "baba"},
				{value: "abba", valueModified: "abba"},
				{value: "baba", valueModified: "baba"},
				{value: "abba", valueModified: "abba"},
				{value: "baba", valueModified: "baba"},
				{value: "abba", valueModified: "abba"},
			},
			want: []string{"abba", "baba"},
		},
		{
			name: "All duplicates",
			input: []IndexedInput{
				{value: "test", valueModified: "test"},
				{value: "test", valueModified: "test"},
				{value: "test", valueModified: "test"},
				{value: "test", valueModified: "test"},
			},
			want: []string{"test"},
		},
		{
			name: "No duplicates",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
				{value: "cherry", valueModified: "cherry"},
				{value: "date", valueModified: "date"},
			},
			want: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:  "Empty slice",
			input: []IndexedInput{},
			want:  []string{},
		},
		{
			name: "Single element",
			input: []IndexedInput{
				{value: "hello", valueModified: "hello"},
			},
			want: []string{"hello"},
		},
		{
			name: "Mixed duplicates",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
				{value: "apple", valueModified: "apple"},
				{value: "cherry", valueModified: "cherry"},
				{value: "banana", valueModified: "banana"},
				{value: "date", valueModified: "date"},
				{value: "cherry", valueModified: "cherry"},
			},
			want: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name: "Case-sensitive duplicates",
			input: []IndexedInput{
				{value: "Apple", valueModified: "Apple"},
				{value: "apple", valueModified: "apple"},
				{value: "APPLE", valueModified: "APPLE"},
				{value: "Apple", valueModified: "Apple"},
				{value: "apple", valueModified: "apple"},
			},
			want: []string{"APPLE", "Apple", "apple"},
		},
		{
			name: "Duplicates with special characters",
			input: []IndexedInput{
				{value: "test!", valueModified: "test!"},
				{value: "test", valueModified: "test"},
				{value: "test!", valueModified: "test!"},
				{value: "test", valueModified: "test"},
				{value: "test?", valueModified: "test?"},
			},
			want: []string{"test", "test!", "test?"},
		},
		{
			name: "Duplicates with numbers",
			input: []IndexedInput{
				{value: "file1", valueModified: "file1"},
				{value: "file10", valueModified: "file10"},
				{value: "file1", valueModified: "file1"},
				{value: "file2", valueModified: "file2"},
				{value: "file10", valueModified: "file10"},
				{value: "file2", valueModified: "file2"},
			},
			want: []string{"file1", "file10", "file2"},
		},
		{
			name: "Duplicates with spaces",
			input: []IndexedInput{
				{value: "apple pie", valueModified: "apple pie"},
				{value: "apple", valueModified: "apple"},
				{value: "apple pie", valueModified: "apple pie"},
				{value: "apple cider", valueModified: "apple cider"},
				{value: "apple", valueModified: "apple"},
			},
			want: []string{"apple", "apple cider", "apple pie"},
		},
		{
			name: "Large number of duplicates",
			input: func() []IndexedInput {
				input := make([]IndexedInput, 0, 300)
				words := []IndexedInput{{value: "apple", valueModified: "apple"}, {value: "banana", valueModified: "banana"}, {value: "cherry", valueModified: "cherry"}}
				for i := 0; i < 100; i++ {
					input = append(input, words...)
				}
				return input
			}(),
			want: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Empty strings as duplicates",
			input: []IndexedInput{
				{value: "", valueModified: ""},
				{value: "test", valueModified: "test"},
				{value: "", valueModified: ""},
				{value: "test", valueModified: "test"},
				{value: "", valueModified: ""},
			},
			want: []string{"", "test"},
		},
		{
			name: "Whitespace duplicates",
			input: []IndexedInput{
				{value: "  apple", valueModified: "  apple"},
				{value: " apple", valueModified: " apple"},
				{value: "apple", valueModified: "apple"},
				{value: "  apple", valueModified: "  apple"},
				{value: "apple", valueModified: "apple"},
			},
			want: []string{"  apple", " apple", "apple"},
		},
		{
			name: "Emoji and special symbols duplicates",
			input: []IndexedInput{
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
				{value: "🚀 rocket", valueModified: "🚀 rocket"},
				{value: "⭐ star", valueModified: "⭐ star"},
				{value: "apple", valueModified: "apple"},
				{value: "✓ check", valueModified: "✓ check"},
				{value: "© copyright", valueModified: "© copyright"},
			},
			want: []string{"apple", "© copyright", "✓ check", "⭐ star", "🚀 rocket"},
		},
		{
			name: "Very long strings",
			input: []IndexedInput{
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
				{value: "a" + strings.Repeat("b", 1000) + "c", valueModified: "a" + strings.Repeat("b", 1000) + "c"},
				{value: "a" + strings.Repeat("b", 100) + "c", valueModified: "a" + strings.Repeat("b", 100) + "c"},
				{value: "abc", valueModified: "abc"},
				{value: "a" + strings.Repeat("b", 500) + "c", valueModified: "a" + strings.Repeat("b", 500) + "c"},
			},
			want: []string{
				"a" + strings.Repeat("b", 1000) + "c",
				"a" + strings.Repeat("b", 500) + "c",
				"a" + strings.Repeat("b", 100) + "c",
				"abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomSort(CustomUnicSort(tt.input))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomReverseSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomCollumnSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		col   int
		want  []string
	}{
		{
			name:  "Empty slice",
			input: []IndexedInput{},
			col:   1,
			want:  []string{},
		},
		{
			name: "Single element",
			input: []IndexedInput{
				{value: "hello", valueModified: "hello"},
			},
			col:  1,
			want: []string{"hello"},
		},
		{
			name: "Sort by first column (col=1)",
			input: []IndexedInput{
				{value: "3\tcherry", valueModified: "3\tcherry"},
				{value: "1\tapple", valueModified: "1\tapple"},
				{value: "2\tbanana", valueModified: "2\tbanana"},
			},
			col:  1,
			want: []string{"1\tapple", "2\tbanana", "3\tcherry"},
		},
		{
			name: "Sort by second column (col=2)",
			input: []IndexedInput{
				{value: "1\tzebra", valueModified: "1\tzebra"},
				{value: "2\tapple", valueModified: "2\tapple"},
				{value: "3\tbanana", valueModified: "3\tbanana"},
				{value: "4\tcherry", valueModified: "4\tcherry"},
			},
			col:  2,
			want: []string{"2\tapple", "3\tbanana", "4\tcherry", "1\tzebra"},
		},
		{
			name: "Sort by third column (col=3)",
			input: []IndexedInput{
				{value: "1\tapple\tred", valueModified: "1\tapple\tred"},
				{value: "2\tbanana\tyellow", valueModified: "2\tbanana\tyellow"},
				{value: "3\tcherry\tred", valueModified: "3\tcherry\tred"},
				{value: "4\tapple\tgreen", valueModified: "4\tapple\tgreen"},
			},
			col:  3,
			want: []string{"4\tapple\tgreen", "1\tapple\tred", "3\tcherry\tred", "2\tbanana\tyellow"},
		},
		{
			name: "Missing columns - sort by col=2",
			input: []IndexedInput{
				{value: "single", valueModified: "single"},
				{value: "two\tcolumns", valueModified: "two\tcolumns"},
				{value: "three\tcolumns\there", valueModified: "three\tcolumns\there"},
				{value: "", valueModified: ""},
			},
			col:  2,
			want: []string{"", "single", "three\tcolumns\there", "two\tcolumns"},
		},
		{
			name: "Column beyond available - sort by col=5",
			input: []IndexedInput{
				{value: "a\tb\tc\td\te", valueModified: "a\tb\tc\td\te"},
				{value: "x\ty", valueModified: "x\ty"},
				{value: "single", valueModified: "single"},
				{value: "p\tq\tr", valueModified: "p\tq\tr"},
			},
			col:  5,
			want: []string{"p\tq\tr", "single", "x\ty", "a\tb\tc\td\te"},
		},
		{
			name: "Empty lines and tabs only",
			input: []IndexedInput{
				{value: "", valueModified: ""},
				{value: "\t", valueModified: "\t"},
				{value: "\t\t", valueModified: "\t\t"},
				{value: "valid\tline", valueModified: "valid\tline"},
				{value: "\tonlytab", valueModified: "\tonlytab"},
			},
			col:  1,
			want: []string{"", "\t", "\t\t", "\tonlytab", "valid\tline"},
		},
		{
			name: "Case sensitivity in column sort",
			input: []IndexedInput{
				{value: "1\tApple", valueModified: "1\tApple"},
				{value: "2\tapple", valueModified: "2\tapple"},
				{value: "3\tBanana", valueModified: "3\tBanana"},
				{value: "4\tbanana", valueModified: "4\tbanana"},
			},
			col:  2,
			want: []string{"1\tApple", "3\tBanana", "2\tapple", "4\tbanana"},
		},
		{
			name: "Numbers in column as strings",
			input: []IndexedInput{
				{value: "file\t100", valueModified: "file\t100"},
				{value: "file\t20", valueModified: "file\t20"},
				{value: "file\t3", valueModified: "file\t3"},
				{value: "file\t1000", valueModified: "file\t1000"},
			},
			col:  2,
			want: []string{"file\t100", "file\t1000", "file\t20", "file\t3"},
		},
		{
			name: "Mixed column types",
			input: []IndexedInput{
				{value: "a\t1", valueModified: "a\t1"},
				{value: "b\t10", valueModified: "b\t10"},
				{value: "c\t2", valueModified: "c\t2"},
				{value: "d\ttext", valueModified: "d\ttext"},
				{value: "e\t", valueModified: "e\t"},
			},
			col:  2,
			want: []string{"e\t", "a\t1", "b\t10", "c\t2", "d\ttext"},
		},
		{
			name: "Special characters in columns",
			input: []IndexedInput{
				{value: "1\t!test", valueModified: "1\t!test"},
				{value: "2\t#test", valueModified: "2\t#test"},
				{value: "3\t$test", valueModified: "3\t$test"},
				{value: "4\ttest", valueModified: "4\ttest"},
			},
			col:  2,
			want: []string{"1\t!test", "2\t#test", "3\t$test", "4\ttest"},
		},
		{
			name: "Whitespace in column values",
			input: []IndexedInput{
				{value: "1\t apple", valueModified: "1\t apple"},
				{value: "2\tapple", valueModified: "2\tapple"},
				{value: "3\t  apple", valueModified: "3\t  apple"},
				{value: "4\tbanana", valueModified: "4\tbanana"},
			},
			col:  2,
			want: []string{"3\t  apple", "1\t apple", "2\tapple", "4\tbanana"},
		},
		{
			name: "Tie-breaking with same column values",
			input: []IndexedInput{
				{value: "2\tapple\tred", valueModified: "2\tapple\tred"},
				{value: "1\tapple\tgreen", valueModified: "1\tapple\tgreen"},
				{value: "3\tbanana\tyellow", valueModified: "3\tbanana\tyellow"},
				{value: "4\tapple\tred", valueModified: "4\tapple\tred"},
			},
			col:  2,
			want: []string{"1\tapple\tgreen", "2\tapple\tred", "4\tapple\tred", "3\tbanana\tyellow"},
		},
		{
			name: "Multiple spaces as delimiters (should still use tab)",
			input: []IndexedInput{
				{value: "1  apple  red", valueModified: "1  apple  red"},
				{value: "2\tbanana\tyellow", valueModified: "2\tbanana\tyellow"},
				{value: "3  cherry  red", valueModified: "3  cherry  red"},
			},
			col:  2,
			want: []string{"1  apple  red", "3  cherry  red", "2\tbanana\tyellow"},
		},
		{
			name: "Column sort with negative column number",
			input: []IndexedInput{
				{value: "1\tapple", valueModified: "1\tapple"},
				{value: "2\tbanana", valueModified: "2\tbanana"},
				{value: "3\tcherry", valueModified: "3\tcherry"},
			},
			col:  -1,
			want: []string{"1\tapple", "2\tbanana", "3\tcherry"},
		},
		{
			name: "Column sort with zero column number",
			input: []IndexedInput{
				{value: "1\tapple", valueModified: "1\tapple"},
				{value: "2\tbanana", valueModified: "2\tbanana"},
				{value: "3\tcherry", valueModified: "3\tcherry"},
			},
			col:  0,
			want: []string{"1\tapple", "2\tbanana", "3\tcherry"},
		},
		{
			name: "Large column number",
			input: []IndexedInput{
				{value: "a", valueModified: "a"},
				{value: "b", valueModified: "b"},
				{value: "c", valueModified: "c"},
			},
			col:  100,
			want: []string{"a", "b", "c"},
		},
		{
			name: "Unicode characters in columns",
			input: []IndexedInput{
				{value: "1\tcafé", valueModified: "1\tcafé"},
				{value: "2\tcafe", valueModified: "2\tcafe"},
				{value: "3\trésumé", valueModified: "3\trésumé"},
				{value: "4\tresume", valueModified: "4\tresume"},
			},
			col:  2,
			want: []string{"2\tcafe", "1\tcafé", "4\tresume", "3\trésumé"},
		},
		{
			name: "Emoji in columns",
			input: []IndexedInput{
				{value: "1\t🚀 rocket", valueModified: "1\t🚀 rocket"},
				{value: "2\t⭐ star", valueModified: "2\t⭐ star"},
				{value: "3\tapple", valueModified: "3\tapple"},
				{value: "4\t✓ check", valueModified: "4\t✓ check"},
			},
			col:  2,
			want: []string{"3\tapple", "4\t✓ check", "2\t⭐ star", "1\t🚀 rocket"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomSort(CustomCollumnSort(CustomSort(tt.input), tt.col))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomCollumnSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomNumericSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{{
		name:  "Empty input",
		input: []IndexedInput{},
		want:  []string{},
	},
		{
			name: "Single element",
			input: []IndexedInput{
				{value: "42", valueModified: "42"},
			},
			want: []string{"42"},
		},
		{
			name: "All same element",
			input: []IndexedInput{
				{value: "42", valueModified: "42"},
				{value: "42", valueModified: "42"},
				{value: "42", valueModified: "42"},
			},
			want: []string{"42", "42", "42"},
		},
		{
			name: "Simple mixed numbers and text",
			input: []IndexedInput{
				{value: "12", valueModified: "12"},
				{value: "4abba", valueModified: "4abba"},
				{value: "4", valueModified: "4"},
				{value: "0", valueModified: "0"},
				{value: "3babba", valueModified: "3babba"},
			},
			want: []string{"3babba", "4abba", "0", "4", "12"},
		},
		{
			name: "Pure numbers ascending",
			input: []IndexedInput{
				{value: "100", valueModified: "100"},
				{value: "2", valueModified: "2"},
				{value: "50", valueModified: "50"},
				{value: "1", valueModified: "1"},
				{value: "10", valueModified: "10"},
			},
			want: []string{"1", "2", "10", "50", "100"},
		},
		{
			name: "Negative numbers",
			input: []IndexedInput{
				{value: "-5", valueModified: "-5"},
				{value: "10", valueModified: "10"},
				{value: "-20", valueModified: "-20"},
				{value: "0", valueModified: "0"},
				{value: "5", valueModified: "5"},
			},
			want: []string{"-20", "-5", "0", "5", "10"},
		},
		{
			name: "Decimal numbers",
			input: []IndexedInput{
				{value: "3.14", valueModified: "3.14"},
				{value: "1.5", valueModified: "1.5"},
				{value: "2.0", valueModified: "2.0"},
				{value: "10.1", valueModified: "10.1"},
				{value: "2.5", valueModified: "2.5"},
			},
			want: []string{"1.5", "2.0", "2.5", "3.14", "10.1"},
		},
		{
			name: "Mixed alphanumeric with no numeric prefix",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
				{value: "123", valueModified: "123"},
				{value: "cherry", valueModified: "cherry"},
				{value: "45", valueModified: "45"},
			},
			want: []string{"apple", "banana", "cherry", "45", "123"},
		},
		{
			name: "Mixed positive and negative decimals",
			input: []IndexedInput{
				{value: "-3.14", valueModified: "-3.14"},
				{value: "1.5", valueModified: "1.5"},
				{value: "-2.0", valueModified: "-2.0"},
				{value: "0.0", valueModified: "0.0"},
				{value: "-0.5", valueModified: "-0.5"},
			},
			want: []string{"-3.14", "-2.0", "-0.5", "0.0", "1.5"},
		},
		{
			name: "Empty strings and whitespace",
			input: []IndexedInput{
				{value: "", valueModified: ""},
				{value: " 100", valueModified: " 100"},
				{value: "100", valueModified: "100"},
				{value: "\t50", valueModified: "\t50"},
				{value: "0", valueModified: "0"},
			},
			want: []string{"", "\t50", " 100", "0", "100"},
		},
		{
			name: "Very large numbers",
			input: []IndexedInput{
				{value: "9999999999999999", valueModified: "9999999999999999"},
				{value: "10000000000000000", valueModified: "10000000000000000"},
				{value: "1", valueModified: "1"},
				{value: "5000000000000000", valueModified: "5000000000000000"},
			},
			want: []string{"1", "5000000000000000", "9999999999999999", "10000000000000000"},
		},
		{
			name: "Numbers with signs",
			input: []IndexedInput{
				{value: "+100", valueModified: "+100"},
				{value: "-50", valueModified: "-50"},
				{value: "100", valueModified: "100"},
				{value: "+10", valueModified: "+10"},
				{value: "-10", valueModified: "-10"},
			},
			want: []string{"-50", "-10", "+10", "+100", "100"},
		},
		{
			name: "Very long numbers",
			input: []IndexedInput{
				{value: "1" + strings.Repeat("0", 1000000), valueModified: "1" + strings.Repeat("0", 1000000)},
				{value: "1" + strings.Repeat("0", 100), valueModified: "1" + strings.Repeat("0", 100)},
				{value: "1" + strings.Repeat("0", 10000), valueModified: "1" + strings.Repeat("0", 10000)},
			},
			want: []string{
				"1" + strings.Repeat("0", 100),
				"1" + strings.Repeat("0", 10000),
				"1" + strings.Repeat("0", 1000000),
			},
		},
		{
			name: "Numbers starting with plus",
			input: []IndexedInput{
				{value: "+100", valueModified: "+100"},
				{value: "100", valueModified: "100"},
				{value: "+10", valueModified: "+10"},
				{value: "10", valueModified: "10"},
			},
			want: []string{"+10", "10", "+100", "100"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomNumericSort(CustomSort(tt.input))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomNumericSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomMonthSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{
		{
			name: "Simple Test",
			input: []IndexedInput{
				{value: "September", valueModified: "September"},
				{value: "January", valueModified: "January"},
				{value: "4", valueModified: "4"},
				{value: "0", valueModified: "0"},
				{value: "May", valueModified: "May"},
			},
			want: []string{"0", "4", "January", "May", "September"},
		},
		{
			name: "Full month names in random order",
			input: []IndexedInput{
				{value: "December", valueModified: "December"},
				{value: "March", valueModified: "March"},
				{value: "August", valueModified: "August"},
				{value: "January", valueModified: "January"},
				{value: "July", valueModified: "July"},
				{value: "April", valueModified: "April"},
				{value: "February", valueModified: "February"},
				{value: "November", valueModified: "November"},
				{value: "October", valueModified: "October"},
				{value: "September", valueModified: "September"},
				{value: "June", valueModified: "June"},
				{value: "May", valueModified: "May"},
			},
			want: []string{
				"January", "February", "March", "April", "May", "June",
				"July", "August", "September", "October", "November", "December",
			},
		},
		{
			name: "Month abbreviations",
			input: []IndexedInput{
				{value: "Dec", valueModified: "Dec"},
				{value: "Mar", valueModified: "Mar"},
				{value: "Aug", valueModified: "Aug"},
				{value: "Jan", valueModified: "Jan"},
				{value: "Jul", valueModified: "Jul"},
				{value: "Apr", valueModified: "Apr"},
				{value: "Feb", valueModified: "Feb"},
				{value: "Nov", valueModified: "Nov"},
				{value: "Oct", valueModified: "Oct"},
				{value: "Sep", valueModified: "Sep"},
				{value: "Jun", valueModified: "Jun"},
				{value: "May", valueModified: "May"},
			},
			want: []string{
				"Jan", "Feb", "Mar", "Apr", "May", "Jun",
				"Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
			},
		},
		{
			name: "Mixed full and abbreviated month names",
			input: []IndexedInput{
				{value: "December", valueModified: "December"},
				{value: "Dec", valueModified: "Dec"},
				{value: "March", valueModified: "March"},
				{value: "Mar", valueModified: "Mar"},
				{value: "January", valueModified: "January"},
				{value: "Jan", valueModified: "Jan"},
			},
			want: []string{"Jan", "January", "Mar", "March", "Dec", "December"},
		},
		{
			name: "Months with numbers and text",
			input: []IndexedInput{
				{value: "March", valueModified: "March"},
				{value: "100", valueModified: "100"},
				{value: "apple", valueModified: "apple"},
				{value: "January", valueModified: "January"},
				{value: "50", valueModified: "50"},
				{value: "December", valueModified: "December"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"100", "50", "apple", "banana", "January", "March", "December"},
		},
		{
			name: "Case insensitive month names",
			input: []IndexedInput{
				{value: "MARCH", valueModified: "MARCH"},
				{value: "march", valueModified: "march"},
				{value: "March", valueModified: "March"},
				{value: "JANUARY", valueModified: "JANUARY"},
				{value: "january", valueModified: "january"},
				{value: "January", valueModified: "January"},
			},
			want: []string{"JANUARY", "January", "january", "MARCH", "March", "march"},
		},
		{
			name: "Invalid month-like strings",
			input: []IndexedInput{
				{value: "Januar", valueModified: "Januar"},
				{value: "Febru", valueModified: "Febru"},
				{value: "March", valueModified: "March"},
				{value: "Jan", valueModified: "Jan"},
				{value: "Feb", valueModified: "Feb"},
				{value: "Marsh", valueModified: "Marsh"},
			},
			want: []string{"Jan", "Januar", "Feb", "Febru", "March", "Marsh"},
		},
		{
			name: "Empty strings and months",
			input: []IndexedInput{
				{value: "", valueModified: ""},
				{value: "December", valueModified: "December"},
				{value: "   ", valueModified: "   "},
				{value: "January", valueModified: "January"},
				{value: "\t", valueModified: "\t"},
			},
			want: []string{"", "\t", "   ", "January", "December"},
		},
		{
			name: "Duplicate month names",
			input: []IndexedInput{
				{value: "March", valueModified: "March"},
				{value: "January", valueModified: "January"},
				{value: "March", valueModified: "March"},
				{value: "January", valueModified: "January"},
				{value: "December", valueModified: "December"},
				{value: "December", valueModified: "December"},
			},
			want: []string{"January", "January", "March", "March", "December", "December"},
		},
		{
			name: "Month names with numbers",
			input: []IndexedInput{
				{value: "Jan2023", valueModified: "Jan2023"},
				{value: "Feb2023", valueModified: "Feb2023"},
				{value: "Mar2023", valueModified: "Mar2023"},
				{value: "January2023", valueModified: "January2023"},
				{value: "February2023", valueModified: "February2023"},
				{value: "March2023", valueModified: "March2023"},
			},
			want: []string{"Jan2023", "January2023", "Feb2023", "February2023", "Mar2023", "March2023"},
		},
		{
			name: "Single month",
			input: []IndexedInput{
				{value: "July", valueModified: "July"},
			},
			want: []string{"July"},
		},
		{
			name:  "Empty input",
			input: []IndexedInput{},
			want:  []string{},
		},
		{
			name: "No month names",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "banana", valueModified: "banana"},
				{value: "123", valueModified: "123"},
				{value: "cherry", valueModified: "cherry"},
			},
			want: []string{"123", "apple", "banana", "cherry"},
		},
		{
			name: "Month names with special characters",
			input: []IndexedInput{
				{value: "January!", valueModified: "January!"},
				{value: "February?", valueModified: "February?"},
				{value: "March.", valueModified: "March."},
				{value: "April", valueModified: "April"},
			},
			want: []string{"January!", "February?", "March.", "April"},
		},
		{
			name: "Mixed case month abbreviations",
			input: []IndexedInput{
				{value: "JAN", valueModified: "JAN"},
				{value: "Feb", valueModified: "Feb"},
				{value: "MAR", valueModified: "MAR"},
				{value: "apr", valueModified: "apr"},
				{value: "MAY", valueModified: "MAY"},
				{value: "jun", valueModified: "jun"},
			},
			want: []string{"JAN", "Feb", "MAR", "apr", "MAY", "jun"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomSort(CustomMonthSort(CustomSort(tt.input)))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomMonthSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomHumanSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{
		{
			name: "Simple Test",
			input: []IndexedInput{
				{value: "September", valueModified: "September"},
				{value: "January", valueModified: "January"},
				{value: "1M", valueModified: "1M"},
				{value: "40000K", valueModified: "40000K"},
				{value: "0.0002K", valueModified: "0.0002K"},
				{value: "May", valueModified: "May"},
			},
			want: []string{"January", "May", "September", "0.0002K", "1M", "40000K"},
		},
		{
			name: "Basic human-readable sizes",
			input: []IndexedInput{
				{value: "1K", valueModified: "1K"},
				{value: "100", valueModified: "100"},
				{value: "2M", valueModified: "2M"},
				{value: "500K", valueModified: "500K"},
				{value: "1G", valueModified: "1G"},
			},
			want: []string{"100", "1K", "500K", "2M", "1G"},
		},
		{
			name: "Mixed units K, M, G, T",
			input: []IndexedInput{
				{value: "1T", valueModified: "1T"},
				{value: "500G", valueModified: "500G"},
				{value: "2M", valueModified: "2M"},
				{value: "100K", valueModified: "100K"},
				{value: "1", valueModified: "1"},
			},
			want: []string{"1", "100K", "2M", "500G", "1T"},
		},
		{
			name: "Decimal human sizes",
			input: []IndexedInput{
				{value: "1.5K", valueModified: "1.5K"},
				{value: "1.2M", valueModified: "1.2M"},
				{value: "1.0K", valueModified: "1.0K"},
				{value: "500", valueModified: "500"},
				{value: "2.5G", valueModified: "2.5G"},
			},
			want: []string{"500", "1.0K", "1.5K", "1.2M", "2.5G"},
		},
		{
			name: "Mixed text and human sizes",
			input: []IndexedInput{
				{value: "apple", valueModified: "apple"},
				{value: "1M", valueModified: "1M"},
				{value: "banana", valueModified: "banana"},
				{value: "500K", valueModified: "500K"},
				{value: "cherry", valueModified: "cherry"},
				{value: "1", valueModified: "1"},
			},
			want: []string{"apple", "banana", "cherry", "1", "500K", "1M"},
		},
		{
			name: "No suffix numbers",
			input: []IndexedInput{
				{value: "3000", valueModified: "3000"},
				{value: "1K", valueModified: "1K"},
				{value: "1024", valueModified: "1024"},
				{value: "2K", valueModified: "2K"},
			},
			want: []string{"1K", "1024", "2K", "3000"},
		},
		{
			name: "Very large human sizes",
			input: []IndexedInput{
				{value: "0.6K", valueModified: "0.6K"},
				{value: "0.001M", valueModified: "0.001M"},
				{value: "0.000002G", valueModified: "0.000002G"},
				{value: "1", valueModified: "1"},
				{value: "500", valueModified: "500"},
			},
			want: []string{"1", "500", "0.6K", "0.001M", "0.000002G"},
		},
		{
			name: "Multiple same unit different values",
			input: []IndexedInput{
				{value: "10K", valueModified: "10K"},
				{value: "1K", valueModified: "1K"},
				{value: "100K", valueModified: "100K"},
				{value: "5K", valueModified: "5K"},
				{value: "50K", valueModified: "50K"},
			},
			want: []string{"1K", "5K", "10K", "50K", "100K"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomNumericSort(CustomHumanSort(CustomSort(tt.input)))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomHumanSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomTrailingBlanksSort(t *testing.T) {
	tests := []struct {
		name  string
		input []IndexedInput
		want  []string
	}{

		{
			name: "Trailing spaces simple",
			input: []IndexedInput{
				{value: "apple ", valueModified: "apple "},
				{value: "apple", valueModified: "apple"},
				{value: "banana  ", valueModified: "banana  "},
				{value: "banana", valueModified: "banana"},
				{value: "cherry", valueModified: "cherry"},
			},
			want: []string{"apple", "apple ", "banana", "banana  ", "cherry"},
		},
		{
			name: "Trailing tabs",
			input: []IndexedInput{
				{value: "apple\t", valueModified: "apple\t"},
				{value: "apple", valueModified: "apple"},
				{value: "banana\t\t", valueModified: "banana\t\t"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"apple", "apple\t", "banana", "banana\t\t"},
		},
		{
			name: "Mixed trailing spaces and tabs",
			input: []IndexedInput{
				{value: "apple ", valueModified: "apple "},
				{value: "apple\t", valueModified: "apple\t"},
				{value: "apple", valueModified: "apple"},
				{value: "banana  ", valueModified: "banana  "},
				{value: "banana\t", valueModified: "banana\t"},
				{value: "banana", valueModified: "banana"},
			},
			want: []string{"apple", "apple ", "apple\t", "banana", "banana  ", "banana\t"},
		},
		{
			name: "Empty strings and all blanks",
			input: []IndexedInput{
				{value: "", valueModified: ""},
				{value: " ", valueModified: " "},
				{value: "\t", valueModified: "\t"},
				{value: "  ", valueModified: "  "},
				{value: "\t\t", valueModified: "\t\t"},
			},
			want: []string{"", " ", "  ", "\t", "\t\t"},
		},
		{
			name: "Only trailing blanks",
			input: []IndexedInput{
				{value: "test ", valueModified: "test "},
				{value: "test\t", valueModified: "test\t"},
				{value: "test  ", valueModified: "test  "},
				{value: "test\t\t", valueModified: "test\t\t"},
			},
			want: []string{"test ", "test  ", "test\t", "test\t\t"},
		},
		{
			name: "Trailing blanks with multiple words",
			input: []IndexedInput{
				{value: "apple pie", valueModified: "apple pie"},
				{value: "apple pie ", valueModified: "apple pie "},
				{value: "apple pie\t", valueModified: "banana split\t"},
			},
			want: []string{"apple pie", "apple pie ", "apple pie\t"},
		},
		{
			name:  "Empty slice",
			input: []IndexedInput{},
			want:  []string{},
		},
		{
			name: "Single element",
			input: []IndexedInput{
				{value: "test ", valueModified: "test "},
			},
			want: []string{"test "},
		},
		{
			name: "Very long trailing blanks",
			input: []IndexedInput{
				{value: "text" + strings.Repeat(" ", 10), valueModified: "text" + strings.Repeat(" ", 10)},
				{value: "text" + strings.Repeat("\t", 5), valueModified: "text" + strings.Repeat("\t", 5)},
				{value: "text", valueModified: "text"},
				{value: "text" + strings.Repeat(" ", 5), valueModified: "text" + strings.Repeat(" ", 5)},
				{value: "text" + strings.Repeat("\t", 10), valueModified: "text" + strings.Repeat("\t", 10)},
			},
			want: []string{
				"text",
				"text" + strings.Repeat(" ", 5),
				"text" + strings.Repeat(" ", 10),
				"text" + strings.Repeat("\t", 5),
				"text" + strings.Repeat("\t", 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomSort(CustomTrailingBlanksSort(CustomSort(tt.input)))
			if !cmpII(got, tt.want) {
				t.Errorf("CustomTrailingBlanksSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n     bool
		r     bool
		u     bool
		b     bool
		m     bool
		c     bool
		h     bool
		k     int
		input []string
		want  []string
	}{

		{
			name:  "Simple sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  []string{"apple", "banana", "zebra"},
		},
		{
			name:  "Reverse sort",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  []string{"zebra", "banana", "apple"},
		},
		{
			name:  "Unique sort",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "Numeric sort",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3"},
			want:  []string{"3", "20", "100"},
		},
		{
			name:  "Column sort (k=2)",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"1\tzebra", "2\tapple", "3\tbanana"},
			want:  []string{"2\tapple", "3\tbanana", "1\tzebra"},
		},
		{
			name:  "Human sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1K", "100", "2M", "500K"},
			want:  []string{"100", "1K", "500K", "2M"},
		},
		{
			name:  "Month sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     true,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March"},
			want:  []string{"January", "March", "December"},
		},
		{
			name:  "Trailing blanks sort",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "banana  "},
			want:  []string{"apple", "apple ", "banana  "},
		},
		{
			name:  "Combination: unique + reverse",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  []string{"cherry", "banana", "apple"},
		},
		{
			name:  "Combination: numeric + human",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K"},
			want:  []string{"500", "2K", "100K", "1M"},
		},
		{
			name:  "All false - default sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  []string{"apple", "banana", "zebra"},
		},
		{
			name:  "Reverse sort",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  []string{"zebra", "banana", "apple"},
		},
		{
			name:  "Unique sort",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  []string{"apple", "banana", "cherry"},
		},
		{
			name:  "Numeric sort",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3"},
			want:  []string{"3", "20", "100"},
		},
		{
			name:  "Column sort (k=2)",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"1\tzebra", "2\tapple", "3\tbanana"},
			want:  []string{"2\tapple", "3\tbanana", "1\tzebra"},
		},
		{
			name:  "Human sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1K", "100", "2M", "500K"},
			want:  []string{"100", "1K", "500K", "2M"},
		},
		{
			name:  "Month sort",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     true,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March"},
			want:  []string{"January", "March", "December"},
		},
		{
			name:  "Trailing blanks sort",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "banana  "},
			want:  []string{"apple", "apple ", "banana  "},
		},
		{
			name:  "Combination: unique + reverse",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  []string{"cherry", "banana", "apple"},
		},
		{
			name:  "Combination: numeric + human",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K"},
			want:  []string{"500", "2K", "100K", "1M"},
		},
		{
			name:  "Combination: month + reverse",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     true,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March"},
			want:  []string{"December", "March", "January"},
		},
		{
			name:  "Combination: numeric + reverse",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3"},
			want:  []string{"100", "20", "3"},
		},
		{
			name:  "Combination: unique + blanks",
			n:     false,
			r:     false,
			u:     true,
			b:     true,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "apple ", "banana", "banana  ", "banana", "banana  "},
			want:  []string{"apple", "apple ", "banana", "banana  "},
		},
		{
			name:  "Combination: column + numeric (k=2)",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"a\t100", "b\t20", "c\t3"},
			want:  []string{"c\t3", "b\t20", "a\t100"},
		},
		{
			name:  "Combination: column + reverse (k=2)",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"1\tapple", "2\tbanana", "3\tcherry"},
			want:  []string{"3\tcherry", "2\tbanana", "1\tapple"},
		},
		{
			name:  "Combination: human + reverse",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1K", "100", "2M", "500K"},
			want:  []string{"2M", "500K", "1K", "100"},
		},
		{
			name:  "Combination: month + unique",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     true,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March", "December", "January"},
			want:  []string{"January", "March", "December"},
		},
		{
			name:  "Combination: numeric + unique",
			n:     true,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3", "100", "20"},
			want:  []string{"3", "20", "100"},
		},
		{
			name:  "Combination: blanks + reverse",
			n:     false,
			r:     true,
			u:     false,
			b:     true,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "banana  "},
			want:  []string{"banana  ", "apple ", "apple"},
		},
		{
			name:  "Combination: column + human (k=2)",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     2,
			input: []string{"a\t1K", "b\t100", "c\t2M", "d\t500K"},
			want:  []string{"b\t100", "a\t1K", "d\t500K", "c\t2M"},
		},
		{
			name:  "Combination: month + blanks",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     true,
			c:     false,
			h:     false,
			k:     1,
			input: []string{"December ", "January", "March  "},
			want:  []string{"January", "March  ", "December "},
		},
		{
			name:  "Combination: column + unique + reverse (k=2)",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"1\tapple", "2\tbanana", "1\tapple", "4\tcherry", "2\tbanana"},
			want:  []string{"4\tcherry", "2\tbanana", "1\tapple"},
		},
		{
			name:  "Complex: numeric + human + reverse",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K", "3G"},
			want:  []string{"3G", "1M", "100K", "2K", "500"},
		},
		{
			name:  "Complex: blanks + column + numeric (k=2)",
			n:     true,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     false,
			h:     false,
			k:     2,
			input: []string{"x\t100 ", "y\t20", "z\t3 "},
			want:  []string{"z\t3 ", "y\t20", "x\t100 "},
		},
		{
			name:  "All flags false with empty input",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     false,
			h:     false,
			k:     1,
			input: []string{},
			want:  []string{},
		},
		{
			name:  "Single element with all flags",
			n:     true,
			r:     true,
			u:     true,
			b:     true,
			m:     true,
			c:     false,
			h:     true,
			k:     1,
			input: []string{"test"},
			want:  []string{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.n, tt.r, tt.u, tt.b, tt.m, tt.c, tt.h, tt.k, tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !slices.Equal(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n     bool
		r     bool
		u     bool
		b     bool
		m     bool
		c     bool
		h     bool
		k     int
		input []string
		want  []string
	}{

		{
			name:  "Simple sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Reverse sort Unsorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Unique sort Unsorted",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Numeric sort Unsorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3"},
			want:  UnSortedMessage,
		},
		{
			name:  "Column sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"1\tzebra", "2\tapple", "3\tbanana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Human sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"1K", "100", "2M", "500K"},
			want:  UnSortedMessage,
		},
		{
			name:  "Month sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March"},
			want:  UnSortedMessage,
		},
		{
			name:  "Trailing blanks sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "banana  "},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: unique + reverse Unsorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: numeric + human Unsorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K"},
			want:  UnSortedMessage,
		},
		{
			name:  "All false - default sort Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"zebra", "apple", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: unique + reverse Unsorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "apple", "cherry", "banana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: numeric + human Unsorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: month + reverse Unsorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: numeric + reverse Unsorted",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"100", "3", "20"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: unique + blanks Unsorted",
			n:     false,
			r:     false,
			u:     true,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "apple ", "banana", "banana  ", "banana", "banana  "},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: column + numeric Unsorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"a\t100", "b\t20", "c\t3"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: column + reverse Unsorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"1\tapple", "2\tbanana", "3\tcherry"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: human + reverse Unsorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"1K", "100", "2M", "500K"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: month + unique Unsorted",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"December", "January", "March", "December", "January"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: numeric + unique Unsorted",
			n:     true,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3", "100", "20"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: blanks + reverse Unsorted",
			n:     false,
			r:     true,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple ", "apple", "banana  "},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: column + human Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     2,
			input: []string{"a\t1K", "b\t100", "c\t2M", "d\t500K"},
			want:  UnSortedMessage,
		},
		{
			name:  "Combination: month + blanks Unsorted",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"December ", "January", "March  "},
			want:  UnSortedMessage,
		},

		{
			name:  "Combination: column + unique + reverse Unsorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"4\tcherry", "1\tapple", "1\tapple", "2\tbanana"},
			want:  UnSortedMessage,
		},
		{
			name:  "Complex: numeric + human + reverse Unsorted",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"1M", "100K", "500", "2K", "3G"},
			want:  UnSortedMessage,
		},
		{
			name:  "Complex: blanks + column + numeric Unsorted",
			n:     true,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"x\t100 ", "y\t20", "z\t3 "},
			want:  UnSortedMessage,
		},
		{
			name:  "All flags false with empty input",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{},
			want:  SortedMessage,
		},
		{
			name:  "Single element with all flags",
			n:     true,
			r:     true,
			u:     true,
			b:     true,
			m:     true,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"test"},
			want:  SortedMessage,
		},

		{
			name:  "Simple sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "zebra"},
			want:  SortedMessage,
		},
		{
			name:  "Reverse sort Sorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"zebra", "banana", "apple"},
			want:  SortedMessage,
		},
		{
			name:  "Unique sort Sorted",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "cherry"},
			want:  SortedMessage,
		},
		{
			name:  "Numeric sort Sorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"3", "20", "100"},
			want:  SortedMessage,
		},
		{
			name:  "Column sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"2\tapple", "3\tbanana", "1\tzebra"},
			want:  SortedMessage,
		},
		{
			name:  "Human sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"100", "1K", "500K", "2M"},
			want:  SortedMessage,
		},
		{
			name:  "Month sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"January", "March", "December"},
			want:  SortedMessage,
		},
		{
			name:  "Trailing blanks sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "apple ", "banana  "},
			want:  SortedMessage,
		},
		{
			name:  "Combination: unique + reverse Sorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"cherry", "banana", "apple"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: numeric + human Sorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"500", "2K", "100K", "1M"},
			want:  SortedMessage,
		},
		{
			name:  "All false - default sort Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "banana", "zebra"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: unique + reverse Sorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"cherry", "banana", "apple"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: month + reverse Sorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"December", "March", "January"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: numeric + reverse Sorted",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"100", "20", "3"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: unique + blanks Sorted",
			n:     false,
			r:     false,
			u:     true,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"apple", "apple ", "banana", "banana  "},
			want:  SortedMessage,
		},
		{
			name:  "Combination: column + numeric Sorted",
			n:     true,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"c\t3", "b\t20", "a\t100"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: column + reverse Sorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"3\tcherry", "2\tbanana", "1\tapple"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: human + reverse Sorted",
			n:     false,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"2M", "500K", "1K", "100"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: month + unique Sorted",
			n:     false,
			r:     false,
			u:     true,
			b:     false,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"January", "March", "December"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: numeric + unique Sorted",
			n:     true,
			r:     false,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"3", "20", "100"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: blanks + reverse Sorted",
			n:     false,
			r:     true,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"banana  ", "apple ", "apple"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: column + human Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     2,
			input: []string{"b\t100", "a\t1K", "d\t500K", "c\t2M"},
			want:  SortedMessage,
		},
		{
			name:  "Combination: month + blanks Sorted",
			n:     false,
			r:     false,
			u:     false,
			b:     true,
			m:     true,
			c:     true,
			h:     false,
			k:     1,
			input: []string{"January", "March  ", "December "},
			want:  SortedMessage,
		},
		{
			name:  "Combination: column + unique + reverse Sorted",
			n:     false,
			r:     true,
			u:     true,
			b:     false,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"4\tcherry", "2\tbanana", "1\tapple"},
			want:  SortedMessage,
		},
		{
			name:  "Complex: numeric + human + reverse Sorted",
			n:     true,
			r:     true,
			u:     false,
			b:     false,
			m:     false,
			c:     true,
			h:     true,
			k:     1,
			input: []string{"3G", "1M", "100K", "2K", "500"},
			want:  SortedMessage,
		},
		{
			name:  "Complex: blanks + column + numeric Sorted",
			n:     true,
			r:     false,
			u:     false,
			b:     true,
			m:     false,
			c:     true,
			h:     false,
			k:     2,
			input: []string{"z\t3 ", "y\t20", "x\t100 "},
			want:  SortedMessage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.n, tt.r, tt.u, tt.b, tt.m, tt.c, tt.h, tt.k, tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !slices.Equal(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
