package handler

import (
	"net/http"
	"strings"

	"github.com/GroupProject2-Kelompok4/BE/features/feedback"
	"github.com/GroupProject2-Kelompok4/BE/utils/helper"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"github.com/GroupProject2-Kelompok4/BE/utils/storages"
	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	service feedback.FeedbackService
}

func New(fs feedback.FeedbackService) feedback.FeedbackHandler {
	return &feedbackHandler{
		service: fs,
	}
}

// RegisterFeedback implements feedback.FeedbackHandler
func (fh *feedbackHandler) RegisterFeedbackMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := RegisterFeedbackMenteeRequest{}
		userId, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT", nil, nil))
		}

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		var imageURL string
		file, err := c.FormFile("proof")
		if err == nil {
			imageURL, err = storages.UploadImage(c, file)
			if err != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to upload image", nil, nil))
			}
			request.Proof = imageURL
		}

		request.Proof = imageURL
		result, err := fh.service.RegisterFeedbackMentee(RequestToCore(request), userId)
		if err != nil {
			if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal server error", nil, nil))
			}
		}

		resp := registerFeedbackMentee(result)
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully created new feedback", resp, nil))
	}
}

// UpdateFeedbackMentee implements feedback.FeedbackHandler
func (fh *feedbackHandler) UpdateFeedbackMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		request := UpdateFeedbackMenteeRequest{}
		userId, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}

		feedbackId := c.Param("id")

		errBind := c.Bind(&request)
		if errBind != nil {
			c.Logger().Error("error on bind login input")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad request", nil, nil))
		}

		var imageURL string
		file, errReq := c.FormFile("proof")
		if errReq == nil {
			imageURL, errReq = storages.UploadImage(c, file)
			if errReq != nil {
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Failed to upload image", nil, nil))
			}
			request.Proof = &imageURL
		}

		request.Proof = &imageURL
		err := fh.service.UpdateFeedbackMentee(RequestToCore(&request), feedbackId, userId)
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

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Successfully updated feedback", nil, nil))
	}
}

// DeleteFeedbackMentee implements feedback.FeedbackHandler
func (fh *feedbackHandler) DeleteFeedbackMentee() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _, errToken := middlewares.ExtractToken(c)
		if errToken != nil {
			c.Logger().Error("missing or malformed JWT")
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT.", nil, nil))
		}

		feedbackId := c.Param("id")
		err := fh.service.DeleteFeedbackMentee(feedbackId, userId)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found", nil, nil))
			}
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil, nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Successfully deleted feedback", nil, nil))
	}
}
