//go:build e2e
// +build e2e

package test

import (
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

const (
	BASE_URL = "http://localhost:8080"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestGetComments(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		// SetHeader("auth", "bearer"+createToken())
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/v1/comment")
	assert.NoError(t, err)
	// assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 401, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "12345", "body": "hello world"}`).
		Post(BASE_URL + "/api/v1/comment")
	assert.NoError(t, err)

	assert.Equal(t, 401, resp.StatusCode())
}
