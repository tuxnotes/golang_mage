# 1 基础介绍

https://github.com/golang/go/wiki/Projects
https://go.dev/

## 1.1 Golang特点

易于开发，快速编译，高效执行

- 执行速度快，编译速度慢(编译型)：C，C++
- 执行速度较慢，编译速度快(解释型)：JAVA .NET
- 执行速度慢，开发难度小(动态脚本)：Python，PHP

特性

- 静态类型并具有丰富的内置类型
- 函数多返回值
- 错误处理机制
- 语言层面支持并发
- 面向对象：使用类型、组合、接口来实现对象对象思想
- 反射
- CGO：用于调用C语言实现的模块
- 自动垃圾回收
- 静态编译
- 交叉编译
- 易于部署
- 基于BSD协议完全开放

## 1.2 Golang应用场景

主要用于服务端开发，其定位是开发大型软件，常用于下面场景：
- 服务器编程：日志处理、虚拟机处理、文件系统、分布式系统等
- 网络编程：web应用、API应用、下载应用等
- 内存数据库
- 云平台
- 机器学习
- 区块链

## 1.3 如何学习

1. 知识点分解
2. 刻意练习(练习不会的、不懂或难以理解的知识点)
3. 反馈
   - 主动反馈：看别人代码
   - 被动反馈：codereview

总结：多写(禁止复制+粘贴)，多看，多问

# 2 Golang基础

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world!!!")
}
```
编译
```bash
go build -x main.go

