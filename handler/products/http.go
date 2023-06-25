package products

import (
	"github.com/krogertechnology/krogo/pkg/errors"
	"github.com/krogertechnology/krogo/pkg/krogo"
	"practice-app/models"
	"practice-app/service/products"
	"regexp"
	"strconv"
)

type Handler struct {
	service products.ProductService
}

func New(service products.ProductService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetByID(ctx *krogo.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	return h.service.GetByID(ctx, id)
}

func (h *Handler) GetAll(ctx *krogo.Context) (interface{}, error) {
	id := ctx.Param("pid")
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"pid"}}
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"pid"}}
	}

	vid := ctx.Param("vid")
	_, err = strconv.Atoi(vid)
	if vid != "" && err != nil {
		return nil, errors.InvalidParam{Param: []string{"vid"}}
	}

	re, _ := regexp.Compile("^[A-Za-z][A-Za-z0-9_]{2,255}$")
	name := ctx.Param("name")
	if name != "" && !re.MatchString(name) {
		return nil, errors.InvalidParam{Param: []string{"name"}}
	}

	return h.service.GetAll(ctx)
}

func (h *Handler) Create(ctx *krogo.Context) (interface{}, error) {
	var product *models.Product

	if err := ctx.Bind(&product); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	return h.service.Create(ctx, product)
}
