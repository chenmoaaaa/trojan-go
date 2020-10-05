//声明go文件属于哪个包
package option
//导入包
import "github.com/p4gefau1t/trojan-go/common"
//定义interface，type是自定义关键字，Handler是借口名
type Handler interface {
        //interface（接口）名同返回值
	Name() string
	Handle() error
	Priority() int
}
//创建一个字典，key是string类型，value是Handler类型
var handlers = make(map[string]Handler)
//func关键字，registerhandler函数名，（）是参数
func RegisterHandler(h Handler) {
        //定义handlers中的一键值对
	handlers[h.Name()] = h
}
//创建函数pop0ptionhandler,返回值为两个包含error值
func PopOptionHandler() (Handler, error) {
        //将handler类型的值maxhandler初始化为nil
	var maxHandler Handler = nil
        //for循环遍历字典，抛弃string类型的key，只保留value
	for _, h := range handlers{
                //if语句进行判断
		if maxHandler == nil || maxHandler.Priority() < h.Priority() {
			maxHandler = h
		}
	}
        //如果没有进行正常赋值操作则返回错误error值
	if maxHandler == nil {
		return nil, common.NewError("no option left")
	}
        //delete删除一键值对
	delete(handlers, maxHandler.Name())
        //返回一个interface类型的值，error值为nil
	return maxHandler, nil
}
