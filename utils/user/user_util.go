package user

import (
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/userDomain"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/server/user_proto"
)

func DataToUser(data *userDomain.User) *user_proto.User {
	return &user_proto.User{
		Id:       data.Id,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Loyalty:  data.Loyalty,
		Payment:  data.Payment,
		Password: data.Password,
		Watched:  DataToProtoWatched(data.Watched),
		Wishlist: DataToProtoWishlist(data.Wishlist),
	}
}

func DataToProtoWishlist(data []*userDomain.Wishlist) []*user_proto.Wishlist {
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

func DataToProtoWatched(data []*userDomain.Watched) []*user_proto.Watched {
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

func DataToDomainWishlist(data []*user_proto.Wishlist) []*userDomain.Wishlist {
	var wishlist []*userDomain.Wishlist
	for _, item := range data {
		wishlist = append(wishlist, &userDomain.Wishlist{
			IdUser:     item.IdUser,
			IdMovie:    item.IdMovie,
			IdWishlist: item.IdWishlist,
		})
	}
	return wishlist
}

func DataToDomainWatched(data []*user_proto.Watched) []*userDomain.Watched {
	var watched []*userDomain.Watched
	for _, item := range data {
		watched = append(watched, &userDomain.Watched{
			IdWatched: item.IdWatched,
			IdMovie:   item.IdMovie,
			IdUser:    item.IdUser,
			Rate:      item.Rate,
		})
	}

	return watched
}