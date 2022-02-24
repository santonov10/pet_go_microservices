package handlers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type errorField struct {
	Message string     `json:"message,omitempty"`
	Code    codes.Code `json:"code,omitempty"`
}

type responseBody struct {
	Error  *errorField `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

type Response struct {
	HTTPStatus int
	Body       responseBody
}

func NewResponseOk(data interface{}) *Response {
	return &Response{
		HTTPStatus: http.StatusOK,
		Body: responseBody{
			Error:  nil,
			Result: data,
		},
	}
}

func NewResponseError(httpStatus int, err error) *Response {
	status.FromError(err)
	var errField *errorField
	st, ok := status.FromError(err)
	if !ok {
		errField = &errorField{Message: err.Error(), Code: codes.Unknown}
	} else {
		errField = &errorField{Message: st.Message(), Code: st.Code()}
	}

	return &Response{
		HTTPStatus: httpStatus,
		Body: responseBody{
			Error:  errField,
			Result: nil,
		},
	}
}

func (r *Response) GinJson(ginC *gin.Context) {
	ginC.JSON(r.HTTPStatus, r.Body)
}
