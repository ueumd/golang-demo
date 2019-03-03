package main

import (
	"fmt"
	"math/big"
	cryptoRand "crypto/rand"
	mathRand "math/rand"
	"time"
)

func mathRandDemo()  {
	mathRand.Seed(int64(time.Now().UnixNano()))
	fmt.Println(mathRand.Int())
}

// GenerateRandnum 生成最大范围内随机数
func GenerateRandnum() int {
	mathRand.Seed(time.Now().Unix())
	randNum := mathRand.Intn(100)

	fmt.Printf("rand is %v\n", randNum)

	return randNum
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	mathRand.Seed(time.Now().Unix())
	randNum := mathRand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func cryptoRandDemo()  {
	// 生成 20 个 [0, 100) 范围的真随机数。
	for i := 0; i < 20; i++ {
		result, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(100))
		fmt.Println(result)
	}
}

func cryptoRandDemo2()  {
	//cryptoRand.Int(cryptoRand.Reader, int64(time.Now().UnixNano()))
}

func main() {
	mathRandDemo()
	GenerateRandnum()
	GenerateRangeNum(1, 5)
}