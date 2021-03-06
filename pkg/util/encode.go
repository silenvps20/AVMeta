package util

import (
	"encoding/base64"
	"os"
)

// Base64 将指定文件编码为 Base64 字符串，
// 返回编码信息及错误信息。
//
// file 字符串参数，传入文件路径
func Base64(file string) (string, error) {
	// 检查错误
	f, err := os.Open(file)
	// 如果出错
	if err != nil {
		return "", err
	}
	// 关闭
	defer f.Close()

	// 初始化byte
	buff := make([]byte, 500000)
	// 读取文件
	n, err := f.Read(buff)
	// 检查错误
	if err != nil {
		return "", err
	}

	// Base64编码
	source := base64.StdEncoding.EncodeToString(buff[:n])

	return source, nil
}
