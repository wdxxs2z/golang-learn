package main

import (
	"sync"
	"fmt"
)

type Person struct  {
	Name string
	Age int
	Sex string
}

var person *Person

var personMu sync.Mutex

var personOnce sync.Once

func GetPerson_1() *Person {
	if person == nil {
		person = &Person{"No safe.",10,"famle"}
	}
	return person
}

//用法不同once mutex做为本身对象的子对象来使用
func GetPerson_2() *Person {
	personMu.Lock()
	defer personMu.Unlock()
	if person == nil {
		person = &Person{"Safe but not highLevel.",12,"xxx"}
	}
	return person
}

//单例的使用，提供sync.once
func GetPerson_3() *Person {
	personOnce.Do(func(){
		person = &Person{"Safe and use once.",14, "yyy"}
	})
	return person
}

func main(){
	fmt.Println(GetPerson_2())
	fmt.Println(GetPerson_1())
	fmt.Println(GetPerson_3())
}