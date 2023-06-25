package variants

import (
	"github.com/golang/mock/gomock"
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"github.com/stretchr/testify/assert"
	"practice-app/models"
	"practice-app/store/variants"
	"testing"
)

func TestHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockService := New(mockVariantStore)

	ctx := krogo.NewContext(nil, nil, krogo.New())

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
		ExpectedErr    error
		ID             string
		Pid            string
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			ID:   "1",
			Pid:  "1",
			ExpectedResult: &models.Variant{
				ID:        "1",
				ProductID: "1",
				Name:      "variant_1",
				Details:   "details",
			},
			ExpectedErr: nil,
			Calls: []*gomock.Call{
				mockVariantStore.EXPECT().GetByID(gomock.Any(), "1", "1").Return(&models.Variant{
					ID:        "1",
					ProductID: "1",
					Name:      "variant_1",
					Details:   "details",
				}, nil),
			},
		},
	}

	for i, test := range testcases {
		res, err := mockService.GetByID(ctx, test.ID, test.Pid)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVariantStore := variants.NewMockVariantStore(ctrl)
	mockService := New(mockVariantStore)

	ctx := krogo.NewContext(nil, nil, krogo.New())

	testcases := []struct {
		Desc           string
		ExpectedResult *models.Variant
		ExpectedErr    error
		Pid            string
		Body           *models.Variant
		Calls          []*gomock.Call
	}{
		{
			Desc: "Success",
			ExpectedResult: &models.Variant{
				ID:        "1",
				ProductID: "1",
				Name:      "variant_1",
				Details:   "details",
			},
			ExpectedErr: nil,
			Pid:         "1",
			Body: &models.Variant{
				ID:        "1",
				ProductID: "1",
				Name:      "variant_1",
				Details:   "details",
			},
			Calls: []*gomock.Call{
				mockVariantStore.EXPECT().Create(gomock.Any(), &models.Variant{
					ID:        "1",
					ProductID: "1",
					Name:      "variant_1",
					Details:   "details",
				}).Return(&models.Variant{
					ID:        "1",
					ProductID: "1",
					Name:      "variant_1",
					Details:   "details",
				}, nil),
			},
		},
		{
			Desc:           "Failure: Missing params",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"id", "product_id", "name", "details"}},
			Pid:            "",
			Body:           &models.Variant{},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		res, err := mockService.Create(ctx, test.Body)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
