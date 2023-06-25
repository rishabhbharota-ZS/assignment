package variants

import (
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
	"practice-app/store/variants"
)

type Service struct {
	store variants.VariantStore
}

func New(store variants.VariantStore) *Service {
	return &Service{store: store}
}

func (s *Service) GetByID(ctx *krogo.Context, id, pID string) (*models.Variant, error) {
	return s.store.GetByID(ctx, id, pID)
}

func (s *Service) Create(ctx *krogo.Context, variant *models.Variant) (*models.Variant, error) {
	missingAttributes := findMissingAttributes(variant)

	if len(missingAttributes) > 0 {
		return nil, errors.MissingParam{Param: missingAttributes}
	}

	return s.store.Create(ctx, variant)
}

func findMissingAttributes(variant *models.Variant) (res []string) {
	if variant.ID == "" {
		res = append(res, "id")
	}

	if variant.ProductID == "" {
		res = append(res, "product_id")
	}

	if variant.Name == "" {
		res = append(res, "name")
	}

	if variant.Details == "" {
		res = append(res, "details")
	}

	return res
}
