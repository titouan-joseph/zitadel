package grpc

import (
	"context"
	"github.com/caos/zitadel/internal/errors"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) GetUserByID(ctx context.Context, request *UserID) (*User, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-fj74G", "Not implemented")
}

func (s *Server) IsUserUnique(ctx context.Context, request *UniqueUserRequest) (*UniqueUserResponse, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dke9f", "Not implemented")
}

func (s *Server) GetMyUserProfile(ctx context.Context, _ *empty.Empty) (*UserProfile, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-fis93", "Not implemented")
}

func (s *Server) GetMyUserEmail(ctx context.Context, _ *empty.Empty) (*UserEmail, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-93j5d", "Not implemented")
}

func (s *Server) GetMyUserPhone(ctx context.Context, _ *empty.Empty) (*UserPhone, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-Hj75G", "Not implemented")
}

func (s *Server) GetMyUserAddress(ctx context.Context, _ *empty.Empty) (*UserAddress, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-21jd4", "Not implemented")
}

func (s *Server) GetMyMfas(ctx context.Context, _ *empty.Empty) (*MultiFactors, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-vkl9i", "Not implemented")
}

func (s *Server) RegisterUser(ctx context.Context, request *RegisterUserRequest) (*User, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-03idx", "Not implemented")
}

func (s *Server) RegisterUserWithExternal(ctx context.Context, request *RegisterUserExternalIDPRequest) (*User, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dlwpd", "Not implemented")
}

func (s *Server) UpdateMyUserProfile(ctx context.Context, request *UpdateUserProfileRequest) (*UserProfile, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dlep3", "Not implemented")
}

func (s *Server) ChangeMyUserEmail(ctx context.Context, request *UpdateUserEmailRequest) (*UserEmail, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-lme45", "Not implemented")
}

func (s *Server) VerifyMyUserEmail(ctx context.Context, request *VerifyMyUserEmailRequest) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-poru7", "Not implemented")
}

func (s *Server) VerifyUserEmail(ctx context.Context, request *VerifyUserEmailRequest) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dlep3", "Not implemented")
}

func (s *Server) ResendMyEmailVerificationMail(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dh69i", "Not implemented")
}

func (s *Server) ResendEmailVerificationMail(ctx context.Context, id *UserID) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-mAq23", "Not implemented")
}

func (s *Server) ChangeMyUserPhone(ctx context.Context, request *UpdateUserPhoneRequest) (*UserPhone, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dk45g", "Not implemented")
}

func (s *Server) VerifyMyUserPhone(ctx context.Context, request *VerifyUserPhoneRequest) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-ol6gE", "Not implemented")
}

func (s *Server) ResendMyPhoneVerificationCode(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-Wk8rf", "Not implemented")
}

func (s *Server) UpdateMyUserAddress(ctx context.Context, request *UpdateUserAddressRequest) (*UserAddress, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-cmt7F", "Not implemented")
}

func (s *Server) SetMyPassword(ctx context.Context, request *PasswordRequest) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-pl9c2", "Not implemented")
}
func (s *Server) RequestPasswordReset(ctx context.Context, request *ResetPasswordRequest) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-xm5hW", "Not implemented")
}

func (s *Server) PasswordReset(ctx context.Context, request *ResetPassword) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-mce5W", "Not implemented")
}

func (s *Server) ChangeMyPassword(ctx context.Context, request *PasswordChange) (*empty.Empty, error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dlo6G", "Not implemented")
}

func (s *Server) AddMfaOTP(ctx context.Context, _ *empty.Empty) (_ *MfaOtpResponse, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-al35G", "Not implemented")
}

func (s *Server) VerifyMfaOTP(ctx context.Context, request *VerifyMfaOtp) (_ *MfaOtpResponse, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-kgjZ7", "Not implemented")
}

func (s *Server) RemoveMfaOTP(ctx context.Context, _ *empty.Empty) (_ *empty.Empty, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-9k46d", "Not implemented")
}

func (s *Server) SkipMfaInit(ctx context.Context, request *SkipMfaInitRequest) (_ *empty.Empty, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dl3i5", "Not implemented")
}

func (s *Server) VerifyUserInit(ctx context.Context, request *VerifyUserInitRequest) (_ *empty.Empty, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-itk6h", "Not implemented")
}

func (s *Server) ResendUserInitMail(ctx context.Context, request *UserID) (_ *empty.Empty, err error) {
	return nil, errors.ThrowUnimplemented(nil, "GRPC-dk467", "Not implemented")
}
