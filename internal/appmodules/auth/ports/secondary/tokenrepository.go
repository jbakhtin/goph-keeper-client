package secondary

import (
	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/dto"
)

type TokenRepository interface {
	SaveTokenPair(*dto.TokensPairDTO) error
	GetTokensPair() (*dto.TokensPairDTO, error)
	RemoveTokensPair() error
}
