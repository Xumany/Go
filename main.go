package main

import (
	"Gozhijiao/vocational"
	"time"
)

//	_ "github.com/mattn/go-sqlite3"

func main() {

	c := vocational.Login("2017165", "Lty1964664291")
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
		ret := vocational.GetToday(*c)
		vocational.NewGetStuFaceActivityList(ret)
		time.Sleep(10 * time.Second)
	}

}
