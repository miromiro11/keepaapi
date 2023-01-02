package api

import "errors"

var (
	ErrInvalidApiKey = errors.New("invalid api key")
	ErrInvalidProductId = errors.New("invalid product id")
	ErrInvalidDomain = errors.New("invalid domain")
	ErrInvalidCategory = errors.New("invalid category")
	ErrInvalidGraph = errors.New("invalid graph")
)