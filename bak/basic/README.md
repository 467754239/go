## 序
```
安装 done
类型 - 常量 - 变量 - 运算符 done
控制流 done
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
1.[安装说明](https://golang.org/doc/install)

2.[包下载地址](https://code.google.com/p/go/downloads/list)

3.确认是否安装成功
```
$ go version
go version go1.7 linux/amd64
```
4.go命令行

```
$ go help
go build        编译
go clean        移除当前源码包里面的编译生成文件
go fmt          格式化代码
go get          动态获取远程代码包
go install      生成结果文件，并将编译好的结果一到$GOPATH/pkg或者$GOPATH/bin
go test         运行测试用的可执行文件
go doc          godoc -http=:8080 查看文档

go fix          修复以前老版本代码到新版本
go version      查看当前版本
go env          查看当前go的环境变量
go list         列出当前所有安装package
go run          编译并运行go语言程序
```
5.整体目录结构
```
通过package组织，只有package名称为main的可以包含main函数
一个程序有且有一个main包
通过import关键字导入其它非main包
```
6.注释
```
单行注释 //
多行注释 /* ..... */
```
7.编辑器
```
sublime
vim
```

## 基本结构及最小示例
1.基本结构
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
2.最小示例
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

## 类型
1.类型默认值
```
声明不赋值，类型为零值、非空值、而是声明后的默认值.
bool: false
integers: 0
float32: 0
string: ""
```
2.保留字
```
break      case   chan     const        continue
default    defer  else     fallthrough  for
func       go     goto     if           import
interface  map    package  range        return
select     struct switch   type         var
```

## 变量
1.变量声明
```
// 第一种，指定变量类型，声明后如不赋值则使用默认值.
var v_name v_type
v_name = value

// 第二种，根据值自动判断变量的类型
var v_name = value

// 第三种，省略var关键字用冒号:代替
// := 左侧的变量不应该是已经声明过的，否则会导致编译错误.
v_name := value

#e.g.
var a int = 10
var b = 10
b := 10
```

示例
```
package main

import (
    "fmt"
)

var a = 1234
var b string = "hello"
var c bool

func main() {
    fmt.Println(a, b, c)
}
```

结果
```
1234 hello false
```


2.多变量声明
```
// 类型相同多个变量、非全局变量
func main() {

    var v_name1, v_name2, v_name3 type
    v_name1, v_name2, v_name3 = value1, value2, value3
    or
    /* 出现在:=左侧的变量不应该是已经被声明过的，否则会导致编译错误 */
    v_name1, v_name2, v_name3 := value1, value2, value3   

}

// 类型不同多个变量、全局变量、局部变量不能使用这种方式
var (
    v_name1 v_type1
    v_name2 v_type2
)
```

示例
```
package main

import (
    "fmt"
)

