package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// s是结果，sep是分隔符
	s, sep := "", ""
	// os包提供了处理操作系统相关的函数/值
	// os.Args 获得命令行参数，返回[] string 字符串切片
	// [0]第一个值是命令本身
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// strings.join 类似于 python ''.join
	// 把string集合组合成一个字符串
	fmt.Printf("strings.Join: %s ", strings.Join(os.Args[1:], " "))
}
