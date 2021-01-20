package user_utils

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/user_domain"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/user_proto"
)

func DataToFullUserPB(data *user_domain.User) *user_proto.FullUserData {

	return &user_proto.FullUserData{
		User:     DataToUser(data),
		Watched:  DataToWatched(data.Watched),
		Wishlist: DataToWishlist(data.Wishlist),
	}
}

func DataToUser(data *user_domain.User) *user_proto.User {
	return &user_proto.User{
		Id:       data.Id,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Loyalty:  data.Loyalty,
		Payment:  data.Payment,
		Password: data.Password,
	}
}

func DataToWishlist(data []*user_domain.Wishlist) []*user_proto.Wishlist {
	var result []*user_proto.Wishlist
	for _, data := range data {
		result = append(result, &user_proto.Wishlist{
			IdMovie:    data.IdMovie,
			IdUser:     data.IdUser,
			IdWishlist: data.IdWishlist,
		})
	}
	return result
}

func DataToWatched(data []*user_domain.Watched) []*user_proto.Watched {
	var result []*user_proto.Watched
	for _, data := range data {
		result = append(result, &user_proto.Watched{
			IdUser:    data.IdUser,
			IdMovie:   data.IdMovie,
			IdWatched: data.IdWatched,
			Rate:      data.Rate,
		})
	}

	return result
}
