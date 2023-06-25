package grpc_api

import (
	"bookkeeper/internal/db"
	"bookkeeper/internal/protogen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *protogen.User {
	return &protogen.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
