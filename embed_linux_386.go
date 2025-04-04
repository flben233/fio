//go:build linux && 386
// +build linux,386

package fio

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed bin/fio-linux-386
var binFiles embed.FS

// GetFioBinary 获取与当前系统匹配的 fio 二进制文件
func GetFioBinary() (string, error) {
	binaryName := "fio-linux-386"
	// 创建临时目录存放二进制文件
	tempDir, err := os.MkdirTemp("", "disktest")
	if err != nil {
		return "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	// 读取嵌入的二进制文件
	binPath := filepath.Join("bin", binaryName)
	fileContent, err := binFiles.ReadFile(binPath)
	if err != nil {
		return "", fmt.Errorf("读取嵌入的 fio 二进制文件失败: %v", err)
	}
	// 写入临时文件
	tempFile := filepath.Join(tempDir, binaryName)
	if err := os.WriteFile(tempFile, fileContent, 0755); err != nil {
		return "", fmt.Errorf("写入临时 fio 文件失败: %v", err)
	}

	return tempFile, nil
}
