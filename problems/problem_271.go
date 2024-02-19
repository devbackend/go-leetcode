package problems

import "strings"

const delimiter = " "

type Codec struct{}

// Encode Encodes a list of strings to a single string.
func (c *Codec) Encode(strs []string) string {
	return strings.Join(strs, delimiter)
}

// Decode Decodes a single string to a list of strings.
func (c *Codec) Decode(strs string) []string {
	return strings.Split(strs, delimiter)
}
