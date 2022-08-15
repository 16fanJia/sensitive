package sensitive

import (
	"bufio"
	"os"
	"sensitive/dfa"
	"sensitive/iface"
	"sensitive/trie"
)

/*
	敏感词库的入口
*/

//模式
const (
	DFA_MODEL = iota + 1
	TRIE_MODEL
)

type Filter struct {
	model int //1:DFA 2:TRIE
	trie  *trie.Trie
	dfa   *dfa.DFA
}

//NewFilter 构造函数
func NewFilter(model int) (filter iface.SensitiveFilter) {
	switch model {
	case TRIE_MODEL:
		filter = &Filter{
			model: TRIE_MODEL,
			trie:  trie.NewTrie(),
			dfa:   nil,
		}
	case DFA_MODEL:
		filter = &Filter{
			model: DFA_MODEL,
			trie:  nil,
			dfa:   dfa.InstanceDfa,
		}
	}
	return filter
}

//AddSensitiveWords 可以添加若干个words
func (f *Filter) AddSensitiveWords(words ...string) {
	switch f.model {
	case TRIE_MODEL:
		for _, word := range words {
			f.trie.Add(word)
		}
	case DFA_MODEL:
		for _, word := range words {
			f.dfa.Add(word)
		}
	}
}

func (f *Filter) DelSensitiveWords(words ...string) {
	switch f.model {
	case TRIE_MODEL:
		for _, word := range words {
			f.trie.Del(word)
		}
	case DFA_MODEL:
		for _, word := range words {
			f.dfa.Del(word)
		}
	}
}

// Replace 和谐敏感词
func (f *Filter) Replace(text string, repl rune) string {
	var result string
	switch f.model {
	case DFA_MODEL:
		result = f.dfa.Replace(text, repl)
	case TRIE_MODEL:
		result = f.trie.Replace(text, repl)
	}
	return result
}

func (f *Filter) LoadWordDict(path string) error {
	var (
		err  error
		file *os.File
	)
	if file, err = os.Open(path); err != nil {
		return err
	}

	defer file.Close()

	//加载文件内容
	if err = f.load(file); err != nil {
		return err
	}
	return nil
}

func (f *Filter) load(file *os.File) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		//添加到词库中
		switch f.model {
		case DFA_MODEL:
			f.dfa.Add(line)
		case TRIE_MODEL:
			f.trie.Add(line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
