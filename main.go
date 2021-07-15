package main

import "fmt"

func main() {
	//fmt.Println("Heelo world")
	minAmount := 5000
	purchase := 15_000
	percentCashback := 15
	limitCashback := 2000
	cashback := 0
	const fullPercent int = 100
	// 1. Можно ли дать кэшбек
	if purchase >= minAmount{
		cashback = purchase * percentCashback / fullPercent
		fmt.Println("Prediction cashback is", cashback)
	}
	// 2. Превҷшение кешбека
	if cashback > limitCashback{
		cashback = limitCashback
	}
	fmt.Println("cashback is", cashback)
}
// git init
// git add .
// TODO: git commit -m "my message"