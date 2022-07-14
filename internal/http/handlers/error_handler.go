package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	TargetCommon = "common"
	TargetField  = "field"
)

type HTTPError struct {
	Code   string `json:"code"`
	Target string `json:"target"`
	Title  string `json:"title,omitempty"`
	Source string `json:"source,omitempty"`
}

type ErrorHandler interface {
	HandleError(err error, c *gin.Context)
}
type RequestParameterError string

func (e RequestParameterError) Error() string {
	return string(e)
}

func HandleError(err error, c *gin.Context) {
	var result *HTTPError
	switch err := err.(type) {
	case *json.UnmarshalTypeError:
		c.JSON(http.StatusBadRequest, gin.H{"errors": []*HTTPError{{
			Code:   "UNMARSHAL_JSON",
			Target: TargetField,
			Source: err.Field,
		}}})
		return
	case RequestParameterError:
		c.JSON(http.StatusBadRequest, gin.H{"errors": []*HTTPError{{
			Code:   "INVALID_REQUEST_PARAMETER",
			Target: TargetCommon,
			Source: err.Error(),
		}}})
		return
	}

	if errors.Cause(err).Error() == "context canceled" {
		return
	}

	status := 0
	switch errors.Cause(err) {
	case io.EOF:
		status = http.StatusBadRequest
		result = &HTTPError{
			Code:   "MISSING_REQUEST_BODY",
			Target: TargetCommon,
		}
	case io.ErrUnexpectedEOF:
		status = http.StatusBadRequest
		result = &HTTPError{
			Code:   "MALFORMED_REQUEST_BODY",
			Target: TargetCommon,
		}
	default:
		status = http.StatusInternalServerError
		result = &HTTPError{
			Code:   "INTERNAL_SERVER_ERROR",
			Target: TargetCommon,
		}
	}

	log.WithFields(log.Fields{
		"endpoint": c.Request.Method + " " + c.Request.URL.String(),
		"error":    err.Error(),
	})
	switch status {
	case http.StatusInternalServerError:
		log.Error("unexpected error")
	default:
		log.Info("handled error")
	}

	c.JSON(status, gin.H{"errors": []*HTTPError{result}})
}
