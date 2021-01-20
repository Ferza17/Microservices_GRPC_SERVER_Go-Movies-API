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
	res, err := user_services.UserServices.Create(user_domain.User{
		Name:     request.GetUser().GetName(),
		Password: crypt_utils.GetMd5(request.GetUser().GetPassword()),
		Email:    request.GetUser().GetEmail(),
		Phone:    request.GetUser().GetPhone(),
	})

	if err != nil {
		logger_utils.Error("Error while trying to service create user", err)
	}

	//TODO u can use OpenID if u want to secure data. in case i dont use that

	return &user_proto.CreateUserResponse{
		User: user_utils.DataToFullUserPB(res),
	}, nil
}
