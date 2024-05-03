package count

import "testing"

func TestCount(t *testing.T) {
	tests := map[string]struct {
		input string
		want  map[rune]int
	}{
		"non-empty string": {
			input: "This is a string. It has both uppercase and lowercase characters.",
			want: map[rune]int{
				' ': 10,
				'.': 2,
				'T': 1,
				'h': 4,
				'i': 3,
				's': 7,
				'a': 7,
				't': 4,
				'r': 5,
				'n': 2,
				'g': 1,
				'I': 1,
				'b': 1,
				'o': 2,
				'u': 1,
				'p': 2,
				'c': 4,
				'e': 5,
				'd': 1,
				'l': 1,
				'w': 1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := characters([]rune(test.input))

			if len(got) != len(test.want) {
				t.Fatalf("maps not equal len. got=%d, want=%d", len(got), len(test.want))
			}

			for k, v := range test.want {
				if v != got[k] {
					t.Fatalf("for %c, got=%d, want=%d", k, got[k], v)
				}
			}
		})
	}
}
