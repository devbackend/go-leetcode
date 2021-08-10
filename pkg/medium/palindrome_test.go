package medium_test

import (
	"github.com/devbackend/go-leetcode/pkg/medium"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		str      string
		expected string
	}{
		{
			name:     "example 1",
			str:      "babad",
			expected: "bab",
		},
		{
			name:     "example 2",
			str:      "cbbd",
			expected: "bb",
		},
		{
			name:     "example 3",
			str:      "a",
			expected: "a",
		},
		{
			name:     "example 4",
			str:      "ac",
			expected: "a",
		},
		{
			name:     "example 5",
			str:      "busislnescsicxpvvysuqgcudefrfjbwwjcchtgqyajdfwvkypfwshnihjdztgmyuuljxgvhdiwphrweyfkbnjgerkmifbirubhseuhrugwrabnjafnbdfjnufdstjbkuwtnpflffaqmjbhssjlnqftgjiglvvequhapasarlkcvbmkwnkuvwktbgfoaxteprobdwswcdyddyvrehvmxrrjiiidatidlpihkbmmruysmhhsncmfdanafdrfpdtfgkglcqpwrrtvacuicohspkounojuziittugpqjyhhkwfnflozbispehrtrnizowrlzcuollagxwtznjwzcumvedjwokueuqktvvouwnsmpxqvvpuwprezrbobrpnwaccwljchdguubjulyilzvmandjjleitweybqkjttschrjjlebnmponvlktzzcdtuybugggcqffkcffpamauvxfbonjrobgpvlyzveiwemmtdvbjciaytvesnocnjrwodtcokgcuoiicxapmrzpkfphjniuvzjrhbnqndfshoduejyktebgdabidxlkstepuwvtrtgbxaeheylicvhrxddijshcvdadxzsccmainyfpfdhqdanfccqkzlmhsfilvoybqojlvbcixjzqpbngdvesuokbxhkomsiqfyukvspqthlzxdnlwthrgaxhtpjzhrugqbfokrdcyurivmzgtynoqfjbafboselxnfupnpqlryvlcxeksirvufepfwczosrrjpudbwqxwldgjyfjhzlzcojxyqjyxxiqvfhjdwtgoqbyeocffnyxhyyiqspnvrpxmrtcnviukrjvpavervvztoxajriuvxqveqsrttjqepvvahywuzwtmgyrzduxfqspeipimyoxmkadrvrdyefekjxcmsmzmtbugyckcbjsrymszftjyllfmoeoylzeahnrxlxpnlvlvzltwnmldi",
			expected: "rbobr",
		},
		{
			name:     "example 6",
			str:      "abcde",
			expected: "a",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := medium.LongestPalindrome(c.str)

			assert.Equal(t, c.expected, actual)
		})
	}
}
