//声明所属那个包
package common
//导入
import (
	"fmt"
)
//创建struct（结构）
type Error struct {
	info string
}
//方法，接受者为指针类型对其修改会影响原始数据
func (e *Error) Error() string {
	return e.info
}

func (e *Error) Base(err error) *Error {
	if err != nil {
		e.info += " | " + err.Error()
	}
	return e
}
//工厂变量go语言中New变量一般为工厂变量，用于创建error值
func NewError(info string) *Error {
	return &Error{
		info: info,
	}
}
//must函数用于检测到重大错误时终止程序
func Must(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Must2(_ interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
