package test

import "fmt"

// SayHI 打印 "Hello world!"
func SayHI() {
	fmt.Println("Hello world!")
}

var name string = "Kai"

func SayName() {
	fmt.Println("Hi! " + name)
}

func PrintVar() {
	var variableName int = 1
	variableName = variableName + 1
	fmt.Println(variableName)
}

func NineToNine() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func Show(msg string) {
	fmt.Println(msg)
}

func Multipy(n1 int, n2 int) {
	var result int = n1 * n2
	fmt.Println(result)
}

func ShowHello(msg string) string {
	if msg == "Hello" {
		return "Good"
	}
	return msg

}

type Day int

const (
	SunDay Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// iota 是 Go 語言專為常量列舉而設計的關鍵字，它在簡單的數值枚舉情況特別好用。當你只需要定義一組整數枚舉時，iota 可以讓程式碼更簡潔，並自動生成遞增的整數值。
// 所以iota 只能用在數字的場合 ,它預設開始值是 0 ,每新增一行就+1

func (d Day) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}

// Currency 枚舉的結構體
type Currency struct {
	Code        string // 貨幣代碼
	Description string // 描述
}

var (
	VND = Currency{"vnd", "越南盾"}
	INR = Currency{"inr", "印度卢比"}
	BRL = Currency{"brl", "巴西里奥"}
	THB = Currency{"thb", "泰国泰铢"}
)

func PrintArray() {

	var arr [10]int
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	arr[5] = 5

	fmt.Println(arr)
}

func PrintSlice() {
	// slice 為動態陣列 不需限定陣列長度
	slice := []int{1, 2, 3}
	// := 是短變數宣告（Short Variable Declaration）的語法糖
	fmt.Println(slice)
}
