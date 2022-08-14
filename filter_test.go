package sensitive

import (
	"fmt"
	"testing"
)

func TestNewFilter(t *testing.T) {
	f := NewFilter(DFA_MODEL)
	err := f.LoadWordDict("/Users/fanjia/Desktop/sensitive/doc/sensitiveDict.txt")
	if err != nil {
		fmt.Println("加载文件失败=====》", err)
	}

	res := f.Replace("AV傻逼大傻逼", '*')
	fmt.Println(res)

	f1 := NewFilter(TRIE_MODEL)
	err1 := f1.LoadWordDict("/Users/fanjia/Desktop/sensitive/doc/sensitiveDict.txt")
	if err1 != nil {
		fmt.Println("加载文件失败=====》", err1)
	}

	res1 := f1.Replace("AV傻逼大傻逼", '*')
	fmt.Println(res1)

}
