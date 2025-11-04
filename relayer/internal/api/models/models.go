package models

import (
	"time"
)

// HealthStatus represents system health status
type HealthStatus struct {
	Status     string           `json:"status"`
	Timestamp  time.Time        `json:"timestamp"`
	Components HealthComponents `json:"components"`
}

type HealthComponents struct {
	Database     string `json:"database"`
	EthereumNode string `json:"ethereumNode"`
	PolygonNode  string `json:"polygonNode"`
	MessageQueue string `json:"messageQueue"`
}

// TransactionStatus represents cross-chain transaction status
type TransactionStatus struct {
	TransactionHash string    `json:"transactionHash"`
	Status          string    `json:"status"`
	SourceNetwork   string    `json:"sourceNetwork"`
	TargetNetwork   string    `json:"targetNetwork"`
	Amount          string    `json:"amount"`
	Timestamp       time.Time `json:"timestamp"`
	Confirmations   int       `json:"confirmations"`
	BridgeFee       string    `json:"bridgeFee,omitempty"`
}

// TransactionList represents paginated transaction list
type TransactionList struct {
	Transactions []TransactionStatus `json:"transactions"`
	Pagination   Pagination          `json:"pagination"`
}

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

// SystemStatus represents detailed system status
type SystemStatus struct {
	Overall    string                     `json:"overall"`
	Components map[string]ComponentStatus `json:"components"`
}

type ComponentStatus struct {
	Status    string    `json:"status"`
	Latency   int       `json:"latency,omitempty"`
	LastCheck time.Time `json:"lastCheck"`
}

// ErrorResponse represents API error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Details interface{} `json:"details,omitempty"`
}

// LockRequest represents token lock request
type LockRequest struct {
	SourceNetwork string `json:"sourceNetwork" binding:"required,oneof=ethereum polygon"`
	TargetNetwork string `json:"targetNetwork" binding:"required,oneof=ethereum polygon"`
	Amount        string `json:"amount" binding:"required,numeric"`
	TargetAddress string `json:"targetAddress" binding:"required,eth_addr"`
	TokenAddress  string `json:"tokenAddress,omitempty" binding:"omitempty,eth_addr"`
}

// LockResponse represents lock operation response
type LockResponse struct {
	TransactionHash         string    `json:"transactionHash"`
	Status                  string    `json:"status"`
	EstimatedCompletionTime time.Time `json:"estimatedCompletionTime"`
}

// AdminTransaction extends TransactionStatus with admin fields
type AdminTransaction struct {
	TransactionStatus
	RetryCount int    `json:"retryCount"`
	LastError  string `json:"lastError,omitempty"`
	InternalID string `json:"internalId"`
}

// AdminTransactionList represents paginated admin transaction list
type AdminTransactionList struct {
	Transactions []AdminTransaction `json:"transactions"`
	Pagination   Pagination         `json:"pagination"`
}

// FeeStructure represents bridge fee structure
type FeeStructure struct {
	EthereumToPolygon NetworkFee `json:"ethereumToPolygon"`
	PolygonToEthereum NetworkFee `json:"polygonToEthereum"`
}

type NetworkFee struct {
	Percentage string `json:"percentage"`
	Minimum    string `json:"minimum"`
	Maximum    string `json:"maximum"`
}

// TransferLimits represents transfer limits
type TransferLimits struct {
	Minimum    string `json:"minimum"`
	Maximum    string `json:"maximum"`
	DailyLimit string `json:"dailyLimit"`
}

// SupportedTokens represents supported tokens
type SupportedTokens struct {
	Ethereum []TokenInfo `json:"ethereum"`
	Polygon  []TokenInfo `json:"polygon"`
}

type TokenInfo struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
}

// TransferEstimate represents transfer estimate
type TransferEstimate struct {
	EstimatedTime int         `json:"estimatedTime"`
	BridgeFee     string      `json:"bridgeFee"`
	TotalAmount   string      `json:"totalAmount"`
	NetworkFees   NetworkFees `json:"networkFees"`
}

type NetworkFees struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
