package variants

import (
	"database/sql"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
)

type Store struct {
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetByID(ctx *krogo.Context, id, pID string) (*models.Variant, error) {
	query := "SELECT id, product_id, variant_name, variant_details FROM variants WHERE id=$1 AND product_id=$2"

	var v models.Variant

	err := ctx.DB().QueryRowContext(ctx, query, id, pID).
		Scan(&v.ID, &v.ProductID, &v.Name, &v.Details)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}

		return nil, errors.DB{Err: err}
	}

	return &v, nil
}

func (s *Store) Create(ctx *krogo.Context, variant *models.Variant) (*models.Variant, error) {
	query := "INSERT INTO variants(id, product_id, variant_name, variant_details) VALUES ($1,$2,$3,$4)"

	_, err := ctx.DB().ExecContext(ctx, query, variant.ID, variant.ProductID, variant.Name, variant.Details)

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return variant, nil
}

func (s *Store) GetVariantData(ctx *krogo.Context, productID string) ([]models.VariantInfo, error) {
	query := "SELECT id, variant_name, variant_details FROM variants WHERE product_id=$1"

	var variantInfo []models.VariantInfo

	rows, err := ctx.DB().QueryContext(ctx, query, productID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var v models.VariantInfo

		err = rows.Scan(&v.ID, &v.Name, &v.Details)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		variantInfo = append(variantInfo, v)
	}

	return variantInfo, nil
}
