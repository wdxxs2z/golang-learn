package main

import (
	"fmt"
)

type TZ int

type T struct {
	name string
	sex  string
	age  TZ
}

func (tz *TZ) increaseNumber(num int){
	*tz += TZ(num)
}

func NewT(name string, sex string, age TZ) *T {
	return &T{
		name,
		sex,
		age,
	}
}

func main() {
	//注意，这里是别名的形式 所以不能声明为 ：= 要
	var a TZ
	a.increaseNumber(100)
	fmt.Println(a)

	b := T{"huahua","famale",10}
	fmt.Println(b.age)
}