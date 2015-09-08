package main

import(
	"fmt"
	"reflect"
)
//参数 大写表示public 小写表示private，private是不允许被反射的
type User struct {
	Name string
	Id   int
	Age  int
}

type Manager struct {
	//匿名字段
	User
	title string
}

func (u User) SayHello(say string) string{
	return u.Name + " say:" + say
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Printf("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name,f.Type,val)
	}

	for i :=0;i< t.NumMethod();i++{
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}

func Set(o interface{}){
	u := reflect.ValueOf(o)

	if u.Kind() == reflect.Ptr && !u.Elem().CanSet() {
		return
	}else{
		u = u.Elem()
	}

	if f := u.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}

func main() {
	u := User{"byebye",1,12}
	Info(u)

	m := Manager{u,"123"}
	fmt.Println(m)

	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	//int 类型的slice []int{...} User -> Name
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1}))

	x := 555
	//因为要修改值，所以这里要传地址
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(567)
	fmt.Println(x)

	Set(&u)
	fmt.Println(u)

	user := User{"jack",2,22}
//	hello := SayHello("byby2")
//	fmt.Println(hello)
	v2 := reflect.ValueOf(user)
	mv := v2.MethodByName("SayHello")
	args := []reflect.Value{reflect.ValueOf("joe")}
	s := mv.Call(args)
	fmt.Println(s)
}
