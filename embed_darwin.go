//go:build darwin && !(amd64 || arm64)
// +build darwin,!amd64,!arm64

package fio

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetFIO 返回适用于当前系统的 fio 命令字符串（带不带 sudo）和临时文件路径（用于清理）
func GetFIO() (fioCmd string, tempFile string, err error) {
	var errors []string
	// 1. 尝试系统自带 fio
	if path, lookErr := exec.LookPath("fio"); lookErr == nil {
		// 尝试 sudo fio
		testCmd := exec.Command("sudo", path, "--help")
		if runErr := testCmd.Run(); runErr == nil {
			return "sudo fio", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("sudo fio 测试失败: %v", runErr))
		}
		// 直接尝试 fio
		testCmd = exec.Command(path, "--help")
		if runErr := testCmd.Run(); runErr == nil {
			return "fio", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("fio 直接运行失败: %v", runErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("无法找到 fio: %v", lookErr))
	}
	// 返回所有错误信息
	return "", "", fmt.Errorf("无法找到可用的 fio 命令:\n%s", strings.Join(errors, "\n"))
}

// ExecuteFIO 执行拼好的 fio 命令字符串（包括 sudo、fio 等）
func ExecuteFIO(fioCmd string, args []string) error {
	// 拼接命令字符串
	fullCmd := fmt.Sprintf("%s %s", fioCmd, strings.Join(args, " "))
	cmd := exec.Command("sh", "-c", fullCmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}