package products

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
)

type ProductStore interface {
	GetByID(ctx *krogo.Context, id string) (*models.ProductWithVariants, error)
	GetAll(ctx *krogo.Context, params map[string]string) ([]models.ProductWithVariants, error)
	Create(ctx *krogo.Context, product *models.Product) (*models.Product, error)
}
