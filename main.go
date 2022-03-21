package main

import (
	"Gozhijiao/vocational"
	"time"

	"github.com/imroc/req/v3"
)

//	_ "github.com/mattn/go-sqlite3"

func main() {
	req.DevMode()
	c := vocational.Login("2017165", "Lty1964664291")
	c.GetToday()
	// c.NewGetStuFaceActivityList()
	//c.GetToday()
	// f, err := json.Marshal(c)

	// if err != nil {
	// 	panic(err)
	// }
	// file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fr, err := ioutil.ReadFile("config.json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(fr))
	// file.Write(f)
	// file.Close()
	for {
		c.NewGetStuFaceActivityList()
		time.Sleep(10 * time.Second)
	}

}
