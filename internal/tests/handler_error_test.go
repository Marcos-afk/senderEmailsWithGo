package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"senderEmails/internal/endpoints"
	internalerrors "senderEmails/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_When_Returning_InternalError(t *testing.T){
  assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request)(interface{}, int, error){
		return nil, 0, internalerrors.ErrInternal
	}

	handlerFunc := endpoints.HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())
}


func Test_HandlerError_When_Returning_BadRequest(t *testing.T){
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request)(interface{}, int, error){
		return nil, 0, errors.New("Bad Request")
	}

	handlerFunc := endpoints.HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Bad Request")
}


func Test_HandlerError_When_Returning_Object(t *testing.T){
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request)(interface{}, int, error){
		return struct {
			Name string `json:"name"`
		}{
			Name: "Test",
		}, 200, nil
	}

	handlerFunc := endpoints.HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Contains(res.Body.String(), "Test")
}

func Test_HandlerError_When_Returning_NoContent(t *testing.T){
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request)(interface{}, int, error){
		return nil, 0, nil
	}

	handlerFunc := endpoints.HandlerError(endpoint)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusNoContent, res.Code)
	assert.Empty(res.Body.String())
}