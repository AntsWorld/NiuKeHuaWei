package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	// 输入随机数个数
	length := 0
	if _, err := fmt.Scanf("%d", &length); err != nil {
		os.Exit(1)
	}
	// 判断数据范围
	minLen, maxLen := 1, 1000
	minValue, maxValue := 1, 500
	if length < minLen || length > maxLen {
		os.Exit(1)
	}
	// 输入随机数
	var numberSli = make(map[int]struct{}, 0)
	for i := 0; i < length; i++ {
		numberValue := 0
		if _, err := fmt.Scanf("%d", &numberValue); err != nil {
			os.Exit(1)
		}
		// 判断输入值范围
		if numberValue < minValue || numberValue > maxValue {
			os.Exit(1)
		}
		numberSli[numberValue] = struct{}{}
	}
	var res []int
	for k, _ := range numberSli {
		res = append(res, k)
	}
	sort.Ints(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
