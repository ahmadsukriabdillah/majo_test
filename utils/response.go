package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type BaseResponse struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"errors"`
}

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

func NewResponseOk(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
		Code:    http.StatusOK,
		Data:    data,
		Message: "OK",
	})
}

var verr validator.ValidationErrors

func NewResponseError(ctx *gin.Context, err error) {
	if errors.As(err, &verr) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "error: validation error.",
			Error:   Descriptive(verr),
		})
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

}
