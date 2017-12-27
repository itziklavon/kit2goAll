package dao

import (
	"strconv"

	"../log"

	"../configuration"
	"../mysql"
)

func RemoveWallet(mylist []int) {
	db, err := mysql.OpenSqlConnection(configuration.GetPropertyValue("WALLET_SCHEMA_NAME"))
	defer db.Close()
	playerQuery := "DELETE FROM fin_accounts WHERE id_user IN (" + constructQuery(mylist) + ")"
	stmt, err := db.Prepare(playerQuery)
	checkErr(err)
	res, err := stmt.Exec()
	affects, err := res.RowsAffected()
	checkErr(err)
	log.Debug(":RemoveWallet: affected rows players- player details" + strconv.Itoa(int(affects)))
}
