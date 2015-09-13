package main

import (
	"net/http"
	"fmt"
	"html/template"
	"time"
	"os"
	"path/filepath"
	"strconv"
	"io"
)

type MyMux struct  {
}

//定义自己的路由就需要实现ServeHTTP方法
func (myMux *MyMux) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//处理逻辑都在request.URL.Path定义的路径里
	switch request.URL.Path {
	case "/":
		doSomething(response, request)
		return
	case "/download":
		downloadFile(response, request)
		return
	case "/upload":
		uploadFile(response, request)
		return
	case "/register":
		registerForm(response, request)
		return
	default:
		http.NotFound(response,request)
		return
	}
}

func downloadFile(resw http.ResponseWriter,req *http.Request){
	http.StripPrefix("/download",http.FileServer(http.Dir("./upload"))).ServeHTTP(resw, req)
}

func uploadFile(resw http.ResponseWriter,req *http.Request){
	if req.Method == "GET" {
		t,_ := template.ParseFiles("web/upload.gtpl")
		t.Execute(resw, nil)
	}else {
		req.ParseMultipartForm(32 << 20)
		file, handler, err :=  req.FormFile("uploadFile")
		if err != nil {
			fmt.Fprintf(resw, "%v", "上传有问题!")
			return
		}
		//后缀,上传类型
		fileext := filepath.Ext(handler.Filename)
		if check(fileext) == false {
			fmt.Fprintf(resw, "%v", "不允许的上传类型")
			return
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + fileext
		fileName = handler.Filename
		//最好是0666的权限，否则windows环境可能会有问题
		f,_ := os.OpenFile("./upload/" + fileName, os.O_CREATE|os.O_WRONLY,0666)
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(resw, "%v", "上传失败!")
			return
		}
		fileDir,_ := filepath.Abs("./upload/" + fileName)
		fmt.Fprintf(resw, "%v", fileName + "上传成功，地址为:" + fileDir)
	}
}

func registerForm(resw http.ResponseWriter,req *http.Request){
	if req.Method == "GET" {
		t,_ := template.ParseFiles("web/register.gtpl")
		t.Execute(resw, nil)
	}else {
		fmt.Fprintf(resw, "注册成功:", req.Form["username"])
	}
}

func doSomething(resw http.ResponseWriter,req *http.Request) {
	fmt.Fprintf(resw, "This is my router.")
}

func check(name string) bool {
	ext := []string{".exe", ".js", ".png"}
	for _,v := range ext {
		if v == name {
			return false
		}
	}
	return true
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9000",mux)
}