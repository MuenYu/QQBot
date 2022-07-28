package util

import "strings"

// 字符串工具类

// InstructExtractor 从字符串中截取命令
func InstructExtractor(content string) string {
	var index = strings.LastIndex(content, "#")
	if index < 0 {
		return ""
	}
	return content[strings.LastIndex(content, "#")+1:]
}
