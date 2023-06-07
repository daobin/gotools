package internal

import (
	"io/fs"
	"os"
	"path"
	"strings"
)

type fileTool struct{}

// CreateDir 创建目录（不存在时创建，且可创建多级）
// @dirname 目录名
// @mode 目录权限
func (receiver fileTool) CreateDir(dirname string, mode fs.FileMode) error {

	_, err := os.Stat(dirname)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dirname, mode)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

// GetFileExt 获取文件格式，如：.txt / .png / .xlsx 等
// @filename 文件名
func (receiver fileTool) GetFileExt(filename string) string {
	return path.Ext(filename)
}

// CheckFileExtValid 校验文件格式是否有效
// @filename 文件名
// @validExt 有效的文件格式
func (receiver fileTool) CheckFileExtValid(filename string, validExt []string) bool {
	fileExt := strings.ToLower(strings.Trim(path.Ext(filename), "."))
	validExt = Slice.ToLower(validExt)

	for _, ext := range validExt {
		ext = strings.Trim(ext, ".")
		if ext == fileExt {
			return true
		}
	}

	return false
}
