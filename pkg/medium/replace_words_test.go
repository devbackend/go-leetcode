package medium_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/medium"
)

func TestReplaceWords(t *testing.T) {
	testCases := map[string]struct {
		dictionary []string
		sentence   string
		expected   string
	}{
		"example 1": {
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			expected:   "the cat was rat by the bat",
		},
		"example 2": {
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			expected:   "a a b c",
		},
		"example 3": {
			dictionary: []string{"catt", "cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			expected:   "the cat was rat by the bat",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := medium.ReplaceWords(tc.dictionary, tc.sentence)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
