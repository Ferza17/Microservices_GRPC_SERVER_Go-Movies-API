package user_domain

import (
	"database/sql"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/datasources/mysql"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/errors_util"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/query_utils"
)

func (u *User) Save() error {

	builder := query_utils.NewQueryBuilder()
	query := builder.
		Insert("users").
		Columns("Name", "Email", "Password", "Phone").
		Values().
		Build()

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

func (u *User) GetUser() error {
	builder := query_utils.NewQueryBuilder()
	query := builder.
		Select("users").
		Columns().
		Where("Id").
		Build()

	logger_utils.Info(query)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger_utils.Error("Error While Preparing query ", err)
		return errors_util.Internal("Error to process data ")
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	// It can be error if payment null
	var Payment sql.NullFloat64
	// Handle Null Loyalty
	var Loyalty sql.NullInt32

	//

	//var IdWishlist sql.NullInt64
	//var IdMovie sql.NullString
	//var IdUser sql.NullInt64

	//var Wishlist []interface{}
	var Wishlist sql.NullString

	//Wishlist = append(Wishlist, IdWishlist)
	//Wishlist = append(Wishlist, IdMovie)
	//Wishlist = append(Wishlist, IdUser)

	//
	//var Watched []interface{}
	var Watched sql.NullString

	//var IdWatched sql.NullInt64
	//var Rate sql.NullFloat64

	//Watched = append(Watched, IdWatched)
	//Watched = append(Watched, IdMovie)
	//Watched = append(Watched, IdUser)
	//Watched = append(Watched, Rate)

	if err := result.Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Phone,
		&Payment,
		&Loyalty,
		&Wishlist,
		&Watched,
	); err != nil {
		logger_utils.Error("error when trying to get user by id", err)
		return errors_util.Internal("error when trying to get user by id")
	}

	return nil

}

func (u *User) Update() error {
	builder := query_utils.NewQueryBuilder()
	query := builder.
		Update("users").
		Columns("Id", "Name", "Email", "Password", "Phone").
		Where("Id").
		Build()
	logger_utils.Info(query)

	stmt, err := mysql.Client.Prepare(query)
	if err != nil {
		logger_utils.Error("error when trying to prepare save user statement", err)
		return errors_util.Internal("Database Error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(u.Id, u.Name, u.Email, u.Password, u.Phone, u.Id); err != nil {
		logger_utils.Error("error when trying to execute user statement", err)
		return errors_util.Internal("Database Error")
	}
	return nil

}
