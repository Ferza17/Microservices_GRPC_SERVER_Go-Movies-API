package user_domain

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/query_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/datasources/mysql"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/errors_util"
)

func (u *User) Save() error {

	builder := query_utils.NewQueryBuilder()
	query := builder.
		Insert("users").
		Columns("Name", "Email", "Password", "Phone").Build()

	logger_utils.Info(query)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger_utils.Error("Error While Preparing query ", err)
		return errors_util.Internal("Error to process data ")
	}
	defer stmt.Close()

	insetResult, err := stmt.Exec(u.Name, u.Email, u.Password, u.Phone)
	if err != nil {
		logger_utils.Error("Error while trying to exec query ", err)
		return errors_util.Internal("Error to save to database ")
	}

	userId, err := insetResult.LastInsertId()
	if err != nil {
		logger_utils.Error("Error while trying to get id", err)
		return errors_util.Internal("Error while trying to get id")
	}
	u.Id = userId

	return nil
}
