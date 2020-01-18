package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/saman2000hoseini/http-monitor/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSuccessfulRegister(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"username":"mammad","password":"mySecret"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/user/register", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, handler.Register(c))
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		assert.Contains(t, rec.Body.String(), "token")
	}
}

func TestFailedRegister(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"username":"ali","password":"Secret"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/user/register", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, handler.Register(c))
	fmt.Println(rec.Body.String())
	if assert.Equal(t, http.StatusConflict, rec.Code) {
		assert.Contains(t, rec.Body.String(), "min")
	}
}

func TestSuccessfulLogin(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"username":"user1","password":"secretpass1"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/user/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, handler.Login(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		assert.Contains(t, rec.Body.String(), "token")
	}
}

func TestFailedLogin(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user":{"username":"user2","password":"secretpass2"}}`
	)
	req := httptest.NewRequest(echo.POST, "/api/user/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, handler.Login(c))
	if assert.Equal(t, http.StatusForbidden, rec.Code) {
		fmt.Println(rec.Body.String())
	}
}

func TestSuccessfulUpdate(t *testing.T) {
	tearDown()
	setup()
	f := make(url.Values)
	f.Set("username", "mynewusername")
	f.Set("password", "password")
	req := httptest.NewRequest(echo.POST, "/api/user/update", strings.NewReader(f.Encode()))
	token, _ := utils.GenerateJWT(1)
	req.Header.Set(echo.HeaderAuthorization, "token: "+token)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		fmt.Println(rec.Body.String())
	}
}
