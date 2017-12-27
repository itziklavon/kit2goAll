package main

import (
	"fmt"

	"./log"
	"./mongo"
)

func main() {
	log.SetLogOutput("/mongo/mongo.log")
	bson := mongo.Find("financials", "reasons")
	if len(bson) > 0 {
		fmt.Println(bson)
	}
	fmt.Println("hello")
}
