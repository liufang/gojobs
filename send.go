package main

import (
	"github.com/jordan-wright/email"
	"net/smtp"
	"fmt"
)

func main() {
	//test list
	list := new(List)
	list.Put(Job{1,"test by fang", nil, nil})
	list.Put(Job{1,"test by fang", nil, nil})
	list.Put(Job{1,"test by fang", nil, nil})
	list.Put(Job{1,"test by fang", nil, nil})
	list.Put(Job{1,"test by fang", nil, nil})

	for i:=1; i<=5; i++ {
		job, error := list.Push()
		if(error == nil) {
			fmt.Println(job.content)
		} else {
			fmt.Println("error")
		}
	}

	_, error := list.Push()
	if(error != nil) {
		fmt.Println("empty list test by fang");
	}

	fmt.Println("list count: ", list.count)

	send();
}


func send() {
	e := email.NewEmail()
	e.From = "\"fang.liu<script>alert(\"test\");</script>\" <system@lisa.org.cn>"
	//e.From = "fang.liu <info@glod.vip>"

	e.To = []string{"1033289127@qq.com"}
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1><script>alert(\"test\");</script>")
	err := e.Send("smtp.mxhichina.com:25", smtp.PlainAuth("", "system@lisa.org.cn", "fangadmin@123456", "smtp.mxhichina.com"))

	//err := e.Send("email-smtp.us-west-2.amazonaws.com:25", smtp.PlainAuth("", "AKIAIO3WQDLMJIZKFZOQ", "AnEHbG08Fzz1LBsUyAvqaeq8hINZDDZzGFgcdvY+oYBG", "email-smtp.us-west-2.amazonaws.com"))

	if(err != nil) {
		fmt.Println("send error:", err)
	} else {
		fmt.Println("send success:");
	}
}