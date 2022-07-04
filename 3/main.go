package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
  TODO 第三题：rand问题
*/

var (
	Limit     = 5
	KnowLimit = 14
)

func main() {
	fmt.Printf("num=%d", randNum())
}

func randNum() int {

	flag := true
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(KnowLimit)
	for flag {
		if num > Limit {
			num = randNum()
		} else {
			flag = false
		}
	}
	return num
}
