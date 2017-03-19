package main 

import (
	"testing"
	"net"
	"bytes"
	"strconv"
)

//测试服务器功能
func TestServer(t *testing.T) {
	start()
	return
	ln, _ := net.Listen("tcp", ":8080")
	defer ln.Close()
	conn, _ := ln.Accept()
	defer conn.Close()
	h(conn, t)
}



//服务处理句柄
func h(conn net.Conn, t *testing.T) {
	defer conn.Close()
	var b bytes.Buffer

	//for {
		test,_:=readNBuffer2(conn, b, 5, t)

		t.Log(test.String())

		len := test.Next(5)
		datalen, err := strconv.Atoi(string(len))
		t.Log(datalen)
		if err != nil {
			t.Log("conver data len error")
		} else {
			bb,_ := readNBuffer2(conn, b, datalen, t)
			t.Log(bb.String())
		}


	//}
}


//读取指定字节的数据流
func readNBuffer2(conn net.Conn, b bytes.Buffer, n int, t *testing.T) (bytes.Buffer, error) {
	t.Log("start read from conn")
	var buf = make([]byte, n)
	for {
		//conn.SetReadBuffer(buf)
		len,err:=conn.Read(buf)
		t.Log(len)
		t.Log(err)
		if(err==nil) {
			b.Write(buf[:len])
		} else {
			break
		}
		if(b.Len()>=n) {
			break
		}
	}

	return b, nil
}