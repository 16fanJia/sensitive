package trie

import "fmt"

/*
	使用前缀树的方法构建敏感词库
*/

type Trie struct {
	Root *Node
}

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool //是否为根节点
	isPathEnd  bool //是否完结
	character  rune //词
	children   map[rune]*Node
}

//NewTrie 新建前缀树
func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			isRootNode: true,
			character:  0,
			children:   make(map[rune]*Node, 0),
		},
	}
}

//newNode 新建子节点
func newNode(character rune) *Node {
	return &Node{
		character: character,
		children:  make(map[rune]*Node, 0),
	}
}

//Add 添加敏感词
func (t *Trie) Add(word string) {
	var (
		current *Node
		runes   []rune
	)
	current = t.Root //初始化为根目录
	runes = []rune(word)

	//fmt.Printf("debug runes %d\n", len(runes))

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		//从根节点查看是否存在前缀
		if next, ok := current.children[r]; ok {
			current = next
		} else {
			//新建一个节点 放在根节点下
			node := newNode(r)
			current.children[r] = node
			current = node
		}
		//判断是否为最后一个字符
		if i == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

//Del 删除敏感词
func (t *Trie) Del(word string) {
	var (
		current *Node
		runes   []rune
	)
	current = t.Root //初始化为根目录
	runes = []rune(word)

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if next, ok := current.children[r]; ok {
			current = next
		} else {
			//未找到对应的敏感词
			fmt.Printf("库中未找到对应的敏感词！%s", word)
			return
		}

		if i == len(runes)-1 {
			current.isPathEnd = false
		}
	}
}

//Replace 替换敏感词
func (t *Trie) Replace(txt string, replace rune) string {
	var (
		parent  = t.Root
		current *Node
		runes   = []rune(txt)
		left    = 0 //记录敏感词开始的位置
		found   bool
	)

	for i := 0; i < len(runes); i++ {
		//从父节点下的字节点开始循环查找
		if current, found = parent.children[runes[i]]; !found || (i == len(runes)-1 && !current.isPathEnd) {
			parent = t.Root
			i = left
			left++
			continue
		}

		if current.isPathEnd && left <= i {
			for j := left; j <= i; j++ {
				runes[j] = replace
			}
		}

		parent = current
	}
	return string(runes)
}
