package main

import (
	"strings"
	"time"
)

//消息
type ChatMessage struct {
	Text string//消息内容
	SendTime time.Time//发送时间
	SendName string//发送账号
}

//https://github.com/RobertJGabriel/Google-profanity-words/blob/ma ster/list.txt
//网址无法访问 写了个测试的字符串
var ProfanityWords = [] string {"Word0","Word1","Word2","Word3" }

//替换屏蔽字符串
func ProfanityFilter(old string) string {
	var text = old
	for _, dword := range ProfanityWords {
		text = strings.Replace(text, dword, strings.Repeat("*", len(dword)), -1)
	}
	return text
}

//Popular
func (user GameUser) Popular() {
	var maxkey = ""
	mapword := make(map[string]int)
	dts, _ := time.ParseDuration("-5s")
	var drove = time.Now().Add(dts)
	var need = CheckMessageNeedPop(drove)
	for need {
		PopMessage()
		need = CheckMessageNeedPop(drove)
	}
	var maxvalue = 0
	for _, v := range QueueMessage {
		var keys = MessageSplitToWord(v.Text)
		for _, vk := range keys {
			var amount, ok = mapword [vk]
			var count = 1
			if ok {
				count = amount + 1
			}
			mapword[vk] = count
			if count > maxvalue {
				maxkey = vk
				maxvalue = count
			}
		}
	}
	user.Conn.Write([]byte(maxkey))
}
//分割聊天信息为单词
func MessageSplitToWord(text string) []string{
	f := func(c rune) bool {
		if c == '*' || c == ','||c==';'||c==' '||c=='.' {
			return true
			} else {
			return false
			}
		}
	result := strings.FieldsFunc(text, f)
	return  result
}