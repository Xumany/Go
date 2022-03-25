package main

import (
	"Gozhijiao/vocational"
	"github.com/imroc/req/v3"
)

func main() {
	req.DevMode()
	c := vocational.Login("2011305", "xuduo123A")
	c.GetDate("2022-03-24")
	c.NewGetStuFaceActivityList()
}
