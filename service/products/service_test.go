package products

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/krogertechnology/krogo/pkg/krogo/request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"practice-app/models"
	"practice-app/store/products"
	"practice-app/store/variants"
	"testing"
)

func TestHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductStore := products.NewMockProductStore(ctrl)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockService := New(mockProductStore, mockVariantStore)

	ctx := krogo.NewContext(nil, nil, krogo.New())

	testcases := []struct {
		Desc           string
		ExpectedResult *models.ProductWithVariants
		ExpectedErr    error
		ID             string
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			ID:   "1",
			ExpectedResult: &models.ProductWithVariants{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
				Variant: []models.VariantInfo{{
					ID:      "1",
					Name:    "variant_1",
					Details: "details",
				}},
			},
			ExpectedErr: nil,
			Calls: []*gomock.Call{
				mockProductStore.EXPECT().GetByID(ctx, "1").Return(&models.ProductWithVariants{
					ID:        "1",
					Name:      "product_1",
					BrandName: "brand_1",
					Details:   "details",
					ImageUrl:  "url",
				}, nil),
				mockVariantStore.EXPECT().GetVariantData(ctx, "1").Return([]models.VariantInfo{{
					ID:      "1",
					Name:    "variant_1",
					Details: "details",
				}}, nil),
			},
		},
		{
			Desc:           "Failure: entity not found",
			ID:             "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.EntityNotFound{ID: "1", Entity: "products"},
			Calls: []*gomock.Call{
				mockProductStore.EXPECT().GetByID(ctx, "1").Return(nil, sql.ErrNoRows),
			},
		},
		{
			Desc:           "Failure: DB error",
			ID:             "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("db error")},
			Calls: []*gomock.Call{
				mockProductStore.EXPECT().GetByID(ctx, "1").Return(nil, errors.DB{Err: errors.Error("db error")}),
			},
		},
	}

	for i, test := range testcases {
		res, err := mockService.GetByID(ctx, test.ID)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductStore := products.NewMockProductStore(ctrl)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockService := New(mockProductStore, mockVariantStore)

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
		ExpectedErr    error
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			ExpectedResult: []models.ProductWithVariants{{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
				Variant: []models.VariantInfo{{
					ID:      "1",
					Name:    "variant_name",
					Details: "detail",
				}},
			}},
			ExpectedErr: nil,
			Calls: []*gomock.Call{
				mockProductStore.EXPECT().GetAll(gomock.Any(), map[string]string{"pid": "1"}).Return([]models.ProductWithVariants{{
					ID:        "1",
					Name:      "product_1",
					BrandName: "brand_1",
					Details:   "details",
					ImageUrl:  "url",
					Variant: []models.VariantInfo{{
						ID:      "1",
						Name:    "variant_name",
						Details: "detail",
					}},
				}}, nil),
			},
		},
	}

	for i, test := range testcases {
		target := "/products?pid=1"
		r := httptest.NewRequest(http.MethodGet, target, nil)
		req := request.NewHTTPRequest(r)
		ctx := krogo.NewContext(nil, req, krogo.New())
		res, err := mockService.GetAll(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductStore := products.NewMockProductStore(ctrl)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockService := New(mockProductStore, mockVariantStore)

	testcases := []struct {
		Desc           string
		ExpectedResult *models.Product
		ExpectedErr    error
		Body           *models.Product
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			ExpectedResult: &models.Product{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			},
			ExpectedErr: nil,
			Body: &models.Product{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			},
			Calls: []*gomock.Call{
				mockProductStore.EXPECT().Create(gomock.Any(), &models.Product{
					ID:        "1",
					Name:      "product_1",
					BrandName: "brand_1",
					Details:   "details",
					ImageUrl:  "url",
				}).Return(&models.Product{
					ID:        "1",
					Name:      "product_1",
					BrandName: "brand_1",
					Details:   "details",
					ImageUrl:  "url",
				}, nil),
			},
		},
		{
			Desc:           "Failure missing params",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"id", "name", "brand_name", "details", "image_url"}},
			Body:           &models.Product{},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		ctx := krogo.NewContext(nil, nil, krogo.New())
		res, err := mockService.Create(ctx, test.Body)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
