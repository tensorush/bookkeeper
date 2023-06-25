package grpc_api

import (
	"context"

	"bookkeeper/internal/db"
	"bookkeeper/internal/protogen"
	"bookkeeper/internal/utils"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *protogen.VerifyEmailRequest) (*protogen.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	txResult, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email")
	}

	rsp := &protogen.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *protogen.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, fieldViolation("email_id", err))
	}

	if err := utils.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))
	}

	return violations
}
