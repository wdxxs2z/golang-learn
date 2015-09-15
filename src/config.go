package main

import (
	"fmt"
	"flag"
)

var configStr string

func init(){
	flag.StringVar(&configStr, "c", "","This is config filepath")
	flag.Parse()
}

func main(){
	if configStr !="" {
		fmt.Println(configStr)
	}else {
		fmt.Println("filepath is not flag.")
	}
}
