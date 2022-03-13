package main

import (
	"Gozhijiao/zhijiaoyun"
)

func main() {
	var user string = "2011305"
	var userPwd string = "xuduo123A"

	r := zhijiaoyun.New(user, userPwd)
	r.Login()
}
