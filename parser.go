package gateway

import (
	"strings"
)

// ParseCreateTransactionResponse parses provided response to format transaction create result
func ParseCreateTransactionResponse(res string) CreateTransactionResult {
	if len(res) == 0 {
		return CreateTransactionResult{}
	}
	ctr := CreateTransactionResult{
		CommandResult{Raw: res},
		CreateTransactionResponse{},
	}

	if strings.Index(res, responseParamTransactionID) == -1 || len(res) < len(responseParamTransactionID)+28 {
		return ctr
	}
	ctr.TransactionID = res[len(res)-28:]
	ctr.IsSuccessful = true
	return ctr
}

// ParseTransactionStatusResponse parses provided response to format transaction status result
func ParseTransactionStatusResponse(res string, cmd command) TransactionStatusResult {
	if len(res) == 0 {
		return TransactionStatusResult{}
	}
	tsr := TransactionStatusResult{
		CommandResult{Raw: res},
		TransactionStatusResponse{},
	}
	// try to parse result code
	resultCodeAt := strings.Index(res, responseParamResultCode)
	if resultCodeAt != -1 && len(res) >= resultCodeAt+len(responseParamResultCode)+5 {
		tsr.ResultCode = res[resultCodeAt+len(responseParamResultCode)+2 : resultCodeAt+len(responseParamResultCode)+5]
	}
	if tsr.ResultCode == string(successCodeByCommand(cmd)) {
		tsr.IsSuccessful = true
	}
	// try to parse result
	resultAt := strings.Index(res, responseParamResult+":")
	if resultAt != -1 && len(res) > resultAt+len(responseParamResult)+2+2 {
		if strings.ToUpper(res[resultAt+len(responseParamResult)+2:resultAt+len(responseParamResult)+4]) == "OK" {
			tsr.Result = "OK"
		} else {
			tsr.Result = "FAILED"
		}
	}
	// try to parse rrn number
	rrnAt := strings.Index(res, responseParamRRN)
	if rrnAt != -1 && len(res) > rrnAt+len(responseParamRRN)+2+12 {
		tsr.RRN = res[rrnAt+len(responseParamRRN)+2 : rrnAt+len(responseParamRRN)+14]
	}
	// try to parse approval code
	approvalCodeAt := strings.Index(res, responseParamApprovalCode)
	if approvalCodeAt != -1 && len(res) > approvalCodeAt+len(responseParamApprovalCode)+2 {
		for i := approvalCodeAt + len(responseParamApprovalCode) + 2; i < len(res); i++ {
			if isCharNumeric(string(res[i])) {
				tsr.ApprovalCode += string(res[i])
			}
			if len(tsr.ApprovalCode) == 6 {
				break
			}
		}
	}
	// try to parse masked card number
	cardNumberAt := strings.Index(res, responseParamCardNumber)
	if cardNumberAt != -1 && len(res) > cardNumberAt+len(responseParamCardNumber)+2 {
		for i := cardNumberAt + len(responseParamCardNumber) + 2; i < len(res); i++ {
			if string(res[i]) == "*" || isCharNumeric(string(res[i])) {
				tsr.CardNumber += string(res[i])
			}
		}
	}

	return tsr
}

// ParseCancelTransactionResponse parses provided response to format transaction cancellation result
func ParseCancelTransactionResponse(res string, cmd command) CancelTransactionResult {
	if len(res) == 0 {
		return CancelTransactionResult{}
	}
	ctr := CancelTransactionResult{
		CommandResult{
			Raw: res,
		},
		CancelTransactionResponse{},
	}
	// try to parse result code
	resultCodeAt := strings.Index(res, responseParamResultCode)
	if resultCodeAt != -1 && len(res) >= resultCodeAt+len(responseParamResultCode)+5 {
		ctr.ResultCode = res[resultCodeAt+len(responseParamResultCode)+2 : resultCodeAt+len(responseParamResultCode)+5]
	}
	if ctr.ResultCode == string(successCodeByCommand(cmd)) {
		ctr.IsSuccessful = true
	}
	// try to parse result
	resultAt := strings.Index(res, responseParamResult+":")
	if resultAt != -1 && len(res) > resultAt+len(responseParamResult)+2+2 {
		if strings.ToUpper(res[resultAt+len(responseParamResult)+2:resultAt+len(responseParamResult)+4]) == "OK" {
			ctr.Result = "OK"
		} else {
			ctr.Result = "FAILED"
		}
	}

	return ctr
}

func isCharNumeric(s string) bool {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}
