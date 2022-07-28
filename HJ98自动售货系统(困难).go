package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 商品信息
type Good struct {
	Name  string  // 商品名称(A1 A2 A3 A4 A5)[名称不重复]
	Price float32 // 单价
	Num   int     // 数量
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
		bytes, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		command := strings.TrimSpace(string(bytes))
		cmdArr := strings.Split(command, ";")
		for _, cmd := range cmdArr {
			ch <- strings.TrimSpace(cmd)
		}
	}
}

// 启动销售系统
func startSaleSystem(ch chan string) {
	saleSys := getBaseSaleSys()
	for true {
		// 读取命令
		cmd := <-ch
		cmd = strings.TrimSpace(cmd)
		fmt.Print(cmd, " --> ")
		if strings.HasPrefix(cmd, "r") {
			saleSys.rCmdHandler(cmd)
		} else if strings.HasPrefix(cmd, "c") {
			saleSys.cCmdHandler(cmd)
		} else if strings.HasPrefix(cmd, "q") {
			saleSys.qCmdHandler(cmd)
		} else if strings.HasPrefix(cmd, "p") {
			saleSys.pCmdHandler(cmd)
		} else if strings.HasPrefix(cmd, "b") {
			bCmdHandler(cmd, saleSys)
		} else {
			unknownCmdHandler(cmd, saleSys)
		}

	}
}

func getBaseSaleSys() *SaleSys {
	saleSys := SaleSys{
		GoodMap: map[string]Good{
			"A1": {
				Name:  "A1",
				Price: 2,
				Num:   0,
			},
			"A2": {
				Name:  "A2",
				Price: 3,
				Num:   0,
			},
			"A3": {
				Name:  "A3",
				Price: 4,
				Num:   0,
			},
			"A4": {
				Name:  "A4",
				Price: 5,
				Num:   0,
			},
			"A5": {
				Name:  "A5",
				Price: 8,
				Num:   0,
			},
			"A6": {
				Name:  "A6",
				Price: 6,
				Num:   0,
			},
		},
		MoneyBoxMap: map[int]MoneyBox{
			10: {
				Denominations: 10,
				Num:           0,
			},
			5: {
				Denominations: 5,
				Num:           0,
			},
			2: {
				Denominations: 2,
				Num:           0,
			},
			1: {
				Denominations: 1,
				Num:           0,
			},
		},
		InMoney: 0,
		YuMoney: 0,
	}
	return &saleSys
}

func unknownCmdHandler(cmd string, sys *SaleSys) {

}

// b A5
func bCmdHandler(cmd string, sys *SaleSys) {

}

// 投币指令
// 命令格式：p 钱币面额
// 功能说明：
//（1） 如果投入非1元、2元、5元、10元的钱币面额（钱币面额不考虑负数、字符等非正整数的情况），输出“E002:Denomination error”；
//（2） 如果存钱盒中1元和2元面额钱币总额小于本次投入的钱币面额，输出“E003:Change is not enough, pay fail”，但投入1元和2元面额钱币不受此限制。
//（3） 如果自动售货机中商品全部销售完毕，投币失败。输出“E005:All the goods sold out”；
//（4） 如果投币成功，输出“S002:Pay success,balance=X”；
// 约束说明：
//（1） 系统在任意阶段都可以投币；
//（2） 一次投币只能投一张钱币；
//（3） 同等条件下，错误码的优先级：E002 > E003 > E005；
// 输出说明：如果投币成功，输出“S002:Pay success,balance=X”。
func (saleSys *SaleSys) pCmdHandler(cmd string) {
	// 解析命令参数
	split := strings.Split(strings.TrimSpace(cmd), " ")
	if len(split) != 2 {
		return
	}
	val, err := strconv.Atoi(split[1])
	if err != nil {
		return
	}
	// 同等条件下，错误码的优先级：E002 > E003 > E005；
	switch val {
	case 1, 2, 5, 10:
		// 判断1元和2元面额钱币总额是否小于本次投入的钱币面额
		if val > 2 && saleSys.inMoneyBiggerThanOneAndTwo(val) {
			fmt.Println("E003:Change is not enough, pay fail")
			return
		}
		// 判断商品是否已售完
		if saleSys.isAllGoodSoldOut() {
			fmt.Println("E005:All the goods sold out")
			return
		}
		// 投币成功
		saleSys.onInMoneySuc(val)
		fmt.Printf("S002:Pay success,balance=%d\n", saleSys.YuMoney)
	default:
		fmt.Println("E002:Denomination error")
		return
	}
}

