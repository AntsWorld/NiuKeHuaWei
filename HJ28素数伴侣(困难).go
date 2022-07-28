package main

// 数据范围： 1 ≤ n ≤ 100 ,输入的数据大小满足 2 ≤ val ≤ 30000
func main() {

}

// 判断一个数是否为素数
// 质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数；否则称为合数（规定1既不是质数也不是合数）
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
