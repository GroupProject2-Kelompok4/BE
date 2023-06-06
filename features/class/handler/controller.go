package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	service class.ClassService
}

func New(cs class.ClassService) class.ClassHandler {
	return &userHandler{
		service: cs,
	}
}

// RegisterClass implements class.ClassHandler
func (ch *userHandler) RegisterClass() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterClassRequest{}
		_, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT", nil, nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		result, pic, err := ch.service.RegisterClass(RequestToCore(request))
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
			}
			if strings.Contains(err.Error(), "duplicated") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		resp := registerClass(result, pic)
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created a class", resp, nil))
	}
}
