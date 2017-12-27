package dao

import (
	"strconv"

	"../configuration"
	"../log"
	"../mysql"
)

func RemovePlayerDetails(mylist []int) {
	db, err := mysql.OpenSqlConnection(configuration.GetPropertyValue("PLAYER_DETAILS_SCHEMA_NAME"))
	defer db.Close()
	playerQuery := "DELETE FROM player_details WHERE player_id IN (" + constructQuery(mylist) + ")"
	stmt, err := db.Prepare(playerQuery)
	checkErr(err)
	res, err := stmt.Exec()
	affects, err := res.RowsAffected()
	checkErr(err)
	log.Debug(":RemovePlayerDetails: affected rows players- player details" + strconv.Itoa(int(affects)))
}
