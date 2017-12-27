package mysql

import (
	"database/sql"

	"../configuration"
	"../log"
	_ "github.com/go-sql-driver/mysql"
)

func OpenSqlConnection(schemaType string) (*sql.DB, error) {
	user := configuration.GetPropertyValue("JDBC_USER_NAME")
	password := configuration.GetPropertyValue("JDBC_PASSWORD")
	host := configuration.GetPropertyValue("JDBC_HOST")
	connection := user + ":" + password + "@tcp(" + host + ":3306)/" + schemaType
	log.Debug(":OpenSqlConnection: connecting to db, user name, password, jdbc, scehma: " + connection)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.ErrorException("OpenSqlConnection:: couldn't connect to DB", err)
	}
	return db, err
}
