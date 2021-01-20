package user_controller

import (
	"context"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/user_domain"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/user_proto"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/services/user_services"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/crypt_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger_utils"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/user_utils"
)

type Server struct {
}

func (s *Server) CreateUser(ctx context.Context, request *user_proto.CreateUserRequest) (*user_proto.CreateUserResponse, error) {
	res, err := user_services.Services.Create(user_domain.User{
		Name:     request.GetUser().GetName(),
		Password: crypt_utils.GetMd5(request.GetUser().GetPassword()),
		Email:    request.GetUser().GetEmail(),
		Phone:    request.GetUser().GetPhone(),
	})

	if err != nil {
		logger_utils.Error("Error while trying to service create user", err)
	}

	// u can use OpenID if u want to secure data. in case i dont use that

	return &user_proto.CreateUserResponse{
		User: user_utils.DataToUser(res),
	}, nil
}

func (s *Server) GetUserById(ctx context.Context, request *user_proto.GetUserByIdRequest) (*user_proto.GetUserByIdResponse, error) {
	res, err := user_services.Services.GetUser(request.GetIdUser())
	if err != nil {
		return nil, err
	}

	return &user_proto.GetUserByIdResponse{
		User: user_utils.DataToUser(res),
	}, nil

}

func (s *Server) UpdateUser(ctx context.Context, request *user_proto.UpdateUserRequest) (*user_proto.UpdateUserResponse, error) {



	UserUpdate := user_domain.User{
		Id:       request.GetUser().GetId(),
		Name:     request.GetUser().GetName(),
		Email:    request.GetUser().GetEmail(),
		Password: crypt_utils.GetMd5(request.GetUser().GetPassword()),
		Phone:    request.GetUser().GetPhone(),
		Payment:  request.GetUser().GetPayment(),
		Loyalty:  request.GetUser().GetLoyalty(),
		Wishlist: user_utils.DataToDomainWishlist(request.GetUser().GetWishlist()),
		Watched:  user_utils.DataToDomainWatched(request.GetUser().GetWatched()),
	}

	res, err := user_services.Services.Update(UserUpdate)

	if err != nil {
		return nil, err
	}

	return &user_proto.UpdateUserResponse{
		User: user_utils.DataToUser(res),
	}, nil
}
