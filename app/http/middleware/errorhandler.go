package middleware

import (
	"bytes"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jadevelopmentgrp/Tickets-Import-API/app"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type copyWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func (cw copyWriter) Write(b []byte) (int, error) {
	return cw.buf.Write(b)
}

func ErrorHandler(c *gin.Context) {
	cw := &copyWriter{buf: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = cw

	c.Next()

	if len(c.Errors) > 0 {
		var message string

		var apiError *app.ApiError
		if errors.As(c.Errors[0], &apiError) {
			message = apiError.ExternalMessage
		} else {
			message = "An error occurred processing your request"
		}

		c.Writer = cw.ResponseWriter
		c.JSON(-1, ErrorResponse{
			Error: message,
		})

		return
	}

	if c.Writer.Status() >= 500 {
		c.Writer = cw.ResponseWriter

		c.JSON(-1, ErrorResponse{
			Error: "An internal server error occurred",
		})

		return
	}

	cw.ResponseWriter.Write(cw.buf.Bytes())
}
