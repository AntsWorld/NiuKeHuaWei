package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 读取第一行数据
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		return
	}
	lineData := strings.ToLower(string(line))
	// 读取第二行数据
	matchKey, _, err := reader.ReadLine()
	if err != nil {
		return
	}
	matchKeyData := strings.ToLower(string(matchKey))
	// 数据范围判断
	min, max := 1, 1000
	n := len(lineData)
	if n < min || n > max {
		fmt.Println("数据范围：1 ≤ n ≤ 1000")
		return
	}
	// 将字符放入map
	letters := make(map[string]int)
	for i := 0; i < len(lineData); i++ {
		letter := fmt.Sprintf("%c", lineData[i])
		if _, ok := letters[letter]; ok {
			letters[letter] = letters[letter] + 1
		} else {
			letters[letter] = 1
		}
	}
	if _, ok := letters[matchKeyData]; ok {
		fmt.Println(letters[matchKeyData])
	} else {
		fmt.Println(0)
	}
}
