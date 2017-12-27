package dao

import (
	"fmt"
	"strconv"
	"strings"

	"../configuration"
	"../log"
	"../mysql"
)

type Row struct {
	ID int `json:"id"`
}

func GetPlayerIds() []int {
	db, err := mysql.OpenSqlConnection(configuration.GetPropertyValue("GSS_SCHEMA_NAME"))
	defer db.Close()
	results, err := db.Query("SELECT ID FROM MdlGm_tblPlayers WHERE (Login like ? OR Email like ?)",
		configuration.GetPropertyValue("AUTOMATION_NAME"), configuration.GetPropertyValue("AUTOMATION_EMAIL"))
	if err != nil {
		log.ErrorException(":GetPlayerIds: couldn't connect to DB", err)
	}
	var myList []int
	for results.Next() {
		var row Row
		err = results.Scan(&row.ID)
		if err != nil {
			log.ErrorException(":GetPlayerIds: couldn't scan for results", err)
		}
		myList = append(myList, row.ID)
	}
	return myList
}

func RemovePlayerIds(mylist []int) {
	db, err := mysql.OpenSqlConnection(configuration.GetPropertyValue("GSS_SCHEMA_NAME"))
	defer db.Close()
	playerQuery := "DELETE FROM MdlGm_tblPlayers WHERE ID IN (" + constructQuery(mylist) + ")"
	playerL01Query := "DELETE FROM MdlGm_tblPlayers_L01 WHERE RefId IN (" + constructQuery(mylist) + ")"
	stmt, err := db.Prepare(playerQuery)
	checkErr(err)
	res, err := stmt.Exec()
	affects, err := res.RowsAffected()
	checkErr(err)
	log.Debug(":RemovePlayerIds: affected rows players :" + strconv.Itoa(int(affects)))
	stmt, err = db.Prepare(playerL01Query)
	checkErr(err)
	res, err = stmt.Exec()
	affects, err = res.RowsAffected()
	checkErr(err)
	log.Debug(":RemovePlayerIds: affected rows playersL01 :" + strconv.Itoa(int(affects)))

}

func constructQuery(myList []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(myList), " ", ",", -1), "[]")
}

func checkErr(err error) {
	if err != nil {
		log.ErrorException(":checkErr: couldn't excecute query", err)
	}
}
