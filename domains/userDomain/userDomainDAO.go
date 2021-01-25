package userDomain

import (
	"database/sql"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/datasources/mysql/userDB"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/errors"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/query"
)

func (u *User) Save() error {
	builder := query.NewQueryBuilder()

	query := builder.
		Insert("users").
		Columns(u).
		Values().
		BuildQuery()
	stmt, err := userDB.Client.Prepare(query)
	if err != nil {
		logger.Error("Error While Preparing query ", err)
		return errors.Internal("Error to process data ")
	}
	defer stmt.Close()

	insetResult, err := stmt.Exec(builder.GetValueOf(u)...)
	if err != nil {
		logger.Error("Error while trying to exec query ", err)
		return errors.Internal("Error to save to database ")
	}

	userId, err := insetResult.LastInsertId()
	if err != nil {
		logger.Error("Error while trying to get id", err)
		return errors.Internal("Error while trying to get id")
	}
	u.Id = userId

	return nil
}

func (u *User) GetUser() error {
	builder := query.NewQueryBuilder()
	query := builder.
		Select("users").
		Columns(nil).
		Where("Id").
		BuildQuery()

	stmt, err := userDB.Client.Prepare(query)
	if err != nil {
		logger.Error("Error While Preparing query ", err)
		return errors.Internal("Error while preparing id")
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
		logger.Error("error when trying to get user by id", err)
		return errors.NotFound("no rows in result set")
	}

	return nil

}

func (u *User) Update() error {
	builder := query.NewQueryBuilder()
	query := builder.
		Update("users").
		Columns(u).
		Where("Id").
		BuildQuery()

	stmt, err := userDB.Client.Prepare(query)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.Internal("Database Error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(builder.GetValueOf(u)...); err != nil {
		logger.Error("error when trying to execute user statement", err)
		return errors.Internal("Database Error")
	}
	return nil

}

// TODO:Login @User
func (u *User) Login() error {
	return nil
}

// TODO Add wishlist with id Film
func (u *User) AddWishlist() error {
	//...
	return nil
}

// TODO Delete wishlist with id Film
func (u *User) DeleteWishlist(idMovie string) error {
	//...
	return nil
}

// TODO get wishlist with id Film
func (u *User) GetWishlist(idMovie string) error {
	//...
	return nil
}

// TODO get All Wishlist by id
func (u *User) GetWishlists(idUser int) error {
	//...
	return nil
}

// WATCHED
