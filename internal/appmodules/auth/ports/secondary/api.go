package secondary

import (
	"github.com/jbakhtin/goph-keeper/internal/client/appmodules/auth/domain/dto"
)

type ApiPort interface {
	Registration(dto.RegistrationDTO) (dto.TokensPairDTO, error)
	Login(dto.LoginDTO) (dto.TokensPairDTO, error)
	RefreshToken(string) (dto.TokensPairDTO, error)
	Logout() error
}
