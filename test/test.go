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

func (d Day) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[d]
}
