package variants

import (
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
	"practice-app/service/variants"
)

type Handler struct {
	service variants.VariantService
}

func New(service variants.VariantService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetByID(ctx *krogo.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	pID := ctx.PathParam("pid")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if pID == "" {
		return nil, errors.MissingParam{Param: []string{"pid"}}
	}

	return h.service.GetByID(ctx, id, pID)
}

func (h *Handler) Create(ctx *krogo.Context) (interface{}, error) {
	var variant *models.Variant

	pID := ctx.PathParam("pid")

	if pID == "" {
		return nil, errors.MissingParam{Param: []string{"pid"}}
	}

	if err := ctx.Bind(&variant); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if pID != variant.ProductID {
		return nil, errors.InvalidParam{Param: []string{"pid"}}
	}

	return h.service.Create(ctx, variant)
}
