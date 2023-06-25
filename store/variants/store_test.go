package variants

import (
	"context"
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/krogertechnology/krogo/pkg/datastore"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/stretchr/testify/assert"
	"practice-app/models"
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
	s := New()

	testcases := []struct {
		Desc           string
		ID             string
		pID            string
		ExpectedResult *models.Variant
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedQuery
	}{
		{
			Desc: "Success",
			ID:   "1",
			pID:  "1",
			ExpectedResult: &models.Variant{
				ID:        "1",
				Name:      "variant_1",
				ProductID: "1",
				Details:   "details",
			},
			ExpectedErr: nil,
			MockCall: mock.ExpectQuery("SELECT").WithArgs("1", "1").WillReturnRows(
				sqlmock.NewRows([]string{"id", "product_id", "variant_name", "variant_details"}).
					AddRow("1", "1", "variant_1", "details")),
		},
		{
			Desc:           "sql no rows",
			ID:             "1",
			pID:            "1",
			ExpectedResult: nil,
			ExpectedErr:    sql.ErrNoRows,
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1", "1").WillReturnError(sql.ErrNoRows),
		},
		{
			Desc:           "db error",
			ID:             "1",
			pID:            "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1", "1").WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := s.GetByID(ctx, test.ID, test.pID)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func Test_Create(t *testing.T) {
	ctx, mock := getSqlMock(t)
	s := New()

	testcases := []struct {
		Desc           string
		ID             string
		Pid            string
		Body           *models.Variant
		ExpectedResult *models.Variant
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedExec
	}{
		{
			Desc: "Success",
			ID:   "1",
			Body: &models.Variant{
				ID:        "1",
				Name:      "variant_1",
				ProductID: "1",
				Details:   "details",
			},
			ExpectedResult: &models.Variant{
				ID:        "1",
				Name:      "variant_1",
				ProductID: "1",
				Details:   "details",
			},
			ExpectedErr: nil,
			MockCall:    mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil),
		},
		{
			Desc: "Failure: DB error",
			ID:   "1",
			Body: &models.Variant{
				ID:        "1",
				Name:      "variant_1",
				ProductID: "1",
				Details:   "details",
			},
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := s.Create(ctx, test.Body)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func Test_GetVariantData(t *testing.T) {
	ctx, mock := getSqlMock(t)
	s := New()

	testcases := []struct {
		Desc           string
		Pid            string
		ExpectedResult []models.VariantInfo
		ExpectedErr    error
		MockCall       *sqlmock.ExpectedQuery
	}{
		{
			Desc: "Success",
			Pid:  "1",
			ExpectedResult: []models.VariantInfo{{
				ID:      "1",
				Name:    "variant_1",
				Details: "details",
			}},
			ExpectedErr: nil,
			MockCall: mock.ExpectQuery("SELECT").WithArgs("1").WillReturnRows(
				sqlmock.NewRows([]string{"id", "variant_name", "variant_details"}).
					AddRow("1", "variant_1", "details")),
		},
		{
			Desc:           "Failure: No rows",
			Pid:            "1",
			ExpectedResult: nil,
			ExpectedErr:    nil,
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(sql.ErrNoRows),
		},
		{
			Desc:           "Failure: DB error",
			Pid:            "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.DB{Err: errors.Error("DB Error")},
			MockCall:       mock.ExpectQuery("SELECT").WithArgs("1").WillReturnError(errors.Error("DB Error")),
		},
	}

	for i, test := range testcases {
		res, err := s.GetVariantData(ctx, test.Pid)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
