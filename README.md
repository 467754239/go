
## Golang

命令行帮助
```
// 查看fmt包
$ go doc fmt

// 查看单个函数Printf
$ godoc fmt Printf

// .代表本地目录
// ...代表从当前目录递归往下走，找到没有下载的包全部下载
$ go get ./...
```

生成本地Golang官方文档
```
$ godoc -http=:8080
```

注意事项
```
- 单引号和双引号区别
字符串一定要用双引号引起来，用单引号会出现异常。
单个字符可以用单引号引起来，表示这个字符的ascii码。
```

数组
```
# 声明与赋值
- var array_name [length]type
var person [3]string
person[0] = "zhengyscn"
person[1] = "china"
person[2] = "beijing"
fmt.Println("person: ", person)

## 切片 slice
array_name[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。
因此
s[lo:lo] 是空的
s[lo:lo+1] 有一个元素
```

map
```
# 声明
map 映射键到值
map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。
```



### 参考链接

- [Golang官网](https://golang.org/)
- [play.golang 翻墙](https://play.golang.org/)
- [Go Search](http://go-search.org/)
- [Go by Example 中文](http://gobyexample.everyx.in/)  
- [the-little-go-book_ZH_CN](https://github.com/songleo/the-little-go-book_ZH_CN)  
- [awesome-go](https://github.com/avelino/awesome-go)
- [golang-tour-英文版](https://tour.golang.org/welcome/1)  
- [github-golang-tour](https://github.com/golang/tour/tree/master/solutions)  
- [github-golang-tour](https://github.com/fengsp/golang-tour)
- [Go入门指南](http://www.kancloud.cn/kancloud/the-way-to-go/72432)