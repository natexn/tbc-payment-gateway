package gateway

// CommandResult holds information about raw response and command completion success
type CommandResult struct {
	Raw          string
	IsSuccessful bool
}

// CreateTransactionResponse describes result fields of expected response to transaction creation requests
type CreateTransactionResponse struct {
	TransactionID string
}

// CreateTransactionResult describes result of transaction creation related functions
type CreateTransactionResult struct {
	CommandResult
	CreateTransactionResponse
}

// TransactionStatusResponse describes result fields of expected response to requests related to transaction status inquiry/update
type TransactionStatusResponse struct {
	Result       string
	ResultCode   string
	RRN          string
	ApprovalCode string
	CardNumber   string
}

// TransactionStatusResult describes result of transaction status inquiry/update functions
type TransactionStatusResult struct {
	CommandResult
	TransactionStatusResponse
}

// CancelTransactionResponse describes result fields of expected response to transaction cancellation related requests
type CancelTransactionResponse struct {
	Result     string
	ResultCode string
}

// CancelTransactionResult describes result of transaction cancellation related functions
type CancelTransactionResult struct {
	CommandResult
	CancelTransactionResponse
}
