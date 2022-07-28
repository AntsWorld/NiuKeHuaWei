// 数独
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x     int
	y     int
	value int
}

func main() {
	// 从标准输入读取数独数据
	suDoData := ReadSuDoDataFromStdIn()
	// 解析空白坐标数据
	points := ParseEmptyPoints(suDoData)
	fmt.Println(points)
}

func ParseEmptyPoints(data [9][9]int) []Point {
	points := make([]Point, 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if data[i][j] == 0 {
				point := parsePointData(i, j, data)
				points = append(points, point)
			}
		}
	}
	return points
}

// 解析单个空白位的值
func parsePointData(i int, j int, data [9][9]int) Point {
	point := Point{}
	// 获取行空白位可能的值

	// 获取列空白位可能的值

	// 获取3*3区域空白位可能的值

	// 根据三个可能值推测空白位的值

	return point
}

func ReadSuDoDataFromStdIn() [9][9]int {
	sd := [9][9]int{}
	reader := bufio.NewReader(os.Stdin)
	// 读取9*9输入信息
	line := 9
	for i := 0; i < line; i++ {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			os.Exit(1)
		}
		byteStr := strings.TrimSpace(string(bytes))
		split := strings.Split(byteStr, " ")
		if len(split) != 9 {
			os.Exit(1)
		}
		for j := 0; j < len(split); j++ {
			val, err := strconv.Atoi(split[j])
			if err != nil {
				os.Exit(1)
			}
			sd[i][j] = val
		}
	}
	return sd
}
