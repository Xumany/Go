package main

import (
	"Gozhijiao/vocational"
	"time"
)

func main() {
	// req.DevMode()

	c := vocational.Login("2011307", "19164440516XQC@")
	for {
		c.GetToday()
		c.NewGetStuFaceActivityList()
		c.IsJoinActivities()
		time.Sleep(5 * time.Second)
	}

}
