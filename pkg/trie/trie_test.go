package trie_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/devbackend/go-leetcode/pkg/trie"
)

func TestTrie_Search(t *testing.T) {
	tr := trie.Constructor()

	cases := []struct {
		insert         string
		search         string
		prefix         string
		rootCandidate  string
		expectedSearch bool
		expectedPrefix bool
		expectedRoot   string
	}{
		{
			search: "foobar",
			prefix: "foo",
		},
		{
			insert: "apple",
			search: "banana",
			prefix: "ban",
		},
		{
			search: "app",
		},
		{
			prefix:         "apple",
			expectedPrefix: true,
		},
		{
			search: "ast",
		},
		{
			search:         "apple",
			expectedSearch: true,
			prefix:         "app",
			expectedPrefix: true,
		},
		{
			insert:         "app",
			search:         "app",
			expectedSearch: true,
		},
		{
			rootCandidate: "castle",
		},
		{
			insert:        "cat",
			rootCandidate: "cattle",
			expectedRoot:  "cat",
		},
		{
			insert:        "catt",
			rootCandidate: "cattle",
			expectedRoot:  "catt",
		},
	}

	for _, c := range cases {
		if c.insert != "" {
			tr.Insert(c.insert)
		}

		if c.search != "" {
			assert.Equalf(t, c.expectedSearch, tr.Search(c.search), fmt.Sprintf("search %s", c.search))
		}

		if c.prefix != "" {
			assert.Equalf(t, c.expectedPrefix, tr.StartsWith(c.prefix), fmt.Sprintf("prefix %s", c.prefix))
		}

		if c.rootCandidate != "" {
			assert.Equalf(t, c.expectedRoot, tr.BestRoot(c.rootCandidate), fmt.Sprintf("root %s", c.rootCandidate))
		}
	}
}