WORK=C:\Users\aspire\AppData\Local\Temp\go-build1583018854
mkdir -p $WORK\b001\
cat >$WORK\b001\importcfg.link << 'EOF' # internal
packagefile command-line-arguments=C:\Users\aspire\AppData\Local\go-build\20\20205aa8ed3afe25f785036ab1b5f007ea47d351d18513e50598ab1596fd0360-d
packagefile fmt=C:\Go\pkg\windows_amd64\fmt.a
packagefile runtime=C:\Go\pkg\windows_amd64\runtime.a
packagefile errors=C:\Go\pkg\windows_amd64\errors.a
packagefile internal/fmtsort=C:\Go\pkg\windows_amd64\internal\fmtsort.a
packagefile io=C:\Go\pkg\windows_amd64\io.a
packagefile math=C:\Go\pkg\windows_amd64\math.a
packagefile os=C:\Go\pkg\windows_amd64\os.a
packagefile reflect=C:\Go\pkg\windows_amd64\reflect.a
packagefile strconv=C:\Go\pkg\windows_amd64\strconv.a
packagefile sync=C:\Go\pkg\windows_amd64\sync.a
packagefile unicode/utf8=C:\Go\pkg\windows_amd64\unicode\utf8.a
packagefile internal/bytealg=C:\Go\pkg\windows_amd64\internal\bytealg.a
packagefile internal/cpu=C:\Go\pkg\windows_amd64\internal\cpu.a
packagefile runtime/internal/atomic=C:\Go\pkg\windows_amd64\runtime\internal\atomic.a
packagefile runtime/internal/math=C:\Go\pkg\windows_amd64\runtime\internal\math.a
packagefile runtime/internal/sys=C:\Go\pkg\windows_amd64\runtime\internal\sys.a
packagefile internal/reflectlite=C:\Go\pkg\windows_amd64\internal\reflectlite.a
packagefile sort=C:\Go\pkg\windows_amd64\sort.a
packagefile math/bits=C:\Go\pkg\windows_amd64\math\bits.a
packagefile internal/oserror=C:\Go\pkg\windows_amd64\internal\oserror.a
packagefile internal/poll=C:\Go\pkg\windows_amd64\internal\poll.a
packagefile internal/syscall/execenv=C:\Go\pkg\windows_amd64\internal\syscall\execenv.a
packagefile internal/syscall/windows=C:\Go\pkg\windows_amd64\internal\syscall\windows.a
packagefile internal/testlog=C:\Go\pkg\windows_amd64\internal\testlog.a
packagefile io/fs=C:\Go\pkg\windows_amd64\io\fs.a
packagefile sync/atomic=C:\Go\pkg\windows_amd64\sync\atomic.a
packagefile syscall=C:\Go\pkg\windows_amd64\syscall.a
packagefile time=C:\Go\pkg\windows_amd64\time.a
packagefile unicode/utf16=C:\Go\pkg\windows_amd64\unicode\utf16.a
packagefile internal/unsafeheader=C:\Go\pkg\windows_amd64\internal\unsafeheader.a
packagefile unicode=C:\Go\pkg\windows_amd64\unicode.a
packagefile internal/race=C:\Go\pkg\windows_amd64\internal\race.a
packagefile internal/syscall/windows/sysdll=C:\Go\pkg\windows_amd64\internal\syscall\windows\sysdll.a
packagefile path=C:\Go\pkg\windows_amd64\path.a
packagefile internal/syscall/windows/registry=C:\Go\pkg\windows_amd64\internal\syscall\windows\registry.a
EOF
mkdir -p $WORK\b001\exe\
cd .
"C:\\Go\\pkg\\tool\\windows_amd64\\link.exe" -o "$WORK\\b001\\exe\\a.out.exe" -importcfg "$WORK\\b001\\importcfg.link" -buildmode=pie 
-buildid=V3fXC9GiYRn10G5LT5Id/DzKLvAZS7bgXbvtGDS5w/LhSQ0CXyl30YEZgMTRIM/V3fXC9GiYRn10G5LT5Id -extld=gcc "C:\\Users\\aspire\\AppData\\Local\\go-build\\20\\20205aa8ed3afe25f785036ab1b5f007ea47d351d18513e50598ab1596fd0360-d"
"C:\\Go\\pkg\\tool\\windows_amd64\\buildid.exe" -w "$WORK\\b001\\exe\\a.out.exe" # internal
cp $WORK\b001\exe\a.out.exe main.exe
rm -r $WORK\b001\
```
重命名生成的执行文件
```bash
go build -x -o hello.exe main.go
```
## 2.1 基本组成元素
### 标识符
- 标识符是编程是使用的名字，用于给变量、常量、函数、类型、接口、包名等进行命名，以建立名称和使用之间的关系
- Go语言标识符的命名规则：
  1. 只能由非空字母(Unicode)、数字、下划线组成
  2. 只能以字母或下划线开始
  3. 不能使用Go关键字
  4. 避免使用Go语言预定义标识符
  5. 建议使用驼峰式：myName
  6. 标识符区分大小写

Go语言提供了一些预先定义的标识符用来表示内置的常量、类型、函数
在自定义标识符是应避免使用：
1. 内置常量：true false nil iota
2. 内置类型：bool byte rune int int8 int16 int32 int64 uint8 uint16 uint32 uint64 uintptr float32 float64 complex64 complex128 string error
3. 内置函数：make len cap new append copy close delete complex real imag panic recover
4. 空白标识符： _
### 关键字
关键字用于特定的语法结构，Go语言定义25个关键字：
- 声明：import package
- 实体声明和定义：char const func interface map struct type var
- 流程控制：break case continue default defer else fallthrough for go goto if range return select switch

### 字面量
字面量是值的表示方法，常用于对变量/常量进行初始化
主要分为：
- 标识基础数据类型值的字面量，例如：0 1.1 true 3+4i 'a' "我爱中国"
- 构造自定义的复合数据类型的类型字面量，例如：type Interval int
- 用于表示复合数据类型值的复合字面量，用来构造array slice map struct的值，例如:{1, 2, 3}

### 操作符
- 算数运算符：+ - * / % ++ --
- 关系运算符：> >= < <= == !=
- 逻辑运算符：&& || !
- 位运算符：& | ^ << >> &^
- 赋值运算符：= += -= *= /= %= &= |= ^= <<= >>=
- 其他运算符：&(单目)、*(单目)、.(单目)、-(单目)、...、<-

### 分隔符
- 小括号(),中括号[]，大括号{},分号;，逗号,


## 2.2 变量

### 声明
声明语句用于定义程序的各种实体对象，主要有：
- 声明变量的var
- 声明变量的const
- 声明函数的func
- 声明类型的type

### 变量

变量是指对一块存储空间定义名称，通过名称对存储空间的内容进行访问或修改。变量声明，常用的语法为：
1. 定义变量并进行初始化：var name string = "silence"
2. 定义变量使用零值进行初始化：var age int
3. 定义变量，变量类型通过值类型进行推导：var isBody = true
4. 定义多个相同类型的变量并使用零值进行初始化：var prefix, suffix
5. 定义多个相同类型的变量并使用对应的值进行初始化：var prev, next int = 3, 4
6. 定义多个变量并使用对应的值进行初始化，变量的类型使用值类型进行推导，类型可不相同：var name, age = "slicen", 30

## 2.3 常量

常量是值不可更改的。不能赋值。定义多个变量并进行初始化，批量赋值中变量类型可以省略，并且除了第一个常量值外其他常量可同时省略类型和值，表示使用前一个常量的初始化表达式。



- 省略类型：const PI = 3.14
- 定义多个常量：相同类型,const C1, C2 int = 1,2
- 定义多个常量：不同类型
  ```go
  const (
    age int = 20
    name string = "foo"
  )
  ```
- 定义多个常量：省略类型,const C3,C4 = "abc",1

常量一般大写，且必须有默认值

常量的另一种写法：
```go
const (
  C5 int = 1
  C6
  C7
)
```
这种场景用于枚举类型，go语言没有枚举类型，使用const + iota实现

```go
const (
  E1 int = iota
  E2
  E3
)
```

## 2.4 作用域

作用域：定义标识符可以使用的范围。在go中使用{}大括号来定义作用域的范围，大括号包含一连串的语句，叫做语句块。语句块可以嵌套，语句块内定义的变量不能在语句块外使用。

常见隐式语句块：
- 全语句块
- 包语句块
- 文件语句块
- if switch for select case语句块

在不同语句块中变量只能定义一次。变量在语句块中一定要使用

作用域内定义变量只能被声明一次且变量必须使用，否则编译错误。在不同作用域可定义相同的变量，此时局部将覆盖全局

## 2.5 注释

Go支持两种注释方式，行注释和块注释：
- 行注释：以//开头，例如：// 我是行注释
- 块注释：以/*开头，  以 */结尾，例如：/*我是块注释*/

## 2.5 基本类型

### 2.5.1 布尔类型
用于表示真假，类型名称bool,**只有两个值true与false**。占用一个字节。使用fmt.Printf格式化输出的时候使用%t作为占位符
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var zero bool
	isBoby := true
	fmt.Printf("Type of isBoby: %T, bytes take for bool type: %d\n", isBoby, unsafe.Sizeof(isBoby))
	fmt.Println(zero)
	fmt.Printf("Default zero value for bool type: %t\n", zero)
}
```
#### 逻辑运算
运算符两边均为布尔类型
与： && 。两边均为true
或： || 。
非： !

