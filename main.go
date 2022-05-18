package main

import (
	"Gozhijiao/zjy"
	"fmt"
	"time"

	"github.com/fufuok/xdaemon"
	"github.com/imroc/req/v3"
)

func main() {
	req.DevMode()
	logfile := "run.log"
	xdaemon.Background(logfile, true)
	zjyinfo, err := zjy.Login("2011305", "xuduo123A", "SCT11035TaAJBnQTTAJND3rnfQTuPUuEU")
	if err != nil {
		fmt.Println(err)
		return
	}
	zjyinfo.Run(5 * time.Second)
}
