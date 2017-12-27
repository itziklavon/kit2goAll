package main

import (
	"fmt"
	"./dao"
	"./configuration"
)

func main() {
	myList := dao.GetPlayerIds()
	for i := 0; i < len(myList); i++ {
		fmt.Printf("id is: %d\n", myList[i])
	}
	host := configuration.GetPropertyValue("JDBC_HOST")
	fmt.Println(host)
}
