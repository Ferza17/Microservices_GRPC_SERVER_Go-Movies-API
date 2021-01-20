package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/env_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger_utils"
)

var (
	Client   *sql.DB
	username = env_utils.GetEnvironmentVariable("mysqlUsersUsername")
	password = env_utils.GetEnvironmentVariable("mysql_users_password")
	host     = env_utils.GetEnvironmentVariable("mysql_users_host")
	schema   = env_utils.GetEnvironmentVariable("mysqlUsersSchema")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger_utils.Error("error while connect to DB", err)
		panic(err)
	}

	logger_utils.Info("Database Successfully configured!")
}
