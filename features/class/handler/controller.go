package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/class"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

type classHandler struct {
	service class.ClassService
}

func New(cs class.ClassService) class.ClassHandler {
	return &classHandler{
		service: cs,
	}
}

// RegisterClass implements class.ClassHandler
func (ch *classHandler) RegisterClass() echo.HandlerFunc {
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

// ListClasses implements class.ClassHandler
func (ch *classHandler) ListClasses() echo.HandlerFunc {
	return func(c echo.Context) error {
		limitStr := c.QueryParam("limit")
		offsetStr := c.QueryParam("offset")

		limit := 5
		if limitStr != "" {
			limitInt, err := strconv.Atoi(limitStr)
			if err != nil {
				c.Logger().Errorf("limit is not a number: %s", limitStr)
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil, nil))
			}
			limit = limitInt
		}

		offset := 0
		if offsetStr != "" {
			offsetInt, err := strconv.Atoi(offsetStr)
			if err != nil {
				c.Logger().Errorf("offset is not a number: %s", offsetStr)
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil, nil))
			}
			offset = offsetInt
		}

		classes, count, err := ch.service.ListClasses(limit, offset)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		result := make([]listClassesResponse, len(classes))
		for i, class := range classes {
			result[i] = listClasses(class)
			result[i].No = uint(i + 1)
		}

		pagination := helper.Paginate(limit, offset, int(count))
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation", result, pagination))
	}
}
