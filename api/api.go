package api

import (
	"encoding/json"
	"net/http"
)

// Coins Balance Params
type CoinsBalanceParams struct {
	username string
}

// Coins Balance Response
type CoinsBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct{
	// Error code
	Code int

	// Error message
	Message string
}