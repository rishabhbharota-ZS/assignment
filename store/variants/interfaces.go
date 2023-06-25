package variants

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
)

type VariantStore interface {
	GetByID(ctx *krogo.Context, id, vID string) (*models.Variant, error)
	Create(ctx *krogo.Context, variant *models.Variant) (*models.Variant, error)
	GetVariantData(ctx *krogo.Context, productID string) ([]models.VariantInfo, error)
}
