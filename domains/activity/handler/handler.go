package activityhandler

import (
	"devcode/config"
	activitycore "devcode/domains/activity/core"
	"devcode/exceptions"
	"devcode/utils/helpers"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type activityHandler struct {
	service activitycore.IServiceActivity
}

var validate = validator.New()

func New(service activitycore.IServiceActivity) *activityHandler {
	return &activityHandler{
		service: service,
	}
}

func (h *activityHandler) Create(c *fiber.Ctx) error {
	request := Request{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	result := h.service.Create(activitycore.Core{
		Title: request.Title,
		Email: request.Email,
	})

	return c.Status(http.StatusCreated).JSON(helpers.SuccessGetResponseData(Response{
		Id:        result.Id,
		Title:     result.Title,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Format(config.LAYOUT_TIME),
		UpdatedAt: result.UpdatedAt.Format(config.LAYOUT_TIME),
		DeleteAt:  result.DeletedAt.Format(config.LAYOUT_TIME),
	}))
}

func (h *activityHandler) Update(c *fiber.Ctx) error {
	request := Request{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	result := h.service.Update(activitycore.Core{
		Id:    uint(id),
		Title: request.Title,
		Email: request.Email,
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(Response{
		Id:        result.Id,
		Title:     result.Title,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Format(config.LAYOUT_TIME),
		UpdatedAt: result.UpdatedAt.Format(config.LAYOUT_TIME),
		DeleteAt:  result.DeletedAt.Format(config.LAYOUT_TIME),
	}))
}

func (h *activityHandler) Delete(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	h.service.Delete(activitycore.Core{
		Id: uint(id),
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(map[string]interface{}{}))
}

func (h *activityHandler) FindAll(c *fiber.Ctx) error {
	results := h.service.FindAll()

	response := []Response{}

	for _, data := range results {
		response = append(response, Response{
			Id:        data.Id,
			Title:     data.Title,
			Email:     data.Email,
			CreatedAt: data.CreatedAt.Format(config.LAYOUT_TIME),
			UpdatedAt: data.UpdatedAt.Format(config.LAYOUT_TIME),
			DeleteAt:  data.DeletedAt.Format(config.LAYOUT_TIME),
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(response))
}

func (h *activityHandler) FindSingle(c *fiber.Ctx) error {
	paramId := c.Params("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	result := h.service.FindSingle(activitycore.Core{
		Id: uint(id),
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(Response{
		Id:        result.Id,
		Title:     result.Title,
		Email:     result.Email,
		CreatedAt: result.CreatedAt.Format(config.LAYOUT_TIME),
		UpdatedAt: result.UpdatedAt.Format(config.LAYOUT_TIME),
		DeleteAt:  result.DeletedAt.Format(config.LAYOUT_TIME),
	}))
}
