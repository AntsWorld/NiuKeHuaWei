package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// HJ1 字符串最后一个单词的长度
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(0)
		return
	}
	value := strings.TrimSpace(string(line))
	words := strings.Split(value, " ")
	if len(words) <= 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(len(words[len(words)-1]))
}
