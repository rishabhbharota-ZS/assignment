package products

import (
	"context"
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/krogertechnology/krogo/pkg/datastore"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/stretchr/testify/assert"
	"practice-app/models"
	"practice-app/store/variants"
	"testing"
)

func getSqlMock(t *testing.T) (*krogo.Context, sqlmock.Sqlmock) {
	ctx := krogo.NewContext(nil, nil, krogo.New())
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Errorf("error while mocking db %v", err)
	}

	ctx.DataStore = datastore.DataStore{ORM: db}
	ctx.Context = context.Background()

	return ctx, mock
}

func Test_GetByID(t *testing.T) {
	ctx, mock := getSqlMock(t)

	ctrl := gomock.NewController(t)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockProductStore := New(mockVariantStore)

	testcases := []struct {
		Desc           string
		ID             string
		ExpectedResult *models.ProductWithVariants
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedQuery
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
			MockCall: mock.ExpectQuery("SELECT").WithArgs("1").WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "brand_name", "details", "image_url"}).
					AddRow("1", "product_1", "brand_1", "details", "url")),
		},
		{
			Desc:           "Failure: No rows",
			ID:             "1",
			ExpectedResult: nil,
			ExpectedErr:    sql.ErrNoRows,
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(sql.ErrNoRows),
		},
		{
			Desc:           "Failure: DB error",
			ID:             "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := mockProductStore.GetByID(ctx, test.ID)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func Test_GetAll(t *testing.T) {
	ctx, mock := getSqlMock(t)

	ctrl := gomock.NewController(t)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockProductStore := New(mockVariantStore)

	testcases := []struct {
		Desc           string
		Params         map[string]string
		ExpectedResult []models.ProductWithVariants
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedQuery
		Calls          []*gomock.Call
	}{
		{
			Desc:   "Success",
			Params: map[string]string{"pid": "1"},
			ExpectedResult: []models.ProductWithVariants{{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			}},
			ExpectedErr: nil,
			MockCall: mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "brand_name", "details", "image_url"}).
					AddRow("1", "product_1", "brand_1", "details", "url")),
			Calls: []*gomock.Call{
				mockVariantStore.EXPECT().GetVariantData(gomock.Any(), "1").Return(nil, nil),
			},
		},
		{
			Desc:   "Success",
			Params: map[string]string{"pid": "1", "vid": "1", "name": "product_1"},
			ExpectedResult: []models.ProductWithVariants{{
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
			}},
			ExpectedErr: nil,
			MockCall: mock.ExpectQuery("SELECT").WithArgs("1", "product_1").WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "brand_name", "details", "image_url"}).
					AddRow("1", "product_1", "brand_1", "details", "url")),
			Calls: []*gomock.Call{
				mockVariantStore.EXPECT().GetByID(gomock.Any(), "1", "1").Return(&models.Variant{
					ID:      "1",
					Name:    "variant_1",
					Details: "details",
				}, nil),
			},
		},
		{
			Desc:           "Failure: No rows",
			Params:         map[string]string{"pid": "1"},
			ExpectedResult: nil,
			ExpectedErr:    nil,
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(sql.ErrNoRows),
		},
		{
			Desc:           "Failure: DB error",
			Params:         map[string]string{"pid": "1"},
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := mockProductStore.GetAll(ctx, test.Params)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func Test_Create(t *testing.T) {
	ctx, mock := getSqlMock(t)

	ctrl := gomock.NewController(t)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockProductStore := New(mockVariantStore)

	testcases := []struct {
		Desc           string
		ID             string
		Body           *models.Product
		ExpectedResult *models.Product
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedExec
	}{
		{
			Desc: "Success",
			ID:   "1",
			Body: &models.Product{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			},
			ExpectedResult: &models.Product{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			},
			ExpectedErr: nil,
			MockCall:    mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil),
		},
		{
			Desc: "Failure: DB error",
			ID:   "1",
			Body: &models.Product{
				ID:        "1",
				Name:      "product_1",
				BrandName: "brand_1",
				Details:   "details",
				ImageUrl:  "url",
			},
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := mockProductStore.Create(ctx, test.Body)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
