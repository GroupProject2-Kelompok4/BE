package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/mentee"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
)

type menteeHandler struct {
	service mentee.MenteeService
}

func New(cs mentee.MenteeService) mentee.MenteeHandler {
	return &menteeHandler{
		service: cs,
	}
}

// RegisterMentee implements mentee.MenteeHandler
func (mh *menteeHandler) RegisterMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterMenteeRequest{}
		_, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT", nil, nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		result, err := mh.service.RegisterMentee(RequestToCore(request))
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

		resp := registerMentee(result)
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created new mentee", resp, nil))
	}
}

// SearchMentee implements mentee.MenteeHandler
func (mh *menteeHandler) SearchMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("keyword")
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

		mentees, count, err := mh.service.SearchMentee(keyword, limit, offset)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		result := make([]searchMenteeResponse, len(mentees))
		for i, mentee := range mentees {
			result[i] = searchMentee(mentee)
			result[i].No = uint(i + 1)
		}

		pagination := helper.Paginate(limit, offset, int(count))
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation", result, pagination))
	}
}

// ProfileMenteeAndFeedback implements mentee.MenteeHandler
func (mh *menteeHandler) ProfileMenteeAndFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, _, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}

		menteeId := c.Param("id")
		menteeLog, err := mh.service.ProfileMenteeAndFeedback(menteeId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
		}

		resp := profileMenteeAndFeedback(menteeLog)
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Successfully operation.", resp, nil))
	}
}

// UpdateMentee implements mentee.MenteeHandler
func (mh *menteeHandler) UpdateMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := UpdateMenteeRequest{}
		_, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}

		menteeId := c.Param("id")

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		err := mh.service.UpdateMentee(menteeId, RequestToCore(request))
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

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully updated mentee profile", nil, nil))

	}
}
