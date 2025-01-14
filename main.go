package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// SquareOutput represents the square operation response.
type SquareOutput struct {
	Body struct {
		Number int `json:"number" example:"4" doc:"Input number"`
		Square int `json:"square" example:"16" doc:"Square of the input number"`
	} `json:"body"`
}

func main() {
	// Create a new router & API.
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("Square Calculator API", "1.0.0"))

	// Register GET /square/{number} handler.
	huma.Register(api, huma.Operation{
		OperationID: "get-square",
		Method:      http.MethodGet,
		Path:        "/square/{number}",
		Summary:     "Calculate the square",
		Description: "Calculate the square of a given number.",
		Tags:        []string{"Math"},
	}, func(ctx context.Context, input *struct {
		Number string `path:"number" example:"4" doc:"The number to square"`
	}) (*SquareOutput, error) {
		// Parse the number from the input path
		num, err := strconv.Atoi(input.Number)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", input.Number)
		}

		// Calculate the square
		resp := &SquareOutput{}
		resp.Body.Number = num
		resp.Body.Square = num * num
		return resp, nil
	})

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
