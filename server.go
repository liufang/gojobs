package main

import (
	"net"
	"log"
	"bytes"
	"fmt"
	"strconv"
)

type server struct {
	jobs List
}

type job struct {
	conn net.Conn
	buf bytes.Buffer
}

func start() {
	ln, err := net.Listen("tcp", ":8080")
	s:=&server{}
	defer ln.Close()
	if err != nil {
		// handle error
		log.Println("listen error")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Println("accept error")
		}
		go handleConnection(s, conn)
	}
}

//服务处理句柄
//TODO 根据命令行内容, 进行命令行操作
func handleConnection(s *server, conn net.Conn) {
	var j job
	j.conn = conn
	defer j.conn.Close()
	for {
		j.readNBuffer(5)
		datalen,err := strconv.Atoi(string(j.buf.Next(5)))
		//读取数据错误, 直接断开连接
		if(err != nil) {
			return
		}
		fmt.Println("data len:", datalen)
		j.readNBuffer(datalen)
		jobstr:=string(j.buf.Next(datalen))
		fmt.Println(jobstr)
		job:=NewJob(jobstr)
		s.jobs.Put(*job)
		//成功返回客户端成功标识
		j.conn.Write([]byte("OK"))
		return
	}
}

//读取指定字节的数据流
func (c *job) readNBuffer(n int) error {
	var buf = make([]byte, 1024)
	l:=0
	for {
		//conn.SetReadBuffer(buf)
		len,err:=c.conn.Read(buf)
		l+=len
		if(err==nil) {
			c.buf.Write(buf[:len])
		} else {
			return err
		}
		if(l>=n) {
			break
		}
	}

	return nil
}