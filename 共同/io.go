package common

import (
	"io"
	"net"

	"github.com/p4gefau1t/trojan-go/log"
)
//定义struct，rawReader字段是一个io.Reader借口
type RewindReader struct {
	rawReader  io.Reader
	buf        []byte
	bufReadIdx int
	rewound    bool
	buffering  bool
	bufferSize int
}
//方法，接受者为RewindReader类型的指针，实现了io.Reader方法
func (r *RewindReader) Read(p []byte) (int, error) {
        //判断r中的bool类型变量的值
	if r.rewound {
                //判断[]byt类型的长度
		if len(r.buf) > r.bufReadIdx {
                        //将r中的[]byte的一部分复制至p中
			n := copy(p, r.buf[r.bufReadIdx:])
                        //相加
			r.bufReadIdx += n
                        //返回复制完成后的数据
			return n, nil
		}   
                //将r.rewoud修改为false数据
		r.rewound = false //all buffering content has been read
	}
        //读取数据
	n, err := r.rawReader.Read(p)
        //判断r.buffering是否为true
	if r.buffering {
                //向r.buf中追加数据
		r.buf = append(r.buf, p[:n]...)
                //判断r.buf的值，决定是否输出bug日在
		if len(r.buf) > r.bufferSize*2 {
			log.Debug("read too many bytes!")
		}
	}
        //返回n
	return n, err
}
//ReadByte函数用于测试是否能正常读取byte
func (r *RewindReader) ReadByte() (byte, error) {
	buf := [1]byte{}
	_, err := r.Read(buf[:])
	return buf[0], err
}

func (r *RewindReader) Discard(n int) (int, error) {
        //创建数组
	buf := [128]byte{}
        //若是n小于128直接丢弃
	if n < 128 {
		return r.Read(buf[:n])
	}
        //循环，检测是否出错，发生错误直接返回 丢弃
	for discarded := 0; discarded+128 < n; discarded += 128 {
		_, err := r.Read(buf[:])
		if err != nil {
			return discarded, err
		}
	}
        //数据包不完整直接丢弃
	if rest := n % 128; rest != 0 {
		return r.Read(buf[:rest])
	}
        //返回n,并且错误值为nil
	return n, nil
}

func (r *RewindReader) Rewind() {
        //r.bufferSize等于0直接报错
	if r.bufferSize == 0 {
		panic("no buffer")
	}
        //设置r的字段
	r.rewound = true
	r.bufReadIdx = 0
}
//设置r的字段
func (r *RewindReader) StopBuffering() {
	r.buffering = false
}

func (r *RewindReader) SetBufferSize(size int) {
        //判断size是否为0
	if size == 0 { //disable buffering
                //检测是否启用缓冲
		if !r.buffering {
			panic("reader is disabled")
		}
                //禁用缓冲
		r.buffering = false
		r.buf = nil
		r.bufReadIdx = 0
		r.bufferSize = 0
	} else {
		if r.buffering {
			panic("reader is buffering")
		}
                //启用缓冲
		r.buffering = true
		r.bufReadIdx = 0
		r.bufferSize = size
		r.buf = make([]byte, 0, size)
	}
}
//定义struct，字段为匿名字段通过数据类型访问
type RewindConn struct {
	net.Conn
	*RewindReader
}
//RwindConn实现Read方法
func (c *RewindConn) Read(p []byte) (int, error) {
	return c.RewindReader.Read(p)
}
//工厂变量创建新的RewindConn
func NewRewindConn(conn net.Conn) *RewindConn {
	return &RewindConn{
		Conn: conn,
		RewindReader: &RewindReader{
			rawReader: conn,
		},
	}
}
//定义struct，rawWriter是io.Writer类型的interface
type StickyWriter struct {
	rawWriter   io.Writer
	writeBuffer []byte
	MaxBuffered int
}
//StickyWriter实现Writer
func (w *StickyWriter) Write(p []byte) (int, error) {
        //判断缓冲区是否启用
	if w.MaxBuffered > 0 {
		w.MaxBuffered--
                //向缓冲区内添加数据
		w.writeBuffer = append(w.writeBuffer, p...)
		if w.MaxBuffered != 0 {
			return len(p), nil
		}
		w.MaxBuffered = 0
                //检测写入是否出错
		_, err := w.rawWriter.Write(w.writeBuffer)
		w.writeBuffer = nil
		return len(p), err
	}
        返回
	return w.rawWriter.Write(p)
}
