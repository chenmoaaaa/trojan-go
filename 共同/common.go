//声明所属那个包
package common
//导入包
import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
)
//创建interface
type Runnable interface {
	Run() error
	Close() error
}
//func关键字，参数和返回值都是string类型
func SHA224String(password string) string {
        //创建一个基于sha224算法的hash.Hash对象
	hash := sha256.New224()
        //向hash中添加byte类型的切片
	hash.Write([]byte(password))
        //返回一个新切片不会影响底层hash的状态
	val := hash.Sum(nil)
        //设置一个空string，但不为nil
	str := ""
        //遍历切片
	for _, v := range val {
                //格式化输出遍历的切片内容，并将内容保存的str里
		str += fmt.Sprintf("%02x", v)
	}
        //return一个string
	return str
}
//返回string类型
func GetProgramDir() string {
        //获取当前绝对路径filepath.Abs，.Dir获取相当路径
        //os.Args获取简单的命令行参数以[]string类型的
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
        }
        //返回得到的绝对路径
	return dir
}
