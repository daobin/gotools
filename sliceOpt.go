package gotools

type sliceOpt struct{}

// IndexOfInt 返回 int 切片中指定元素的位置，未找到时返回 -1
func (opt sliceOpt) IndexOfInt(data []int, element int) int {
	for idx, dt := range data {
		if dt == element {
			return idx
		}
	}

	return -1
}

// IntersectInt 求两个 int 切片的交集
func (opt sliceOpt) IntersectInt(data1, data2 []int) []int {
	intersect := make([]int, 0)
	if len(data1) == 0 || len(data2) == 0 {
		return intersect
	}

	dataMap := make(map[int]int)
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

// MergeInt 求两个 int 切片的并集
func (opt sliceOpt) MergeInt(data1, data2 []int) []int {
	return append(data1, data2...)
}

// DiffInt 求两个 int 切片的差集（在第一个切片中，但不在第二个切片中）
func (opt sliceOpt) DiffInt(data1, data2 []int) []int {
	diff := make([]int, 0)
	if len(data1) == 0 {
		return diff
	}

	dataMap := make(map[int]int)
	for _, dt := range data2 {
		dataMap[dt] = dt
	}

	for _, dt := range data1 {
		if _, ok := dataMap[dt]; !ok {
			diff = append(diff, dt)
		}
	}

	return diff
}

// IndexOfString 返回 string 切片中指定元素的位置，未找到时返回 -1
func (opt sliceOpt) IndexOfString(data []string, element string) int {
	for idx, dt := range data {
		if dt == element {
			return idx
		}
	}

	return -1
}

// IntersectString 求两个 string 切片的交集
func (opt sliceOpt) IntersectString(data1, data2 []string) []string {
	intersect := make([]string, 0)
	if len(data1) == 0 || len(data2) == 0 {
		return intersect
	}

	dataMap := make(map[string]string)
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

// MergeString 求两个 string 切片的并集
func (opt sliceOpt) MergeString(data1, data2 []string) []string {
	return append(data1, data2...)
}

// DiffString 求两个 string 切片的差集（在第一个切片中，但不在第二个切片中）
func (opt sliceOpt) DiffString(data1, data2 []string) []string {
	diff := make([]string, 0)
	if len(data1) == 0 {
		return diff
	}

	dataMap := make(map[string]string)
	for _, dt := range data2 {
		dataMap[dt] = dt
	}

	for _, dt := range data1 {
		if _, ok := dataMap[dt]; !ok {
			diff = append(diff, dt)
		}
	}

	return diff
}
