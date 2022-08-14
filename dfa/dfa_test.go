package dfa

import (
	"fmt"
	"strings"
	"testing"
)

func TestDFA(t *testing.T) {
	words := strings.Split(InvalidWords, ",")
	for _, v := range words {
		InstanceDfa.Add(v)
	}
	sensitive := []string{"我日", "傻", "逼", "日"}
	for _, v := range sensitive {
		InstanceDfa.Add(v)
	}

	text := "文明用语你&* 妈, 逼的你这个狗 日的，怎么这么傻啊。我也是服了，我日,这些话我都说不出口"
	fmt.Println(InstanceDfa.Replace(text, '*'))

}
