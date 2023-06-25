package variants

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
)

type VariantService interface {
	GetByID(ctx *krogo.Context, id, pID string) (*models.Variant, error)
	Create(ctx *krogo.Context, variant *models.Variant) (*models.Variant, error)
}
