package main

import (
	"fmt"
)

type USB interface {
	Name() string
	//嵌入式的接口
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string{
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func Disconnect(usb interface{}) {
	//fmt.Println("Disconnected")
	//判断usb的类型
//	if pc, ok := usb.(PhoneConnecter);ok {
//		fmt.Println("Disconnected:", pc.Name())
//		return
//	}
//	fmt.Println("Unknown device.")

	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

func main() {
	//申明一下USB接口
	var a USB
	//初始化接口
	a = PhoneConnecter{"My MeiZu Phone"}
	a.Connect()

	Disconnect(a)
}
