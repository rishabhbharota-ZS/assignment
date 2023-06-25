package products

import (
	"database/sql"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
	"practice-app/store/products"
	"practice-app/store/variants"
)

type Service struct {
	store        products.ProductStore
	variantStore variants.VariantStore
}

func New(store products.ProductStore, variantStore variants.VariantStore) *Service {
	return &Service{store: store, variantStore: variantStore}
}

func (s *Service) GetByID(ctx *krogo.Context, id string) (*models.ProductWithVariants, error) {
	p, err := s.store.GetByID(ctx, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.EntityNotFound{ID: id, Entity: "products"}
		}

		return nil, err
	}

	res, _ := s.variantStore.GetVariantData(ctx, id)

	p.Variant = res

	return p, nil
}

func (s *Service) GetAll(ctx *krogo.Context) ([]models.ProductWithVariants, error) {
	return s.store.GetAll(ctx, ctx.Params())
}

func (s *Service) Create(ctx *krogo.Context, product *models.Product) (*models.Product, error) {
	missingAttributes := findMissingAttributes(product)

	if len(missingAttributes) > 0 {
		return nil, errors.MissingParam{Param: missingAttributes}
	}

	return s.store.Create(ctx, product)
}

func findMissingAttributes(product *models.Product) (res []string) {
	if product.ID == "" {
		res = append(res, "id")
	}

	if product.Name == "" {
		res = append(res, "name")
	}

	if product.BrandName == "" {
		res = append(res, "brand_name")
	}

	if product.Details == "" {
		res = append(res, "details")
	}

	if product.ImageUrl == "" {
		res = append(res, "image_url")
	}

	return res
}
