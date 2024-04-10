package secondary

import (
	"context"
	"github.com/jbakhtin/goph-keeper-client/internal/appmodules/auth/domain/models"
	"github.com/jbakhtin/goph-keeper-client/pkg/queryspec"
)

type SessionQuerySpecification interface {
	Limit(queryspec.QuerySpecification, int) queryspec.QuerySpecification
	Where(queryspec.QuerySpecification) queryspec.QuerySpecification
	Or(...queryspec.QuerySpecification) queryspec.QuerySpecification
	And(...queryspec.QuerySpecification) queryspec.QuerySpecification
	UserID(int) queryspec.QuerySpecification
}

type SessionRepository interface {
	Get(ctx context.Context, id int) (models.Session, error)
	Create(ctx context.Context, user models.Session) (models.Session, error)
	Update(ctx context.Context, user models.Session) (models.Session, error)
	Delete(ctx context.Context, user models.Session) (models.Session, error)
	Search(context.Context, queryspec.QuerySpecification) ([]models.Session, error)
}
