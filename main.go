package main

import (
	"GO-demo/test"
	"fmt"
)

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

	test.SayHI()
	test.SayName()
	test.PrintVar()
	test.NineToNine()
	test.Show("FUCK")
	test.Multipy(-2, 6)
	test.Multipy(27, 9)
	fmt.Println(test.ShowHello("Hello"))
	fmt.Println(test.ShowHello("Hi"))

	var today test.Day = test.Monday
	fmt.Println(today) // 輸出：5

	// 也可以使用條件檢查
	if today == test.Friday {
		fmt.Println("It's Friday!")
	} else {
		fmt.Println("False")
	}

	today = test.Friday
	fmt.Println(today)

	var currency test.Currency = test.THB
	var currencyCode string = test.BRL.Code
	var currencyDescription string = test.BRL.Description
	// Code、Description 為string 所以要宣告為string
	fmt.Printf("貨幣代碼: %s, 描述: %s\n", test.VND.Code, test.VND.Description)
	fmt.Println(currency) // 會印出整個enum
	fmt.Println(currencyCode)
	fmt.Println(currencyDescription)

	test.PrintArray()

	test.PrintSlice()
}
