package test

import (
	"bookmark-backend/type/response"
	"bookmark-backend/type/share"
	"bytes"
	"context"
	"encoding/json"
	"github.com/bsthun/gut"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"strconv"
	"testing"
)

type RequestResult[T any] struct {
	Success      *bool
	ErrorMessage *string
	Message      *string
	Data         T
}

func ExecuteTest[T any](t *testing.T, method string, handler fiber.Handler, body any) *RequestResult[*T] {
	return ExecuteTestWithContext[T](t, method, handler, body, nil)
}

func ExecuteTestWithContext[T any](t *testing.T, method string, handler fiber.Handler, body any, ctx context.Context) *RequestResult[*T] {
	// * create fiber app
	app := fiber.New()

	// * register handler
	app.Use(func(c *fiber.Ctx) error {
		if ctx != nil && ctx.Value("userId") != nil {
			c.Locals("l", &jwt.Token{
				Claims: &share.UserClaims{
					UserId: gut.Ptr(strconv.FormatUint(ctx.Value("userId").(uint64), 10)),
				},
			})
		}
		return c.Next()
	})
	app.Add(method, "/", handler)

	var reqBody io.Reader
	if body != nil {
		// * marshal body to json
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// * create request
	req, err := http.NewRequest(method, "/", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// * execute request
	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	// * read response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode >= 400 {
		// * parse error response
		var resp response.ErrorResponse
		if err := json.Unmarshal(resBody, &resp); err != nil {
			t.Fatal(err)
		}

		return &RequestResult[*T]{
			Success:      resp.Success,
			ErrorMessage: resp.Message,
		}
	}

	// * parse response
	var resp response.GenericResponse[*T]
	if err := json.Unmarshal(resBody, &resp); err != nil {
		t.Fatal(err)
	}

	return &RequestResult[*T]{
		Success: resp.Success,
		Message: resp.Message,
		Data:    resp.Data,
	}
}
