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
		result, err := fh.service.RegisterFeedbackMentee(RequestToCore(request))
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
