package iface

/*
	暴露给用户的外部接口
*/

type SensitiveFilter interface {
	//LoadWordDict 加载敏感字典
	LoadWordDict(string) error
	//AddSensitiveWords 添加敏感词
	AddSensitiveWords(...string)
	//DelSensitiveWords 删除敏感词
	DelSensitiveWords(...string)
	//Replace 和谐敏感词
	Replace(string, rune) string
}
