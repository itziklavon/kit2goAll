package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"./log"

	"./dao"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/go/delete/user/")
	var myList []int
	fmt.Println("deleting player: " + id)
	var playerId, err = strconv.Atoi(id)
	if err != nil {
		log.ErrorException("couldn't convert id to integer", err)

	} else {
		myList = append(myList, playerId)
		go dao.RemovePlayerIds(myList)
		go dao.RemovePlayerDetails(myList)
		go dao.RemoveWallet(myList)
	}

}

func handleRequests() {
	http.HandleFunc("/go/delete/user/", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	log.SetLogOutput("/http/http.log")
	handleRequests()
}
