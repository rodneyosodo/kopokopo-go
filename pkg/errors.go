package kopokopo

import "errors"

var (
	// ErrEmptyToken indicates missing or invalid bearer token.
	ErrEmptyToken = errors.New("empty token")

	// ErrInvalidPaymentChannel indicates missing or invalid payment channel.
	ErrInvalidPaymentChannel = errors.New("invalid payment channel")

	// ErrMaxMetadataSize indicates maximum metadata length
	ErrMaxMetadataSize = errors.New("maximum metadata size is 5")

	// ErrEmptyID indicates an empty reference or ID
	ErrEmptyID = errors.New("empty id")

	// ErrInvalidSettlementMethod indicates invalid settlement method
	ErrInvalidSettlementMethod = errors.New("invalid settlement method")
)
