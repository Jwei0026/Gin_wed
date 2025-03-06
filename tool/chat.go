package tool

import (
	"strings"
)

// containsStringInSlice 判断字符串内是否含有切片内某一元素，如果有则返回该元素
func ContainsStringInSlice(str string) (string, bool) {
	var slice = []string{"系列", "表壳", "表底", "颜色", "功能", "款式", "尺寸", "表镜", "表扣", "年份", "机芯", "厚度", "表盘", "防水"}
	for _, item := range slice { //便利切片里的内容
		if strings.Contains(str, item) { //将切片内容与字符串进行匹配
			return item, true
		}
	}
	return "", false
}

// chineseToEnglish 是一个映射表，存储中文词语到英文单词的对应关系
var chineseToEnglish = map[string]string{ //这里将关键词转成目标属性
	"系列": "series_id",
	"表壳": "case_material",
	"表底": "case_back",
	"颜色": "strap_color",
	"功能": "functions",
	"款式": "style",
	"尺寸": "size",
	"表镜": "watch_glass",
	"表扣": "watch_buckle",
	"年份": "launch_year",
	"机芯": "movement",
	"厚度": "thickness",
	"表盘": "dial",
	"防水": "water_resistance",
}

// translateChineseToEnglish 将中文词语转换为对应的英文单词
func TranslateChineseToEnglish(chineseWord string) (string, bool) {
	// 在映射表中查找中文词语对应的英文单词
	englishWord, found := chineseToEnglish[chineseWord]
	return englishWord, found
}
