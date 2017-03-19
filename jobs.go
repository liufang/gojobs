package main

import "math/rand"

//定义job结构
type Job struct{
	id int
	content string
	next *Job
	previous *Job
}

//构造Job
func NewJob(content string) *Job {
	return &Job{id:rand.Intn(32),content:content}
}