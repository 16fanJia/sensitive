package trie

import (
	"testing"
)

func TestFilter_Replace(t *testing.T) {
	tree := NewTrie()
	tree.Add("王八蛋")

	str := "任小玉王八蛋"
	tree.Replace(str, '*')

}