var x, y int
var (
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
    g, h := 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
```
结果
```
0 0 0 false 1 2 123 hello 123 hello
```

注意
```
1. 多变量赋值时，将先行计算所有左侧变量的值，再进行赋值.
......

2. 垃圾桶
func test(int, string) {
    return 123, "abc"
}

a, _ := test()

3. 已声明但是没有使用的变量会在编译阶段报错，较python更严格.
```

## 常量
```
常量可以是字符、字符串、布尔或数字
常量赋值时编译期的行为
```

## 控制语句
> if | for | switch | goto | break | continue

IF  
1.说明
```
条件表达式没有括号
支持一个初始化表达式（可以是多变量初始化语句）
初始化语句中定义的都是只能在block级别中使用的局部变量，不能在block之外使用
左大括号必须和条件语句在同一行(必须与if/else在同一行)
go没有三元运算符

if判断语句条件不需要括号
在判断语句里允许声明一个变量，其变量的作用域只在逻辑块内，其它地方不起作用
花括号一定存在，且必须与if/else在同一行
```

2.语法
```
//基本
if a > 0 {  //无括号
     dosomething()
} else if a == 0 { //必须用花括号
     doothertings()
} else {
     donothing()
}

//单行模式
if a > 0 { a += 100 } else { a -= 100 }
```

3.示例
```
package main

import (
    "fmt"
)

func main() {
    // 方式1
    a := 10

    if a > 0 {
        a += 100
    } else if a == 0 {
        a = 0
    } else {
        a -= 100
    }

    fmt.Println(a)

    // 方式2
    b := 101
    if b > 0 {
        b += 100
    } else {
        b -= 100
    }
    fmt.Println(b)

    // 方式3
    //支持一个初始化参数，变量c的作用域尽在if/else范围内.
    if c := 1; c < 10 {
        c += 1
        fmt.Println("c =", c)
    }

    // 方式4
    if a, b := 1, 2; a+b == 10 {
        fmt.Println("into if")
    } else {
        fmt.Println("into else")
    }

    // 方式5
    if x := computedValue(); x > 10 {
        ...
    } else {
        ...
    }

}
```

> 注意

```
/*
在有返回值的函数中，不允许将“最终的”return语句放到if ... else ... 结构中，否则编译失败
*/
func example(x int) int {
    if x == 0 {
        return 5
    } else {
        return x
    }
}
```

FOR
> go中没有"while"，但支持三种形式的for来实现各种功能.


1.说明
```
- Go只有for一个循环语句关键字，但支持3中形式
- 初始化和步进表达式可以是多个值
- 条件语句每次循环都会被重新检查，因此不建议在条件语句中使用函数，
尽量提前计算好条件并以变量或常量代替
- 左大括号必须和条件语句在同一行

- 初始化和步进表达式可以使多个值：必须使用平行赋值
i, j:=0,len(a)-1; i<j; i,j=i+1,j-1
- 每次循环都会重新检查条件表达式
```

2.语法
```
for init; condition; post {
    //init不支持逗号，只能平行赋值
    //condition每次循环开始都会检查，不建议在里面使用函数，建议用计算好的变量/常量代替
    //post后面必须跟花括号,每轮循环结束的时候调用
}

for i:=0; i<10; i++ {
}

-----------------------------

for condition {
    dosomething()
}

i:=1
for i<10 {
}

-----------------------------

for { //无限循环
    dosomething()
}

for {
    a++
    if a > 10 {
        break
    }
}

-----------------------------
```

3.三种形式示例
```
package main

import (
    "fmt"
)

func main() {

    // 第一种形式
    a := 1
    for {
        a++
        if a > 3 {
            break
        }
        fmt.Println(a)
    }
    fmt.Println("a finish")

    // 第二种形式
    b := 1
    for b < 3 {
        b++
        fmt.Println(b)
    }
    fmt.Println("b finish")

    // 第三种形式
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
    fmt.Println("i finish")
}

```

选择语句SWITCH

1. 说明
```
- 可以使用任何类型或表达式作为条件语句
- 不需要写break，一旦条件符合自动终止
- 如希望继续执行下一个case，需要使用fallthrough语句
- 支持一个初始化表达式(可以是并行方式)，右侧需跟分号
- 左大括号必须和条件语句在同一行
```

2.语法
```
switch a:=5; a { //默认break，匹配成功后不会自动向下执行其他case,而是跳出整个switch
    case 0:         //普通
        println(a)
    case 1, 2:      //多个分支，逗号分隔
        println(a)
    case 100:       //什么都不做
    case 5:
        println(a)
        fallthrough   //进入后面的case 进行处理，而不是跳出block
    default:
        println(a)    //默认
}
注意，不需要break来明确退出一个case，一旦条件符合，自动终止,除非使用fallthough

--------------

可以不带表达式
switch sExpr {
        case expr1: //sExpr和expr1必须类型一致,不限制为常量或者证书，可以用任何类型或表达式
            ...
}

switch {
    case 0 <= Num && Num <= 3:
        fmt.Printf("3")
}
```

3.三种形式示例
```
package main

import (
    "fmt"
)

func main() {

    // 第一种形式
    a := 1
    switch a {
    case 0:
        fmt.Println("a=0")
    case 1:
        fmt.Println("a=1")
    default:
        fmt.Println("None")
    }

    // 第二种形式
    ab := 1
    switch {
    case ab >= 0:
        fmt.Println("ab=0")
        fallthrough // 不终止，继续检查下一个case.
    case ab >= 1:
        fmt.Println("ab=1")
    default:
        fmt.Println("None")
    }

    // 第三种形式
    switch ab := 1; {
    case ab >= 0:
        fmt.Println("ab=0")
        fallthrough // 不终止，继续检查下一个case.
    case ab >= 1:
        fmt.Println("ab=1")
    default:
        fmt.Println("None")
    }

}
```

goto | break | continue

1.说明
```
- 三个语法都可以配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- break与continue配合标签可用于多层循环跳出
- goto是调整执行位置，与其它2个语句配合标签的结果并不相同

break/continue可配合标签用于多重循环跳出
goto是调整执行位置，与其他两个执行结果并不相同
```

2.二种形式示例
```
package main

import (
    "fmt"
)

func main() {
LABEL1:
    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                fmt.Println(i)
                break LABEL1
            }
        }
    }
    fmt.Println("ok")

    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                fmt.Println(i)
                goto LABEL2
            }
        }
    }
LABEL2:
    fmt.Println("ok")

}
```

## 数组Array
1.说明
```
- 定义数组的格式： var <varName> [n] <type>, n>=0
- 数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组的指针和指针数组
- 数组在Go中为值类型
- 数组之间可以使用==或!=进行比较，但不可以使用<或>
- 可以使用new来创建数组，此方法返回一个指向数组的指针
- Go支持多维数组

在Go语言中数组并不是一个统一的数据类型，将数组的长度也作为类型的一部分.
var a [2]int // 长度为2的int型数组
var b [1]int // 长度为1的int型数组
这两个数组是不同类型的，a和b是不同直接==赋值的，必须用循环来完成赋值的操作.
```

## 切片Slice
```
- 其本身并不是数组，它指向底层的数组
- 作为变长数组的替代方案，可以关联底层数组的局部 或 全部
- 为引用类型
- 可以直接创建 或 从底层数组获取生成
- 使用len()获取元素个数，cap()获取容器
- 一般使用make()创建
- 如果多个slice指向相同的底层数组，其中一个的值改变会影响全部 
- make([]T, len, cap)
- 其中cap可以省略，则和len的值相同
- len表示存数的元素个数，cap表示容量
```