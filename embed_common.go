package fio

import (
	"os"
	"path/filepath"
)

// CleanFio 删除临时提取出的 fio 文件
func CleanFio(tempFile string) error {
	if tempFile == "" {
		return nil // 不需要清理
	}
	// 删除整个临时目录
	return os.RemoveAll(filepath.Dir(tempFile))
}
