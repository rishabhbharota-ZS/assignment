package products

import (
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
)

type ProductService interface {
	GetByID(ctx *krogo.Context, id string) (*models.ProductWithVariants, error)
	GetAll(ctx *krogo.Context) ([]models.ProductWithVariants, error)
	Create(ctx *krogo.Context, product *models.Product) (*models.Product, error)
}
