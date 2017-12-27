package main

import (
	"./dao"
	"./log"
	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Day().Do(task)
	<-s.Start()
}

func task() {
	log.SetLogOutput("/automation-cleanup/cleanup.log")
	myList := dao.GetPlayerIds()
	if len(myList) > 0 {
		dao.RemovePlayerIds(myList)
		dao.RemovePlayerDetails(myList)
		dao.RemoveWallet(myList)
	}
}
