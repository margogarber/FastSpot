package payments

import (
	"fmt"
	"time"
)

// PaymentProvider defines the interface for payment providers
type PaymentProvider interface {
	ProcessPayment(amount float64, method string) (*PaymentResult, error)
}

// PaymentResult represents the result of a payment transaction
type PaymentResult struct {
	TransactionID string
	Status        string // "success" | "pending" | "failed"
	Message       string
	ProcessedAt   time.Time
}

// StubProvider implements a stub payment provider for testing
type StubProvider struct{}

// NewStubProvider creates a new stub payment provider
func NewStubProvider() *StubProvider {
	return &StubProvider{}
}

// ProcessPayment simulates payment processing
func (p *StubProvider) ProcessPayment(amount float64, method string) (*PaymentResult, error) {
	// Simulate processing delay
	time.Sleep(100 * time.Millisecond)

	// Generate fake transaction ID
	transactionID := fmt.Sprintf("STUB-%d", time.Now().Unix())

	// Simulate different outcomes based on amount
	status := "success"
	message := "Payment processed successfully"

	// Simulate pending for large amounts
	if amount > 100 {
		status = "pending"
		message = "Payment is being processed"
	}

	// Simulate failure for amounts ending in .13 (unlucky number)
	if int(amount*100)%100 == 13 {
		status = "failed"
		message = "Payment declined"
	}

	return &PaymentResult{
		TransactionID: transactionID,
		Status:        status,
		Message:       message,
		ProcessedAt:   time.Now(),
	}, nil
}