// 投币成功处理逻辑
func (saleSys *SaleSys) onInMoneySuc(val int) {
	saleSys.InMoney = val + saleSys.InMoney
	saleSys.YuMoney = val + saleSys.YuMoney
	if box, ok := saleSys.MoneyBoxMap[val]; ok {
		box.Num++
		saleSys.MoneyBoxMap[val] = box
	} else {
		moneyBox := MoneyBox{
			Denominations: val,
			Num:           1,
		}
		saleSys.MoneyBoxMap[val] = moneyBox
	}
}

// 查询指令
// 命令格式：q 查询类别
// 功能说明：
//（1） 查询自动售货机中商品信息，包含商品名称、单价、数量。 根据商品数量从大到小进行排序；商品数量相同时，按照商品名称的先后顺序进行排序 。
func (saleSys *SaleSys) qCmdHandler(cmd string) {
	// 解析命令
	split := strings.Split(strings.TrimSpace(cmd), " ")
	if len(split) != 2 {
		fmt.Println("E010:Parameter error")
		return
	}
	tp, err := strconv.Atoi(split[1])
	if err != nil {
		return
	}
	switch tp {
	case 0:
		sortAndPrintGoodInfo(saleSys.GoodMap)
	case 1:
		printMoneyBoxInfo(saleSys.MoneyBoxMap)
	default:
		fmt.Println("E010:Parameter error")
		return
	}
}

func printMoneyBoxInfo(boxMap map[int]MoneyBox) {

}

func sortAndPrintGoodInfo(goodMap map[string]Good) {

}

// 退币指令
// c
func (saleSys *SaleSys) cCmdHandler(cmd string) {
	// 如果投币余额等于0的情况下，输出“E009:Work failure”；
	if saleSys.YuMoney == 0 {
		fmt.Println("E009:Work failure")
		return
	}
	// 如果投币余额大于0的情况下，按照 退币原则 进行“找零”，输出退币信息；
	// TODO
}

//系统初始化指令
// r 22-18-21-21-7-20 3-23-10-6
func (saleSys *SaleSys) rCmdHandler(cmd string) {
	// 解析命令数据
	split := strings.Split(strings.TrimSpace(cmd), " ")
	if len(split) != 3 {
		return
	}
	// 解析商品数量
	goodNums := strings.Split(split[1], "-")
	if len(goodNums) != 6 {
		return
	}
	for i, num := range goodNums {
		if val, err := strconv.Atoi(num); err != nil {
			return
		} else {
			key := fmt.Sprintf("A%d", i+1)
			good := saleSys.GoodMap[key]
			good.Num = val
			saleSys.GoodMap[key] = good
		}
	}

	// 解析钱币面额张数
	moneyNums := strings.Split(split[2], "-")
	if len(moneyNums) != 4 {
		return
	}
	for i, num := range moneyNums {
		if val, err := strconv.Atoi(num); err != nil {
			return
		} else {
			key := 1
			switch i {
			case 0:
				key = 1
			case 1:
				key = 2
			case 2:
				key = 5
			case 3:
				key = 10
			default:
				key = 1
			}
			box := saleSys.MoneyBoxMap[key]
			box.Num = val
			saleSys.MoneyBoxMap[key] = box
		}
	}
	fmt.Println("S001:Initialization is successful")
}

func (saleSys *SaleSys) isAllGoodSoldOut() bool {
	total := 0
	for _, good := range saleSys.GoodMap {
		total += good.Num
	}
	return total == 0
}

func (saleSys *SaleSys) inMoneyBiggerThanOneAndTwo(val int) bool {
	// 计算1元和2元面额总金额
	box1 := saleSys.MoneyBoxMap[1]
	box2 := saleSys.MoneyBoxMap[2]
	t := box1.Denominations*box1.Num + box2.Denominations*box2.Num
	return val > t
}
