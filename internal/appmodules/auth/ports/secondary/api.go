package secondary

import (
	"context"
	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/dto"
)

type ApiPort interface {
	Registration(context.Context, dto.RegistrationDTO) error
	Login(context.Context, dto.LoginDTO) (*dto.TokensPairDTO, error)
	RefreshToken(context.Context) (*dto.TokensPairDTO, error)
	Logout(context.Context, int) error
}
