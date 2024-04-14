package primary

import (
	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/dto"
)

type UseCase interface {
	Login(loginDTO dto.LoginDTO) dto.TokensPairDTO
	Registration(registrationDTO dto.RegistrationDTO)
}
