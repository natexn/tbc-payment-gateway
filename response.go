package gateway

// CommandResult holds information about raw response and command completion success
type CommandResult struct {
	Raw          string
	IsSuccessful bool
}

// CreateTransactionResponse describes fields of expected response to transaction creation requests
type CreateTransactionResponse struct {
	TransactionID string
}

// CreateTransactionResult describes result of transaction creation related functions
type CreateTransactionResult struct {
	Result   CommandResult
	Response CreateTransactionResponse
}

// TransactionStatusResponse describes fields of expected response to requests related to transaction status inquiry/update
type TransactionStatusResponse struct {
	Result       string
	ResultCode   string
	RRN          string
	ApprovalCode string
	CardNumber   string
}

// TransactionStatusResult describes result of transaction status inquiry/update functions
type TransactionStatusResult struct {
	Result   CommandResult
	Response TransactionStatusResponse
}

// CancelTransactionResponse describes fields of expected response to transaction cancellation related requests
type CancelTransactionResponse struct {
	Result     string
	ResultCode string
}

// CancelTransactionResult describes result of transaction cancellation related functions
type CancelTransactionResult struct {
	Result   CommandResult
	Response CancelTransactionResponse
}

// CloseDayResponse describes fields of expected response to day close request
type CloseDayResponse struct {
	Result     string
	ResultCode string
	// FLD075 - Credits, Reversal Number
	FLD075 int
	// FLD076 - Debits, Number
	FLD076 int
	// FLD087 - Credits, Reversal Amount
	FLD087 int
	// FLD087 - Debits, Amount.
	FLD088 int
}

// CloseDayResult describes result of day close request
type CloseDayResult struct {
	Result   CommandResult
	Response CloseDayResponse
}
