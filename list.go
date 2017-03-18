package main

import (
	"errors"
)

type List struct{
	first *Job
	last *Job
	count int
}

//添加一个任务到列表中
func (list *List) Put(job Job) error {
	if(list.count == 0) {
		list.first = &job
		list.last =&job
	} else {
		//加入现有的链表
		tmpJob := list.first
		list.first = &job
		job.next = tmpJob
		tmpJob.previous = &job
	}
	list.count++
	return nil
}

//弹出最早进入队列的任务
func (list *List) Push() (*Job, error) {
	if(list.last != nil) {
		tmpJob := list.last
		list.last = tmpJob.previous
		list.count--
		return tmpJob, nil
	}

	return new(Job),errors.New("empty list");
}