#### 关系运算
==相等 !=不相等
关系运算的结果是bool类型
### 2.5.2 数值类型
#### 整型
Go语音提供了5种有符号，5种无符号，1种指针，1种单字节，1种单个unicode字符(unicode code point)，共13种整形类型，零值均为0。

int uint rune(unicode码点,int32) int8 int16 int32 int64 uint8 uint16 uint32 uint64 byte(unint8) unintptr

标识符： int/int*/uint/unint*/unintptr/rune/byte
字面量：十进制，八进制0777=7*8^2 + 7 * 8^1 + 7 * 8^0，十六进制0x10

```go
package main

import (
  "fmt"
)

func main() {
  var age int = 30
  fmt.Printf("%T %d\n", age, age)
  fmt.Println(0777,0x10)
  var a byte = 'A'
  var m rune = '中'
  fmt.Println(a,m)
  fmt.Printf("%T\n",a)
  fmt.Printf("%T\n",m)
}
```
常见操作：算数运算

##### 位运算

& | ^ << >> ^&位清空
十进制转二进制：辗转相除法

##### 类型转换
从大往小转可能发生溢出
#### 浮点型
```go
package main

import "fmt"

func main() {
  var hight float64 = 1.68
  fmt.Printf("%T %f\n",hight,hight)
  var weight float64 = 13.05E1
  fmt.Printf("%T %f\n", weight, weight)
}
```
字面量
十进制表示法
科学计数法
### 2.5.3 字符串类型
字面量，有两种方法表示：一种是使用双引号""，一种是使用反引号``
双引号是可解释的字符串，反引号是原生字符串，特殊字符如\n在原生字符串中不会转义
算数运算：+
关系运算： "bb" > "ba" true 从第一个不相等的比较
赋值运算： ：= +=
如果通过索引(切片)的方式定位字符串的内容，则字符串定义内容只能位ASCII，否则会出现乱码
获取字符串的长度，使用len函数，但字符串内容必须只能是ASCII码。因为Unicode在存储的时候一个字符是多个字节
### 2.5.4 枚举类型
const + iota实现
### 2.5.5 指针类型

```go
package main

import "fmt"

func main() {
	a := 2

	var p *int = &a
	*p = 3
	fmt.Printf("%T %p %d\n", p, p, a)
}
```
获取用户输入数据:使用fmt.Scan(),接受指针,不带缓冲的IO
```go
package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入名字：")
	fmt.Scan(&name) // scan receive pointer
	fmt.Printf("你的名字是: %s\n", name)
}
```
## 2.6 流程控制

### 2.6.1 条件选择
#### if else

#### switch
switch不需要操作数，它仅仅是列出case，每个case是一个布尔表达式
```go
package main

import "fmt"

func conflip() string {
	return "heads"
}

func Signum(x int) int {
	switch { // tagless switch, which is equivalent to switch true
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func main() {
	var heads, tails int
	switch conflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
}
```

