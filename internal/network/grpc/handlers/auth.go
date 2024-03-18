package handlers

import (
	"github.com/jbakhtin/goph-keeper/internal/client/appmodules/auth/domain/dto"
	secondary_ports "github.com/jbakhtin/goph-keeper/internal/client/appmodules/auth/ports/secondary"
	"github.com/jbakhtin/goph-keeper/internal/client/logger/zap"
)

var _ secondary_ports.ApiPort = &AuthHandler{}

type AuthHandler struct {
	lgr *zap.Logger
}

func (a AuthHandler) Registration(dto dto.RegistrationDTO) (dto.TokensPairDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) Login(dto dto.LoginDTO) (dto.TokensPairDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) RefreshToken(s string) (dto.TokensPairDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) Logout() error {
	//TODO implement me
	panic("implement me")
}
