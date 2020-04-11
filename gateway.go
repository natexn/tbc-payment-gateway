package gateway

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/andelf/go-curl"
)

// Gateway is a type for describing tbc payment gateway properies
type Gateway struct {
	submitURL   string
	certAbsPath string
	certPass    string
}

// New returns instance of the tbc payment gateway
func New(submitURL, certAbsPath, certPass string) Gateway {
	return Gateway{
		submitURL:   submitURL,
		certAbsPath: certAbsPath,
		certPass:    certPass,
	}
}

// CreateSMSTransaction creates SMS transaction and returns transaction id if successful or corresponding error
func (g *Gateway) CreateSMSTransaction(amount int, currency Currency, clientIPAddress, description string, language TemplateLanguage) (*CreateTransactionResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandSMSTx
	paramsMap[requestParamClientIPAddress] = clientIPAddress
	paramsMap[requestParamAmount] = amount
	paramsMap[requestParamCurrency] = currency
	paramsMap[requestParamDescription] = description
	paramsMap[requestParamLanguage] = language
	paramsMap[requestParamMessageType] = messageTypeSMS
	// get query string
	queryStr := buildQueryStr(paramsMap)
	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	ctResult := ParseCreateTransactionResponse(*res)
	return &ctResult, nil
}

// CreateDMSTransaction creates DMS transaction and returns transaction id if successful or corresponding error
func (g *Gateway) CreateDMSTransaction(amount int, currency Currency, clientIPAddress, description string, language TemplateLanguage) (*CreateTransactionResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandDMSTxCreate
	paramsMap[requestParamClientIPAddress] = clientIPAddress
	paramsMap[requestParamAmount] = amount
	paramsMap[requestParamCurrency] = currency
	paramsMap[requestParamDescription] = description
	paramsMap[requestParamLanguage] = language
	paramsMap[requestParamMessageType] = messageTypeDMS
	// get query string
	queryStr := buildQueryStr(paramsMap)

	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	ctResult := ParseCreateTransactionResponse(*res)
	return &ctResult, nil
}

// CommitDMSTransaction commits pre-authorized transaction
func (g *Gateway) CommitDMSTransaction(amount int, currency Currency, clientIPAddress, description string, language TemplateLanguage) (*TransactionStatusResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandDMSTxCommit
	paramsMap[requestParamClientIPAddress] = clientIPAddress
	paramsMap[requestParamAmount] = amount
	paramsMap[requestParamCurrency] = currency
	paramsMap[requestParamDescription] = description
	paramsMap[requestParamLanguage] = language
	paramsMap[requestParamMessageType] = messageTypeDMS
	// get query string
	queryStr := buildQueryStr(paramsMap)

	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	tsr := ParseTransactionStatusResponse(*res, commandDMSTxCommit)
	return &tsr, nil
}

// GetTransactionStatus returns transaction status based on transaction id passed
func (g *Gateway) GetTransactionStatus(transactionID string) (*TransactionStatusResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandTxStatus
	paramsMap[requestParamTransID] = transactionID

	// get query string
	queryStr := buildQueryStr(paramsMap)
	// fmt.Println(queryStr)
	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	tsr := ParseTransactionStatusResponse(*res, commandTxStatus)
	return &tsr, nil
}

// ReverseTransaction reverses transaction if it was committed before day was closed for the last time
func (g *Gateway) ReverseTransaction(amount int, transactionID string) (*CancelTransactionResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandTxReverse
	paramsMap[requestParamTransID] = transactionID
	if amount != -1 {
		paramsMap[requestParamAmount] = amount
	}
	// get query string
	queryStr := buildQueryStr(paramsMap)

	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	ctr := ParseCancelTransactionResponse(*res, commandTxReverse)
	return &ctr, nil
}

// RefundTransaction refunds transaction if it was committed before the last date of close day
func (g *Gateway) RefundTransaction(amount int, transactionID string) (*CancelTransactionResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandTxRefund
	paramsMap[requestParamTransID] = transactionID
	if amount != -1 {
		paramsMap[requestParamAmount] = amount
	}
	// get query string
	queryStr := buildQueryStr(paramsMap)

	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	ctr := ParseCancelTransactionResponse(*res, commandTxRefund)
	return &ctr, nil
}

// CloseDay returns transaction status based on transaction id passed
func (g *Gateway) CloseDay() (*CloseDayResult, error) {
	// build params query
	paramsMap := make(map[string]interface{})
	paramsMap[requestParamCommand] = commandCloseDay

	// get query string
	queryStr := buildQueryStr(paramsMap)
	// execute call
	res, err := g.buildAndExecCurl(queryStr)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("empty response")
	}
	cdr := ParseCloseDayResponse(*res)
	return &cdr, nil
}

func (g *Gateway) buildAndExecCurl(queryStr string) (*string, error) {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_POSTFIELDS, queryStr)
	easy.Setopt(curl.OPT_VERBOSE, false)
	easy.Setopt(curl.OPT_SSL_VERIFYPEER, 1)
	easy.Setopt(curl.OPT_SSL_VERIFYHOST, 2)
	easy.Setopt(curl.OPT_TIMEOUT, 120)
	easy.Setopt(curl.OPT_CAINFO, g.certAbsPath) // because of Self-Signed certificate at payment server.
	easy.Setopt(curl.OPT_SSLCERT, g.certAbsPath)
	easy.Setopt(curl.OPT_SSLKEY, g.certAbsPath)
	easy.Setopt(curl.OPT_SSLKEYPASSWD, g.certPass)
	easy.Setopt(curl.OPT_URL, g.submitURL)
	easy.Setopt(curl.OPT_WRITEFUNCTION, writeDataToString)
	var result *string
	easy.Setopt(curl.OPT_WRITEDATA, result)

	if err := easy.Perform(); err != nil {
		return nil, err
	}
	return result, nil
}

func writeDataToString(ptr []byte, outputTo interface{}) bool {
	str, ok := outputTo.(*string)
	if ok {
		*str = string(ptr)
		return true
	}
	return false
}

func buildQueryStr(params map[string]interface{}) string {
	queryStr := ""
	for key, val := range params {
		var paramValue string
		switch val.(type) {
		case string:
			paramValue = url.QueryEscape(val.(string))
		case int:
			paramValue = strconv.Itoa(val.(int))
		default:
			continue
		}

		if len(queryStr) == 0 {
			queryStr = key + "=" + paramValue
		} else {
			queryStr += "&" + key + "=" + paramValue
		}
	}
	return queryStr
}
