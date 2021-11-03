如何调试 golang 程序



```
package main

import "fmt"

func main() {
	var a int64

	if a <= 10*10.1 {
		fmt.Println(a)
	}

	//t := 10 * 10
	//if a == t {
	//fmt.Println("equal")
	//}
}
```

执行以下命令查看 golang 编译时是如何对常量进行转换的
```
$ go tool compile -S -l -m 1.go

"".main STEXT size=121 args=0x0 locals=0x48
	0x0000 00000 (1.go:5)	TEXT	"".main(SB), ABIInternal, $72-0
	0x0000 00000 (1.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (1.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (1.go:5)	PCDATA	$0, $-2
	0x000d 00013 (1.go:5)	JLS	114
	0x000f 00015 (1.go:5)	PCDATA	$0, $-1
	0x000f 00015 (1.go:5)	SUBQ	$72, SP
	0x0013 00019 (1.go:5)	MOVQ	BP, 64(SP)
	0x0018 00024 (1.go:5)	LEAQ	64(SP), BP
	0x001d 00029 (1.go:5)	PCDATA	$0, $-2
	0x001d 00029 (1.go:5)	PCDATA	$1, $-2
	0x001d 00029 (1.go:5)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (1.go:5)	FUNCDATA	$1, gclocals·568470801006e5c0dc3947ea998fe279(SB)
	0x001d 00029 (1.go:5)	FUNCDATA	$2, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x001d 00029 (1.go:5)	FUNCDATA	$3, "".main.stkobj(SB)
	0x001d 00029 (1.go:9)	PCDATA	$0, $0
	0x001d 00029 (1.go:9)	PCDATA	$1, $0
	0x001d 00029 (1.go:9)	MOVQ	$0, (SP)
	0x0025 00037 (1.go:9)	CALL	runtime.convT64(SB)
	0x002a 00042 (1.go:9)	PCDATA	$0, $1
	0x002a 00042 (1.go:9)	MOVQ	8(SP), AX
	0x002f 00047 (1.go:9)	PCDATA	$1, $1
	0x002f 00047 (1.go:9)	XORPS	X0, X0
	0x0032 00050 (1.go:9)	MOVUPS	X0, ""..autotmp_1+48(SP)
	0x0037 00055 (1.go:9)	PCDATA	$0, $2
	0x0037 00055 (1.go:9)	LEAQ	type.int64(SB), CX
	0x003e 00062 (1.go:9)	PCDATA	$0, $1
	0x003e 00062 (1.go:9)	MOVQ	CX, ""..autotmp_1+48(SP)
	0x0043 00067 (1.go:9)	PCDATA	$0, $0
	0x0043 00067 (1.go:9)	MOVQ	AX, ""..autotmp_1+56(SP)
	0x0048 00072 (1.go:9)	PCDATA	$0, $1
	0x0048 00072 (1.go:9)	PCDATA	$1, $0
	0x0048 00072 (1.go:9)	LEAQ	""..autotmp_1+48(SP), AX
	0x004d 00077 (1.go:9)	PCDATA	$0, $0
	0x004d 00077 (1.go:9)	MOVQ	AX, (SP)
	0x0051 00081 (1.go:9)	MOVQ	$1, 8(SP)
	0x005a 00090 (1.go:9)	MOVQ	$1, 16(SP)
	0x0063 00099 (1.go:9)	CALL	fmt.Println(SB)
	0x0068 00104 (<unknown line number>)	MOVQ	64(SP), BP
	0x006d 00109 (<unknown line number>)	ADDQ	$72, SP
	0x0071 00113 (<unknown line number>)	RET
	0x0072 00114 (<unknown line number>)	NOP
	0x0072 00114 (1.go:5)	PCDATA	$1, $-1
	0x0072 00114 (1.go:5)	PCDATA	$0, $-2
	0x0072 00114 (1.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0077 00119 (1.go:5)	PCDATA	$0, $-1
	0x0077 00119 (1.go:5)	JMP	0
```

编译时使用 runtime.convT64 转换为了 int64 类型 ？？


参考：
https://juejin.cn/post/6844903965541285896

https://xie.infoq.cn/article/322f9e84abf217f822f260d46

