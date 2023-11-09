package test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/yaza-putu/online-test-dot/src/config"
	"github.com/yaza-putu/online-test-dot/src/core"
	"github.com/yaza-putu/online-test-dot/src/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type e2eTestSuite struct {
	suite.Suite
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.Require().NoError(utils.EnvTesting("/.."))
	s.Require().NoError(utils.DatabaseTesting())

	go core.HttpServerTesting()
}

func (s *e2eTestSuite) TestCreateToken() {
	reqStr := `{"email":"user@mail.com","password" : "Password1"}`

	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/api/token", config.Host().Port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}

	response, err := client.Do(req)
	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(http.StatusOK, response.StatusCode)
	assert.Contains(s.T(), strings.Trim(string(byteBody), "\n"), "access_token")

	response.Body.Close()
}

func (s *e2eTestSuite) TestWrongCredintial() {
	reqStr := `{"email":"user@mail.com","password" : "1"}`

	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/api/token", config.Host().Port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}

	response, err := client.Do(req)
	s.NoError(err)

	s.Equal(http.StatusUnauthorized, response.StatusCode)

	response.Body.Close()
}
