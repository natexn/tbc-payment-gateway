package gateway

import (
	"strconv"
	"strings"
)

// ParseCreateTransactionResponse parses provided response to format transaction create result
func ParseCreateTransactionResponse(res string) CreateTransactionResult {
	if len(res) == 0 {
		return CreateTransactionResult{}
	}
	ctr := CreateTransactionResult{
		Result:   CommandResult{Raw: res},
		Response: CreateTransactionResponse{},
	}
	if strings.Index(res, responseParamTransactionID) == -1 || len(res) < len(responseParamTransactionID)+28 {
		return ctr
	}
	ctr.Response.TransactionID = res[len(res)-28:]
	ctr.Result.IsSuccessful = true
	return ctr
}

// ParseTransactionStatusResponse parses provided response to format transaction status result
func ParseTransactionStatusResponse(res string, cmd command) TransactionStatusResult {
	if len(res) == 0 {
		return TransactionStatusResult{}
	}
	tsr := TransactionStatusResult{
		Result:   CommandResult{Raw: res},
		Response: TransactionStatusResponse{},
	}
	// try to parse result code
	resultCodeAt := strings.Index(res, responseParamResultCode)
	if resultCodeAt != -1 && len(res) >= resultCodeAt+len(responseParamResultCode)+5 {
		tsr.Response.ResultCode = res[resultCodeAt+len(responseParamResultCode)+2 : resultCodeAt+len(responseParamResultCode)+5]
	}
	if tsr.Response.ResultCode == string(successCodeByCommand(cmd)) {
		tsr.Result.IsSuccessful = true
	}
	// try to parse result
	resultAt := strings.Index(res, responseParamResult+":")
	if resultAt != -1 && len(res) > resultAt+len(responseParamResult)+2+2 {
		if strings.ToUpper(res[resultAt+len(responseParamResult)+2:resultAt+len(responseParamResult)+4]) == "OK" {
			tsr.Response.Result = "OK"
		} else {
			tsr.Response.Result = "FAILED"
		}
	}
	// try to parse rrn number
	rrnAt := strings.Index(res, responseParamRRN)
	if rrnAt != -1 && len(res) > rrnAt+len(responseParamRRN)+2+12 {
		tsr.Response.RRN = res[rrnAt+len(responseParamRRN)+2 : rrnAt+len(responseParamRRN)+14]
	}
	// try to parse approval code
	approvalCodeAt := strings.Index(res, responseParamApprovalCode)
	if approvalCodeAt != -1 && len(res) > approvalCodeAt+len(responseParamApprovalCode)+2 {
		for i := approvalCodeAt + len(responseParamApprovalCode) + 2; i < len(res); i++ {
			if isCharNumeric(string(res[i])) {
				tsr.Response.ApprovalCode += string(res[i])
			}
			if len(tsr.Response.ApprovalCode) == 6 {
				break
			}
		}
	}
	// try to parse masked card number
	cardNumberAt := strings.Index(res, responseParamCardNumber)
	if cardNumberAt != -1 && len(res) > cardNumberAt+len(responseParamCardNumber)+2 {
		for i := cardNumberAt + len(responseParamCardNumber) + 2; i < len(res); i++ {
			if string(res[i]) == "*" || isCharNumeric(string(res[i])) {
				tsr.Response.CardNumber += string(res[i])
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
		Result:   CommandResult{Raw: res},
		Response: CancelTransactionResponse{},
	}
	// try to parse result code
	resultCodeAt := strings.Index(res, responseParamResultCode)
	if resultCodeAt != -1 && len(res) >= resultCodeAt+len(responseParamResultCode)+5 {
		ctr.Response.ResultCode = res[resultCodeAt+len(responseParamResultCode)+2 : resultCodeAt+len(responseParamResultCode)+5]
	}
	if ctr.Response.ResultCode == string(successCodeByCommand(cmd)) {
		ctr.Result.IsSuccessful = true
	}
	// try to parse result
	resultAt := strings.Index(res, responseParamResult+":")
	if resultAt != -1 && len(res) > resultAt+len(responseParamResult)+2+2 {
		if strings.ToUpper(res[resultAt+len(responseParamResult)+2:resultAt+len(responseParamResult)+4]) == "OK" {
			ctr.Response.Result = "OK"
		} else {
			ctr.Response.Result = "FAILED"
		}
	}

	return ctr
}

// ParseCloseDayResponse parses provided response to format day close result
func ParseCloseDayResponse(res string) CloseDayResult {
	if len(res) == 0 {
		return CloseDayResult{}
	}
	cdr := CloseDayResult{
		Result:   CommandResult{Raw: res},
		Response: CloseDayResponse{},
	}
	// try to parse result code
	resultCodeAt := strings.Index(res, responseParamResultCode)
	if resultCodeAt != -1 && len(res) >= resultCodeAt+len(responseParamResultCode)+5 {
		cdr.Response.ResultCode = res[resultCodeAt+len(responseParamResultCode)+2 : resultCodeAt+len(responseParamResultCode)+5]
	}
	if cdr.Response.ResultCode == string(successCodeByCommand(commandCloseDay)) {
		cdr.Result.IsSuccessful = true
	}
	// try to parse result
	resultAt := strings.Index(res, responseParamResult+":")
	if resultAt != -1 && len(res) > resultAt+len(responseParamResult)+2+2 {
		if strings.ToUpper(res[resultAt+len(responseParamResult)+2:resultAt+len(responseParamResult)+4]) == "OK" {
			cdr.Response.Result = "OK"
		} else {
			cdr.Response.Result = "FAILED"
		}
	}
	// try to parse FLD075 field
	fld075At := strings.Index(res, responseParamFLD075)
	if fld075At != -1 && len(res) > fld075At+len(responseParamFLD075)+2 {
		var fld075NumberStr string
		for i := fld075At + len(responseParamFLD075) + 2; i < len(res); i++ {
			if !isCharNumeric(string(res[i])) {
				break
			}
			fld075NumberStr += string(res[i])
		}
		fld075Number, err := strconv.Atoi(fld075NumberStr)
		if err == nil {
			cdr.Response.FLD075 = fld075Number
		}
	}
	// try to parse FLD076 field
	fld076At := strings.Index(res, responseParamFLD076)
	if fld076At != -1 && len(res) > fld076At+len(responseParamFLD076)+2 {
		var fld076NumberStr string
		for i := fld076At + len(responseParamFLD076) + 2; i < len(res); i++ {
			if !isCharNumeric(string(res[i])) {
				break
			}
			fld076NumberStr += string(res[i])
		}
		fld076Number, err := strconv.Atoi(fld076NumberStr)
		if err == nil {
			cdr.Response.FLD076 = fld076Number
		}
	}
	// try to parse FLD087 field
	fld087At := strings.Index(res, responseParamFLD087)
	if fld087At != -1 && len(res) > fld087At+len(responseParamFLD087)+2 {
		var fld087NumberStr string
		for i := fld087At + len(responseParamFLD087) + 2; i < len(res); i++ {
			if !isCharNumeric(string(res[i])) {
				break
			}
			fld087NumberStr += string(res[i])
		}
		fld087Number, err := strconv.Atoi(fld087NumberStr)
		if err == nil {
			cdr.Response.FLD087 = fld087Number
		}
	}
	// try to parse FLD088 field
	fld088At := strings.Index(res, responseParamFLD088)
	if fld088At != -1 && len(res) > fld088At+len(responseParamFLD088)+2 {
		var fld088NumberStr string
		for i := fld088At + len(responseParamFLD088) + 2; i < len(res); i++ {
			if !isCharNumeric(string(res[i])) {
				break
			}
			fld088NumberStr += string(res[i])
		}
		fld088Number, err := strconv.Atoi(fld088NumberStr)
		if err == nil {
			cdr.Response.FLD088 = fld088Number
		}
	}

	return cdr
}

func isCharNumeric(s string) bool {
	switch s {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return true
	default:
		return false
	}
}
