//所有.go文件必须在第一行声明自己所属的包，
//同一个目录下的.go文件必须使用同一个包名。
package main

//导入其他目录下的包，标准库的包不需要加绝对路径
import (
        //flag包用于处理命令行参数。
	"flag"

	"github.com/p4gefau1t/trojan-go/option"

	_ "github.com/p4gefau1t/trojan-go/component"
	"github.com/p4gefau1t/trojan-go/log"
)
//func关键字用来声明函数，main是方法名，
//main方法是一切go程序的入口。
func main() {
        //接受用户传入的命令行参数
        //将其解析为对应变量的值
	flag.Parse()
        //无限for循环
	for {
                //接受函数返回的值和error值
		h, err := option.PopOptionHandler()
                //如果发生错误进行报错
		if err != nil {
                        //输出日在信息
			log.Fatal("invalid options")
		}
                //判断error的值，不为空则退出循环
                //用来退出程序，或者在严重错误时终止
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
