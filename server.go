package main

import (
	"net"
	"math/big"
	"time"
	"log"
)


func start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

//服务处理句柄
func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		//确定长度
		var buf1 = make([]byte, 2)
		conn.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
        _, err := conn.Read(buf1)
        if err == nil {
        	return
        }
        len := new(big.Int)
        len.SetBytes(buf1)
		var buf = make([]byte, 100)
        log.Println("start to read from conn")
        conn.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
        n, err := conn.Read(buf)
        if err != nil {
            log.Printf("conn read %d bytes,  error: %s", n, err)
            if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
                continue
            }
            return
        }
	}
}