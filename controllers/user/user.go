package user

import (
	"context"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/domains/userDomain"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/protos/server/user_proto"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/services/user"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/crypt"
	"github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/logger"
	userUtils "github.com/Ferza17/Microservices_GRPC_SERVER_Go-Users-API/utils/user"
)

type Server struct {
}

func (s *Server) CreateUser(ctx context.Context, request *user_proto.CreateUserRequest) (*user_proto.CreateUserResponse, error) {
	res, err := user.Services.Create(userDomain.User{
		Name:     request.GetName(),
		Password: crypt.GetMd5(request.GetPassword()),
		Email:    request.GetEmail(),
		Phone:    request.GetPhone(),
		Payment:  request.GetPayment(),
		Loyalty:  request.GetLoyalty(),
	})

	if err != nil {
		logger.Error("Error while trying to service create user", err)
	}

	// u can use OpenID if u want to secure data. in this case i dont use that
	return &user_proto.CreateUserResponse{
		User: userUtils.DataToUser(res),
	}, nil
}

func (s *Server) GetUserById(ctx context.Context, request *user_proto.GetUserByIdRequest) (*user_proto.GetUserByIdResponse, error) {
	res, err := user.Services.GetUser(request.GetIdUser())
	if err != nil {
		return nil, err
	}

	return &user_proto.GetUserByIdResponse{
		User: userUtils.DataToUser(res),
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, request *user_proto.UpdateUserRequest) (*user_proto.UpdateUserResponse, error) {

	UserUpdate := userDomain.User{
		Id:       request.GetUser().GetId(),
		Name:     request.GetUser().GetName(),
		Email:    request.GetUser().GetEmail(),
		Password: crypt.GetMd5(request.GetUser().GetPassword()),
		Phone:    request.GetUser().GetPhone(),
		Payment:  request.GetUser().GetPayment(),
		Loyalty:  request.GetUser().GetLoyalty(),
		Wishlist: userUtils.DataToDomainWishlist(request.GetUser().GetWishlist()),
		Watched:  userUtils.DataToDomainWatched(request.GetUser().GetWatched()),
	}

	res, err := user.Services.Update(UserUpdate)

	if err != nil {
		return nil, err
	}

	return &user_proto.UpdateUserResponse{
		User: userUtils.DataToUser(res),
	}, nil
}
