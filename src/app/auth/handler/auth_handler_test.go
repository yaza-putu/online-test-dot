package handler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yaza-putu/online-test-dot/src/utils"
	"net/http"
	"testing"
)

var handler = NewAuthHandler()

func TestCreateSuccessToken(t *testing.T) {
	p := `{"email" : "user@mail.com", "password" : "Password1"}`
	ctx, res := utils.NewServerTest("POST", "/token", p)
	fmt.Println(res.Body.String())
	if assert.NoError(t, handler.Create(ctx)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}
}
