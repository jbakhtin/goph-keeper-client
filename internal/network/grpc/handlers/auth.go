package handlers

import (
	"context"
	"github.com/jbakhtin/goph-keeper-client/gen/go/v1/auth"
	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/dto"
	secondary_ports "github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/ports/secondary"
	"github.com/jbakhtin/goph-keeper-client/internal/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	_ secondary_ports.ApiPort = &AuthHandler{}
)

type AuthHandler struct {
	logger *zap.Logger
	client auth.AuthServiceClient
}

func NewAuthHandler(lgr *zap.Logger, conn *grpc.ClientConn) (*AuthHandler, error) {
	client := auth.NewAuthServiceClient(conn)

	return &AuthHandler{
		logger: lgr,
		client: client,
	}, nil
}

func (a AuthHandler) Registration(ctx context.Context, registrationDTO dto.RegistrationDTO) error {
	pbRegisterRequest := &auth.RegisterRequest{
		Email:                registrationDTO.Login,
		Password:             registrationDTO.Password,
		PasswordConfirmation: registrationDTO.PasswordConfirmation,
	}

	_, err := a.client.Register(context.TODO(), pbRegisterRequest)
	if err != nil {
		return err
	}

	return nil
}

func (a AuthHandler) Login(ctx context.Context, loginDto dto.LoginDTO) (*dto.TokensPairDTO, error) {
	pbLoginRequest := &auth.LoginRequest{
		Email:    loginDto.Login,
		Password: loginDto.Password,
	}

	response, err := a.client.Login(ctx, pbLoginRequest)
	if err != nil {
		return nil, err
	}

	return &dto.TokensPairDTO{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}, nil
}

func (a AuthHandler) RefreshToken(ctx context.Context) (*dto.TokensPairDTO, error) {
	response, err := a.client.RefreshAccessToken(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &dto.TokensPairDTO{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
	}, nil
}

func (a AuthHandler) Logout(ctx context.Context, logoutType int) error {
	pbLogoutRequest := &auth.LogoutRequest{
		Type: auth.LogoutType(logoutType),
	}

	_, err := a.client.Logout(ctx, pbLogoutRequest)
	if err != nil {
		return err
	}

	return nil
}
