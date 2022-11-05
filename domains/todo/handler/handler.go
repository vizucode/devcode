package todohandler

import (
	todocore "devcode/domains/todo/core"
	"devcode/exceptions"
	"devcode/utils/helpers"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type todoHandler struct {
	service todocore.IServiceTodo
}

func New(service todocore.IServiceTodo) *todoHandler {
	return &todoHandler{
		service: service,
	}
}

func (h *todoHandler) Create(ctx *fiber.Ctx) error {
	request := Request{}
	err := ctx.BodyParser(&request)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	err = validate.Struct(&request)
	if err != nil {
		panic(err)
	}

	result := h.service.Create(todocore.Core{
		ActivityGroupId: request.ActivityGroupId,
		Title:           request.Title,
	})

	return ctx.Status(http.StatusCreated).JSON(helpers.SuccessGetResponseData(ToResponse(result)))
}

func (h *todoHandler) Update(ctx *fiber.Ctx) error {
	request := RequestUpdate{}
	err := ctx.BodyParser(&request)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	err = validate.Struct(&request)
	if err != nil {
		panic(err)
	}

	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	result := h.service.Update(todocore.Core{
		Id:       uint(id),
		Title:    request.Title,
		IsActive: request.IsActive,
	})

	return ctx.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(ToResponse(result)))
}

func (h *todoHandler) Delete(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	h.service.Delete(todocore.Core{
		Id: uint(id),
	})

	return ctx.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(map[string]interface{}{}))
}

func (h *todoHandler) FindAll(ctx *fiber.Ctx) error {
	idParam := ctx.Query("activity_group_id")
	core := todocore.Core{
		ActivityGroupId: 0,
	}
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			panic(exceptions.NewBadRequestError(err.Error()))
		}
		core.ActivityGroupId = uint(id)
	}

	results := h.service.FindAll(core)

	response := []Response{}

	for _, data := range results {
		response = append(response, ToResponse(data))
	}

	return ctx.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(response))
}

func (h *todoHandler) FindById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(exceptions.NewBadRequestError(err.Error()))
	}

	result := h.service.FindById(todocore.Core{
		Id: uint(id),
	})

	return ctx.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(ToResponse(result)))
}
