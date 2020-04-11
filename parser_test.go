package gateway

import "testing"

func TestParseCreateTransactionResponse(t *testing.T) {
	testCases := []struct {
		Input  string
		Result CreateTransactionResult
	}{
		{
			Input: "TRANSACTION_ID: LUTI+KP7ikF7x0Ndg75lhIUbKuI=",
			Result: CreateTransactionResult{
				Result:   CommandResult{Raw: "TRANSACTION_ID: LUTI+KP7ikF7x0Ndg75lhIUbKuI=", IsSuccessful: true},
				Response: CreateTransactionResponse{TransactionID: "LUTI+KP7ikF7x0Ndg75lhIUbKuI="},
			},
		},
	}

	for _, tc := range testCases {
		if tc.Result != ParseCreateTransactionResponse(tc.Input) {
			t.Fatalf("unexpected result for input: %s. \nExpected: %v \nGot: %v", tc.Input, tc.Result, ParseCreateTransactionResponse(tc.Input))
		}
	}
}

func TestParseTransactionStatusResponse(t *testing.T) {
	testCases := []struct {
		Input  string
		Cmd    command
		Result TransactionStatusResult
	}{
		{
			Input: "RESULT: OK RESULT_CODE: 000 RRN: 728418142503 APPROVAL_CODE: 414576 CARD_NUMBER: 4***********4813",
			Cmd:   commandSMSTx,
			Result: TransactionStatusResult{
				Result: CommandResult{Raw: "RESULT: OK RESULT_CODE: 000 RRN: 728418142503 APPROVAL_CODE: 414576 CARD_NUMBER: 4***********4813", IsSuccessful: true},
				Response: TransactionStatusResponse{
					Result:       "OK",
					ResultCode:   "000",
					RRN:          "728418142503",
					ApprovalCode: "414576",
					CardNumber:   "4***********4813",
				},
			},
		},
		{
			Input: "RESULT: OK RESULT_CODE: 000 RRN: 728418142503 APPROVAL_CODE: 414576 CARD_NUMBER: 4***********4813",
			Cmd:   commandDMSTxCommit,
			Result: TransactionStatusResult{
				Result: CommandResult{Raw: "RESULT: OK RESULT_CODE: 000 RRN: 728418142503 APPROVAL_CODE: 414576 CARD_NUMBER: 4***********4813", IsSuccessful: true},
				Response: TransactionStatusResponse{
					Result:       "OK",
					ResultCode:   "000",
					RRN:          "728418142503",
					ApprovalCode: "414576",
					CardNumber:   "4***********4813",
				},
			},
		},
	}

	for _, tc := range testCases {
		if tc.Result != ParseTransactionStatusResponse(tc.Input, tc.Cmd) {
			t.Fatalf("unexpected result for input: %s - %s. \nExpected: %v \nGot: %v", tc.Input, tc.Cmd, tc.Result, ParseTransactionStatusResponse(tc.Input, tc.Cmd))
		}
	}
}

func TestParseCancelTransactionResponse(t *testing.T) {
	testCases := []struct {
		Input  string
		Cmd    command
		Result CancelTransactionResult
	}{
		{
			Input: "RESULT: OK RESULT_CODE: 400",
			Cmd:   commandTxReverse,
			Result: CancelTransactionResult{
				Result:   CommandResult{Raw: "RESULT: OK RESULT_CODE: 400", IsSuccessful: true},
				Response: CancelTransactionResponse{Result: "OK", ResultCode: "400"},
			},
		},
		{
			Input: "RESULT: OK RESULT_CODE: 000",
			Cmd:   commandTxRefund,
			Result: CancelTransactionResult{
				Result:   CommandResult{Raw: "RESULT: OK RESULT_CODE: 000", IsSuccessful: true},
				Response: CancelTransactionResponse{Result: "OK", ResultCode: "000"},
			},
		},
	}
	for _, tc := range testCases {
		if tc.Result != ParseCancelTransactionResponse(tc.Input, tc.Cmd) {
			t.Fatalf("unexpected result for input: %s - %s. \nExpected: %v \nGot: %v", tc.Input, tc.Cmd, tc.Result, ParseCancelTransactionResponse(tc.Input, tc.Cmd))
		}
	}
}
