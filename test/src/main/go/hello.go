package main

import (
	"./log"
)

func main() {
	log.SetLogOutput("/hello/hello.log")
	for i := 0; i < 100; i++ {
		log.Debug("error")
		log.Debug("Ah Nu, ze Shalom!")
	}
}
