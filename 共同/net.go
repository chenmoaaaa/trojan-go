package common

import (
	"fmt"
	"net"
	"strconv"
)
//创建常量
const (
	KiB = 1024
	MiB = KiB * 1024
	GiB = MiB * 1024
)
//return string
func HumanFriendlyTraffic(bytes uint64) string {
        //判断大小进行不同的打印输出
	if bytes <= KiB {
		return fmt.Sprintf("%d B", bytes)
	}
	if bytes <= MiB {
		return fmt.Sprintf("%.2f KiB", float32(bytes)/KiB)
	}
	if bytes <= GiB {
		return fmt.Sprintf("%.2f MiB", float32(bytes)/MiB)
	}
	return fmt.Sprintf("%.2f GiB", float32(bytes)/GiB)
}

func PickPort(network string, host string) int {
       //switch判断语句，依次判断case value为true则执行语句
	switch network {
        
	case "tcp":
                
		for retry := 0; retry < 16; retry++ {
                        //net.Listen用于监听IP和端口,tcp是协议类型，host是IP地址
			l, err := net.Listen("tcp", host+":0")
                        //判断error
			if err != nil {
				continue
			}
                        //不论是否匹配都关闭l
			defer l.Close()
                        //获取端口
			_, port, err := net.SplitHostPort(l.Addr().String())
                        //处理错误
			Must(err)
                        //将string端口转换为int端口
			p, err := strconv.ParseInt(port, 10, 32)
			Must(err)
                        //return 端口
			return int(p)
		}
	case "udp":
		for retry := 0; retry < 16; retry++ {
                        //UDP协议监听
			conn, err := net.ListenPacket("udp", host+":0")
			if err != nil {
				continue
			}
                        //关闭关闭listener
			defer conn.Close()
                        //获取端口，并返回
			_, port, err := net.SplitHostPort(conn.LocalAddr().String())
			Must(err)
			p, err := strconv.ParseInt(port, 10, 32)
			Must(err)
			return int(p)
		}
	default:
		return 0
	}
	return 0
}
