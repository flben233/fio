//go:build windows && !(amd64 || 386)
// +build windows,!amd64,!386

package fio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetFIO 返回适用于当前系统的 fio 命令字符串和临时文件路径（用于清理）
func GetFIO() (fioCmd string, tempFile string, err error) {
	var errors []string
	// 1. 尝试系统自带 fio
	if path, lookErr := exec.LookPath("fio.exe"); lookErr == nil {
		// 直接尝试 fio.exe (Windows 通常不使用 sudo)
		testCmd := exec.Command(path, "--help")
		if runErr := testCmd.Run(); runErr == nil {
			return path, "", nil
		} else {
			errors = append(errors, fmt.Sprintf("fio.exe 运行失败: %v", runErr))
		}
	} else if path, lookErr := exec.LookPath("fio"); lookErr == nil {
		// 尝试不带 .exe 后缀
		testCmd := exec.Command(path, "--help")
		if runErr := testCmd.Run(); runErr == nil {
			return path, "", nil
		} else {
			errors = append(errors, fmt.Sprintf("fio 运行失败: %v", runErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("无法找到 fio: %v", lookErr))
	}
	// 返回所有错误信息
	return "", "", fmt.Errorf("无法找到可用的 fio 命令:\n%s", strings.Join(errors, "\n"))
}

// ExecuteFIO 执行拼好的 fio 命令字符串
func ExecuteFIO(fioCmd string, args []string) error {
	// Windows 版本不使用 sh -c 方式执行，而是直接执行命令
	cmd := exec.Command(fioCmd, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
