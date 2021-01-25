package userDB

import (
	"database/sql"
	"fmt"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/env"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = env.GetEnvironmentVariable("mysqlUsersUsername")
	password = env.GetEnvironmentVariable("mysql_users_password")
	host     = env.GetEnvironmentVariable("mysql_users_host")
	schema   = env.GetEnvironmentVariable("mysqlUsersSchema")
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
		logger.Error("error while connect to DB", err)
		panic(err)
	}

	logger.Info("Database Successfully configured!")
}
