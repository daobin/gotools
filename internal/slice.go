package internal

import "strings"

type sliceTool struct{}

// ToLower 将切片中的元素转为小写
// @data 需要转为小写的切片数据
func (receiver sliceTool) ToLower(data []string) []string {
	if data == nil {
		return []string{}
	}

	for idx, val := range data {
		data[idx] = strings.ToLower(val)
	}

	return data
}
