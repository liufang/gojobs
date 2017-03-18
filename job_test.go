package main

import (
	"testing"
	"strconv"
	)

func testPut() *List {
	list := new(List)
	//入队列
	for i:=1; i<=5; i++ {
		list.Put(Job{1,"a test by fang " + strconv.Itoa(i), nil, nil})
	}
	return list
}

//测试出队入队
func TestList(t *testing.T) {
	list := testPut()

	//出队列
	for i:=1; i<=7; i++ {
		job, error := list.Push()
		if(error == nil) {
			t.Log(job.content)
		} else {
			t.Log("error")
		}
		t.Log("push ...");
	}

	t.Log("success")
}

func TestPut(t *testing.T) {
	list:=testPut()
	t.Log(list.count)
}