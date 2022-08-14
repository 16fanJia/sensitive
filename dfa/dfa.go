package dfa

import "fmt"

/*
dfa :使用dfa算法构建敏感词map

敏感词数据结构 eg：
{
	isEnd: false
	王：
	{
		isEnd: false
		八：
		{
			isEnd：false
			蛋：
			{
				isEnd：true
			}
		}
	}
}
无意义符号数据结构：
{
	"@":Null (空结构体)
}
*/

type DFA struct {
	sensitiveWord map[string]interface{}
	invalidWord   map[string]interface{}
}

var InstanceDfa *DFA

//饿汉式加载单例
func init() {
	InstanceDfa = &DFA{
		sensitiveWord: make(map[string]interface{}, 0), //敏感词的map
		invalidWord:   make(map[string]interface{}),    //无效词汇，不参与敏感词汇判断直接忽略
	}
}

//Add 向敏感词库中添加违禁词  生成违禁词集合
func (d *DFA) Add(word string) {
	runes := []rune(word)
	nowMap := d.sensitiveWord
	for i := 0; i < len(runes); i++ {
		if _, ok := nowMap[string(runes[i])]; !ok { //如果该key不存在
			thisMap := make(map[string]interface{})
			thisMap["isEnd"] = false
			nowMap[string(runes[i])] = thisMap
			nowMap = thisMap
		} else {
			nowMap = nowMap[string(runes[i])].(map[string]interface{})
		}
		if i == len(runes)-1 {
			nowMap["isEnd"] = true
		}
	}
}

//Del 删除敏感词
func (d *DFA) Del(word string) {
	runes := []rune(word)
	nowMap := d.sensitiveWord
	for i := 0; i < len(runes); i++ {
		//判断敏感词map中 是否存在
		if _, ok := nowMap[string(runes[i])]; ok {
			nowMap = nowMap[string(runes[i])].(map[string]interface{})
		} else {
			//未找到对应的敏感词
			fmt.Printf("敏感词库中未找到对应的敏感词！%s", word)
			return
		}

		if i == len(runes)-1 {
			nowMap["isEnd"] = false
		}
	}
}

//Replace 将敏感词汇转换为rep
func (d *DFA) Replace(txt string, rep rune) (word string) {
	var (
		runes  = []rune(txt)
		nowMap = d.sensitiveWord
		start  = -1
		tag    = -1
	)

	//从第一个字符开始遍历
	for i := 0; i < len(runes); i++ {
		if _, ok := d.invalidWord[(string(runes[i]))]; ok {
			continue //如果是无效词汇直接跳过
		}
		if thisMap, ok := nowMap[string(runes[i])].(map[string]interface{}); ok {
			//记录敏感词第一个文字的位置
			tag++
			if tag == 0 {
				start = i
			}
			//判断是否为敏感词的最后一个文字
			if isEnd, _ := thisMap["isEnd"].(bool); isEnd {
				//将敏感词的第一个文字到最后一个文字全部替换为指定rune
				for y := start; y < i+1; y++ {
					runes[y] = rep
				}
				//重置标志数据
				nowMap = d.sensitiveWord
				start = -1
				tag = -1
			} else { //不是最后一个，则将其包含的map赋值给nowMap
				nowMap = nowMap[string(runes[i])].(map[string]interface{})
			}
		} else { //如果敏感词不是全匹配，则终止此敏感词查找。从开始位置的第二个文字继续判断
			if start != -1 {
				i = start
			}
			//重置标志参数
			nowMap = d.sensitiveWord
			start = -1
			tag = -1
		}
	}
	return string(runes)
}
