// go run 直接执行
// go build 编译可执行程序
package main
//main包
//main是程序包，是程序的入口

import "fmt"
//main函数
//必须调用一个函数
//调用的函数必须使用

func main() {
	var a, b int //初始a，b
	fmt.Println("helloworld") //helloworld
	a = 1 //赋值a
	b = 2 //赋值b
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	a = 3
	fmt.Println("a =", a)
	c := 30 //c必须是一个没有初始化过的值，如果在前面初始化过，就会报错 no new variables on left side of :=
	fmt.Printf ("now c type is %T\n", c) // %T\n表示显示这个变量的属性
//最终打印以下内容
/*
	helloworld
	a = 1
	b = 2
	a = 3
	now c type is int
*/
}

//教程来自哔哩哔哩
//视频链接//www.bilibili.com/video/av20432910
// to learn 10P 自动推导类型和赋值区别

/*
	git入门
	git clone 下载某库
	git add . 添加更改
	git commit -m "commit名" 添加commit
	git push -u origin <branch> 推送更新
	git pull 同步云端更新
*/