![golang](img/golang.png)

## 序
```
安装 done
类型 - 常量 - 变量 - 运算符
控制流
数据结构：数组、切片、map
函数
结构struct
接口
字符串
文件
异常处理
数学计算
并发
反射
正则
数据库
时间日期
配置文件
日志logging
json/xml文件处理
测试
项目及项目结构
```

## 简介
```
## Go语言最主要的特性
自动垃圾回收
更丰富的内置类型
函数多返回值
错误处理
匿名函数和必包
类型和接口
并发编程
反射
语言交互性
高性能/高效开发
```

## 安装
[安装说明](https://golang.org/doc/install)

[包下载地址](https://code.google.com/p/go/downloads/list)

```
1. 确认是否安装成功
$ go version
go version go1.7 linux/amd64

2. go命令行
$ go help

go build 编译
go clean 移除当前源码包里面的编译生成文件
go fmt 格式化代码
go get 动态获取远程代码包
go install 生成结果文件，并将编译好的结果一到$GOPATH/pkg或者$GOPATH/bin
go test 运行测试用的可执行文件
go doc   godoc -http=:8080 查看文档

go fix 修复以前老版本代码到新版本
go version查看当前版本
go env 查看当前go的环境变量
go list 列出当前所有安装package
go run 编译并运行go语言程序

3. 整体目录结构
通过package组织，只有package名称为main的可以包含main函数
一个程序有且有一个main包
通过import关键字导入其它非main包

4. 注释
单行注释 //
多行注释 /* ..... */

5. 编辑器
sublime
```

## 最小示例
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

## 基本结构
>basic_structure.go

```
// 当前程序的包名
package main

// 导入其它的包
import (
    "fmt"
)

// 变量的定义
const PI = 3.14

// 全局变量的声明与赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
    fmt.Println("Hello world!你好，世界! ")
}

```

## 类型、变量、常量
## 控制语句
## array
## slice
## map
## function
## struct
## method
## interface
## reflection
## concurrency
