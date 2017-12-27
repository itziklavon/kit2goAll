package dao

import "database/sql"

func RmoveWallet() {
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
			panic(err.Error())
		}
		myList = append(myList, tag.ID)
	}
}
