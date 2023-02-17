package prints

import "fmt"

func PrintVar(v interface{}) {
	fmt.Printf("######\n")
	fmt.Printf("type:%T;\nvalue:%v;\n+value:%+v;\n", v, v, v)
}

/**
Println：(终端中有显示)
1. 用默认的类型格式T显示方式将传入的参数写入到标准输出里面
2. 多个传入参数之间使用空格分隔，
3. 在显示的最后追加换行符，
4. 返回值为 写入标准输出的字节数和写入过程中遇到的问题。

Printf：(终端中有显示)
1. 用传入的格式化规则符将传入的变量写入到标准输出里面
2. 返回值为 写入标准输出的字节数和写入过程中遇到的问题。

Sprintf: (终端中不会显示,格式化字符串使用)
1. 用传入的格式化规则符将传入的变量格式化，
2. 返回为 格式化后的字符串。

格式化规则符如下：
	%v 输出结构体 {10 30}
	%+v 输出结构体显示字段名 {one:10 tow:30}
	%#v 输出结构体源代码片段 main.Point{one:10, tow:30}
	%T 输出值的类型            main.Point
	%t 输出格式化布尔值      true
	%d`输出标准的十进制格式化 100
	%b`输出标准的二进制格式化 99 对应 1100011
	%c`输出定整数的对应字符  99 对应 c
	%x`输出十六进制编码  99 对应 63
	%f`输出十进制格式化  99 对应 63
	%e`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	%E`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	%s 进行基本的字符串输出   "\"string\""  对应 "string"
	%q 源代码中那样带有双引号的输出   "\"string\""  对应 "\"string\""
	%p 输出一个指针的值   &jgt 对应 0xc00004a090
	% 后面使用数字来控制输出宽度 默认结果使用右对齐并且通过空格来填充空白部分
	%2.2f  指定浮点型的输出宽度 1.2 对应  1.20
	%*2.2f  指定浮点型的输出宽度对齐，使用 `-` 标志 1.2 对应  *1.20
*/
