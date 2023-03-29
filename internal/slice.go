package internal

import (
	"reflect"
	"strings"
)

type sliceTool struct{}

// IndexOf 返回切片中指定元素的位置，未找到时返回 -1
func (receiver sliceTool) IndexOf(data []any, elem any) int {
	if data == nil || len(data) == 0 {
		return -1
	}

	if reflect.TypeOf(data[0]) != reflect.TypeOf(elem) {
		return -1
	}

	for idx, val := range data {
		if val == elem {
			return idx
		}
	}

	return -1
}

// Intersect 求两个切片的交集
// @data1 切片数据
// @data2 切片数据
func (receiver sliceTool) Intersect(data1, data2 []any) []any {
	intersect := make([]any, 0)
	if data1 == nil || data2 == nil || len(data1) == 0 || len(data2) == 0 {
		return intersect
	}

	if reflect.TypeOf(data1[0]) != reflect.TypeOf(data2[0]) {
		return intersect
	}

	dataMap := make(map[any]any)
	for _, dt := range data1 {
		dataMap[dt] = dt
	}

	for _, dt := range data2 {
		if _, ok := dataMap[dt]; ok {
			intersect = append(intersect, dt)
		}
	}

	return intersect
}

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
