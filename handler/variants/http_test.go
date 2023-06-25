package variants

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
	"practice-app/service/variants"
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
	mockVariantService := variants.NewMockVariantService(ctrl)
	mockHandler := New(mockVariantService)

	ctx := getContext()

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
				mockVariantService.EXPECT().GetByID(ctx, "1", "1").Return(&models.Variant{
					ID:        "1",
					ProductID: "1",
					Name:      "variant_1",
					Details:   "details",
				}, nil),
			},
		},
		{
			Desc:           "Failure: missing variant id",
			ID:             "",
			Pid:            "1",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"id"}},
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "Failure: missing product id",
			ID:             "1",
			Pid:            "",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"pid"}},
			Calls:          []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		ctx.SetPathParams(map[string]string{"id": test.ID, "pid": test.Pid})
		res, err := mockHandler.GetByID(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockVariantService := variants.NewMockVariantService(ctrl)
	mockHandler := New(mockVariantService)

	testcases := []struct {
		Desc           string
		ExpectedResult interface{}
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
				mockVariantService.EXPECT().Create(gomock.Any(), &models.Variant{
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
			Desc:           "Failure: product id not provided",
			ExpectedResult: nil,
			ExpectedErr:    errors.MissingParam{Param: []string{"pid"}},
			Pid:            "",
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "bind error",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"body"}},
			Pid:            "1",
			Calls:          []*gomock.Call{},
		},
		{
			Desc:           "pid invalid",
			ExpectedResult: nil,
			ExpectedErr:    errors.InvalidParam{Param: []string{"pid"}},
			Pid:            "1",
			Body: &models.Variant{
				ID:        "1",
				ProductID: "2",
				Name:      "variant_1",
				Details:   "details",
			},
			Calls: []*gomock.Call{},
		},
	}

	for i, test := range testcases {
		body, _ := json.Marshal(test.Body)

		if test.Desc == "bind error" {
			body = []byte("invalid Body")
		}

		target := "/products/" + test.Pid + "/variant"
		r := httptest.NewRequest(http.MethodGet, target, bytes.NewBuffer(body))
		req := request.NewHTTPRequest(r)
		ctx := krogo.NewContext(nil, req, krogo.New())
		ctx.SetPathParams(map[string]string{"pid": test.Pid})

		res, err := mockHandler.Create(ctx)

		assert.Equalf(t, test.ExpectedResult, res, "TEST[%v] FAILED - %s", i, test.Desc)
		assert.Equalf(t, test.ExpectedErr, err, "TEST[%v] FAILED - %s", i, test.Desc)
	}
}
