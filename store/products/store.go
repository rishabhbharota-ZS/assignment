package products

import (
	"database/sql"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
	"practice-app/store/variants"
	"strconv"
)

type Store struct {
	variantStore variants.VariantStore
}

func New(variantStore variants.VariantStore) *Store {
	return &Store{variantStore: variantStore}
}

func (s *Store) GetByID(ctx *krogo.Context, id string) (*models.ProductWithVariants, error) {
	query := "SELECT id, name, brand_name, details, image_url FROM products where id=$1"

	var p models.ProductWithVariants

	err := ctx.DB().QueryRowContext(ctx, query, id).
		Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}

		return nil, errors.DB{Err: err}
	}

	return &p, nil
}

func (s *Store) GetAll(ctx *krogo.Context, params map[string]string) ([]models.ProductWithVariants, error) {
	query := "SELECT id, name, brand_name, details, image_url FROM products "

	whereClause, values := generateWhereClause(params)

	var productArray []models.ProductWithVariants

	rows, err := ctx.DB().QueryContext(ctx, query+whereClause, values...)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	for rows.Next() {
		var p models.ProductWithVariants

		err = rows.Scan(&p.ID, &p.Name, &p.BrandName, &p.Details, &p.ImageUrl)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		var variantInfo []models.VariantInfo

		vid, ok := params["vid"]

		if ok {
			variant, err := s.variantStore.GetByID(ctx, p.ID, vid)

			if err != nil {
				return nil, err
			}

			var varInfo = models.VariantInfo{
				ID:      variant.ID,
				Name:    variant.Name,
				Details: variant.Details,
			}

			variantInfo = append(variantInfo, varInfo)
		} else {
			variantInfo, _ = s.variantStore.GetVariantData(ctx, p.ID)
		}

		p.Variant = variantInfo

		productArray = append(productArray, p)
	}

	return productArray, nil
}

func (s *Store) Create(ctx *krogo.Context, product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products(id, name, brand_name, details, image_url) VALUES ($1,$2,$3,$4,$5)"

	_, err := ctx.DB().ExecContext(ctx, query, product.ID, product.Name, product.BrandName, product.Details, product.ImageUrl)

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return product, nil
}

func generateWhereClause(params map[string]string) (string, []interface{}) {
	clause := "WHERE "
	i := 1
	and := ""

	var values []interface{}

	for key, value := range params {
		if key != "vid" {
			values = append(values, value)
			column := ""
			if key == "pid" {
				column = "id"
			}
			if key == "name" {
				column = "name"
			}
			if i > 1 {
				and = " AND "
			}
			clause += and + column + "=$" + strconv.Itoa(i)
			i++
		}
	}

	return clause, values
}
