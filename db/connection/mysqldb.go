package connection

import (
	"database/sql"
	"vending-machine/utils"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMysql() *sql.DB {
	var err error

	dbDriver := utils.GetEnvByKey("DB_DRIVER")
	dbName := utils.GetEnvByKey("DB_NAME")
	dbUser := utils.GetEnvByKey("DB_USER")
	dbPass := utils.GetEnvByKey("DB_PASS")

	DB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	return DB
}
