package products

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/krogertechnology/krogo/pkg/krogo/request"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"practice-app/models"
	"practice-app/service/products"
	"testing"
)

func getContext() *krogo.Context {
	r := httptest.NewRequest(http.MethodGet, "/products/{id}", nil)
	req := request.NewHTTPRequest(r)
	ctx := krogo.NewContext(nil, req, krogo.New())

	return ctx
}

func TestHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductService := products.NewMockProductService(ctrl)
	mockHandler := New(mockProductService)

	ctx := getContext()

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
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
			},
			ExpectedErr: nil,
			Calls: []*gomock.Call{
				mockProductService.EXPECT().GetByID(ctx, "1").Return(&models.ProductWithVariants{
					ID:        "1",
					Name:      "product_1",
					BrandName: "brand_1",
					Details:   "details",
					ImageUrl:  "url",
				}, nil),
			},
		},
		{
			Desc:           "Failure",
			ID:             "",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"id"}},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		ctx.SetPathParams(map[string]string{"id": test.ID})
		res, err := mockHandler.GetByID(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductService := products.NewMockProductService(ctrl)
	mockHandler := New(mockProductService)

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
		ExpectedErr    error
		Pid            string
		Vid            string
		Name           string
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			Pid:  "1",
			Vid:  "1",
			Name: "product_1",
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
				mockProductService.EXPECT().GetAll(gomock.Any()).Return([]models.ProductWithVariants{{
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
		{
			Desc:           "Failure: pid not provided",
			Pid:            "",
			Vid:            "1",
			Name:           "product_1",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"pid"}},
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "Failure: pid not valid",
			Pid:            "a",
			Vid:            "1",
			Name:           "product_1",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"pid"}},
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "Failure: vid not valid",
			Pid:            "1",
			Vid:            "a",
			Name:           "product_1",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"vid"}},
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "Failure: pid not valid",
			Pid:            "1",
			Vid:            "1",
			Name:           "p",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"name"}},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		target := "/products?pid=" + test.Pid + "&vid=" + test.Vid + "&name=" + test.Name
		r := httptest.NewRequest(http.MethodGet, target, nil)
		req := request.NewHTTPRequest(r)
		ctx := krogo.NewContext(nil, req, krogo.New())

		res, err := mockHandler.GetAll(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductService := products.NewMockProductService(ctrl)
	mockHandler := New(mockProductService)

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
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
				mockProductService.EXPECT().Create(gomock.Any(), &models.Product{
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
			Desc:           "bind error",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"body"}},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		body, _ := json.Marshal(test.Body)

		if test.Desc == "bind error" {
			body = []byte("invalid Body")
		}

		r := httptest.NewRequest(http.MethodGet, "/products", bytes.NewBuffer(body))
		req := request.NewHTTPRequest(r)
		ctx := krogo.NewContext(nil, req, krogo.New())

		res, err := mockHandler.Create(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
