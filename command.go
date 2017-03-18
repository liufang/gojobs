package main

import "net"

type Command struct {
	name string
	content string
}

//从网络链接创建命令
func CreateByConn(conn Conn) {
	//读取两个字节命令行数据总长度

	//读取命令行名字

	//读取命令行参数
}