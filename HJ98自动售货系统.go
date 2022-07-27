package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 商品信息
type Good struct {
	Name  string // 商品名称(A1 A2 A3 A4 A5)[名称不重复]
	Price int    // 单价(题目中都是整数使用int)
	Num   int    // 数量
}

// 存钱盒信息
type MoneyBox struct {
	Denominations int // 面额(1元 2元 5元 10元)
	Num           int // 张数
}

// 售货系统
type SaleSys struct {
	GoodMap     map[string]Good  // 商品(key为商品名称)
	MoneyBoxMap map[int]MoneyBox // 存钱盒(key为钱币面额)
	InMoney     int              // 投币金额
	YuMoney     int              // 可用余额
}

func main() {
	ch := make(chan string)
	// 启动售货系统
	go startSaleSystem(ch)
	// 读取命令行输入
	for true {
		reader := bufio.NewReader(os.Stdin)
		cmdData, _, err := reader.ReadLine()
		if err != nil {
			os.Exit(1)
		}
		cmdArr := strings.Split(string(cmdData), ";")
		for _, cmd := range cmdArr {
			ch <- strings.TrimSpace(cmd)
		}
	}
}

// 启动销售系统
func startSaleSystem(ch chan string) {
	saleSys := &SaleSys{}
	for true {
		// 读取命令
		cmd := <-ch
		cmd = strings.TrimSpace(cmd)
		fmt.Println(cmd)
		if strings.HasPrefix(cmd, "r") {
			rCmdHandler(cmd, saleSys)
		} else if strings.HasPrefix(cmd, "c") {
			cCmdHandler(cmd, saleSys)
		} else if strings.HasPrefix(cmd, "q") {
			qCmdHandler(cmd, saleSys)
		} else if strings.HasPrefix(cmd, "p") {
			pCmdHandler(cmd, saleSys)
		} else if strings.HasPrefix(cmd, "b") {
			bCmdHandler(cmd, saleSys)
		} else {
			unknownCmdHandler(cmd, saleSys)
		}

	}
}

func unknownCmdHandler(cmd string, sys *SaleSys) {

}

// b A5
func bCmdHandler(cmd string, sys *SaleSys) {

}

// p 5
func pCmdHandler(cmd string, sys *SaleSys) {

}

// q1
func qCmdHandler(cmd string, sys *SaleSys) {

}

// c
func cCmdHandler(cmd string, sys *SaleSys) {

}

// r 22-18-21-21-7-20 3-23-10-6
func rCmdHandler(cmd string, sys *SaleSys) {
	// 解析命令数据
	split := strings.Split(strings.TrimSpace(cmd), " ")
	if len(split) != 3 {
		return
	}
	// 解析商品数量
	// 解析钱币面额张数

}
