package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"../configuration"
	//"log"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetPlayerIds() []int {
	db, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT ID, Login as Name FROM dev_bo.mdlgm_tblplayers")
	if err != nil {
		panic(err.Error())
	}
	var myList []int
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		myList = append(myList, tag.ID)
		//log.Println("current list size is: " + string(len(myList)))
	}
	return myList
}

func RmovePlayerIds() {
	db, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT ID, Login as Name FROM dev_bo.mdlgm_tblplayers WHERE (Login like ? OR Email like ?)",
		configuration.GetPropertyValue("JDBC_HOST"), configuration.GetPropertyValue("JDBC_HOST"))
	if err != nil {
		panic(err.Error())
	}
	var myList []int
	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		myList = append(myList, tag.ID)
	}
}