### 2.6.2 循环
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		/*
			每次调用input.Scan(),就会读取下一行，并将行尾的换行符去除，使用input.Text()获取内容.
			如果没有可读取的行了，则Scan()函数返回false
		*/
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Dup line: %s count: %d\n", line, n)
		}
	}
}
```
### 作业

乘法表
```go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d ", j, i)
		}
		fmt.Println()
	}
}
```
## 2.7 复合数据类型
复合数据类型可以理解为可以存储多个单值，类似于集合
### 2.7.1 数组
固定长度，元素类型相同，数组长度也是类型的一部分，长度是一个常亮，如字面量。声明数组的时候必须有数组长度和元素类型。
```go
package main

import "fmt"

func main() {
	var numbers [10]int
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
	var s1 [3]string
	fmt.Printf("%q\n", s1)
}
```
#### 赋值
字面量的方式赋值
```go
var nums [10]int{10, 20, 30}
```
索引的方式
```go
var nums [10]int{0:10, 9:30}
```
数组长度根据数组长度进行推导
```go
[...]int{1,2,3} // 数组长度为3
```
#### 获取数组长度
```go
fmt.Println(len(numbers))
```
#### 数组索引
索引范围0-(len(array)-1)
numbers[0] = 1000

使用短声明给多个变量赋值时，:=左侧至少有一个新的变量
```go
var value int
_, value, xx := 1,2,3 // xx为新的变量
fmt.Println(value,xx)
```
### 2.7.2 切片
切片是长度可变的数组(由数据类型相同的元素组成的一组长度可变的序列)，由三部分组成：
- 指针：指向通过切片可访问到的底层数组的第一个元素，不必须是低等数组的第一个元素
- 长度：切片元素的数量，不能超过容量
- 容量：切片开始到底层数组结束位置元素的数量
数组切片操作后得到的类型是slice.slice有长度和容量的属性。

#### 声明
切片的声明需要指定元素的类型，但不需要指定存储元素的数量。切片声明后，会被初始化为nil，表示暂不存在的切片
```go
var nums []int
numbers := make([]int, 3)
```
#### 切片操作
切片操作的写法为s[i:j]，其中0 ≤ i ≤ j ≤ cap(s)。此操作会生成一个包含i到j-1的新切片。其中s可以是数组变量，指向数组的指针，或者是其他的切片。
切片操作超出容量将会导致panic，但是可以超过length。
因为切片包含指针，因此将切片传给函数，则允许函数修改底层数组。
与数组不同的是，切片不支持比较。标准库提供了byte.Equal函数，用于比较byte类型的切片，但是其他的类型就得自己实现了。
#### 切片的零值
切片类型的零值是nil,零值的切片表示没有底层数组。nil slice 的length和capacity均为0.但也有length和capacity为0，单slice不为slice的情况：
```go
var s []int    // len(s) == 0, s == nil
s = nil        // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}    // len(s) == 0, s != nil
```
因此要测试slice是否为空，使用len(s) == 0，而不是s == nil
#### 删除元素
使用切片操作
1. 从切片中间位置移除元素，需要保留切片顺序可以使用copy

```go
func remove(slice []int, i int) []int{
	copy(slice[:i],slice[i+1:])
	return slice[:len(slice)-1]
}
```
2. 如果不需要保留顺序，则直接将最后一个元素赋值给要移除的元素的位置

```go
func remove(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
```

#### copy append

### 2.7.3 多维切片
元素也是切片

## 2.8 map映射
映射是存储一系列无序的key/value对，通过key来对value进行增删改查
key/value规则：
map的key只能是可以使用==运算的类型，如字符串，数字，布尔，数组；value可以是任意类型
### 2.8.1 操作
增删改查
如果查找的key在map中不存在，则返回value类型的零值。**但是map的元素不是变量，不能进行取地址操作**。原因之一是map的增长将导致元素重新hash到新的存储位置。
删除：delete内置函数
### 2.8.2 map的零值
map的零值为nil，意味着还没有指向任何的hash表
```go
	var ages map[string]int
	fmt.Println(ages == nil)  // true
	fmt.Println(len(ages) == 0) // true
	ages["carol"] = 21 // panic: assignment to entry in nil map
```
对nil map 的大多数操作如lookup,delete,range,len都是安全的，其行为如同空map，但是往nil的map存储值将引发panic。因此存储前首先需要分配map。
```go
if age, ok := ages["Bob"]; !ok { /*...*/ }
```
### 2.8.3 比较
与slice一样，map不支持比较，唯一可以比较的对象是nil。如果想比较则要手动实现：
```go
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || y[k] != x[k] {
			return false
		}
	}
	return true
}
```
### 2.8.4 基于map实现集合
golang没有提供set类型，但是map的key是不同的，基于map可以实现set，map的value为bool类型。
去重示例：
```go
func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
```









