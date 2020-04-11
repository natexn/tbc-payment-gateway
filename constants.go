package gateway

// command is a type for describing types of commands for communicating with merchant handler endpoint
type command string

const (
	// commandTxStatus is for  getting transaction status
	commandTxStatus command = "c"
	// commandCloseDay is for closing merchant day
	commandCloseDay command = "b"
	// commandSMSTx is for creating a directly charging SMS transaction
	commandSMSTx command = "v"
	// commandDMSTxCreate is for initiating a DMS transaction (blocking the specified amount of money)
	commandDMSTxCreate command = "a"
	// commandDMSTxCommit is for completing a DMS transaction (confirming withdrawal of the specified amount of money within preauthorized transaction and unblocking the rest)
	commandDMSTxCommit command = "t"
	// commandTxReverse is for reversing specified transaction (should be called for a transaction in case merchant day has not been closed after its creation)
	commandTxReverse command = "r"
	// commandTxRefund is for refunding specified transaction (should be called in for a transaction in case it merchant day has been closed after its creation)
	commandTxRefund command = "k"
)

type messageType string

const (
	messageTypeSMS messageType = "SMS"
	messageTypeDMS messageType = "DMS"
)

const (
	requestParamCommand         string = "command"
	requestParamAmount          string = "amount"
	requestParamCurrency        string = "currency"
	requestParamClientIPAddress string = "client_ip_addr"
	requestParamTransID         string = "trans_id"
	requestParamDescription     string = "desc"
	requestParamMessageType     string = "msg_type"
	requestParamLanguage        string = "language"

	responseParamTransactionID string = "TRANSACTION_ID"
	responseParamResult        string = "RESULT"
	responseParamResultCode    string = "RESULT_CODE"
	responseParamRRN           string = "RRN"
	responseParamApprovalCode  string = "APPROVAL_CODE"
	responseParamCardNumber    string = "CARD_NUMBER"
	responseParamFLD075        string = "FLD_075"
	responseParamFLD076        string = "FLD_076"
	responseParamFLD087        string = "FLD_087"
	responseParamFLD088        string = "FLD_088"
)

// TemplateLanguage is a type for describing available template languages
type TemplateLanguage string

const (
	// TemplateLanguageEn is for English
	TemplateLanguageEn TemplateLanguage = "EN"
	// TemplateLanguageGe is for Georgian
	TemplateLanguageGe TemplateLanguage = "GE"
)

// Currency is a type for describing currencies codes correspondingly to ISO 4217 standard (TODO: the list is not complete)
type Currency int

const (
	// CurrencyGEL is for Georgian lari
	CurrencyGEL Currency = 981
	// CurrencyUSD is for United States dollar
	CurrencyUSD Currency = 840
	// CurrencyEUR is for Euro
	CurrencyEUR Currency = 978
	// CurrencyGBP is for Pound sterling
	CurrencyGBP Currency = 826
)

// successCode describes a list of codes which indicate different operations' result as successful
type successCode string

const (
	// successCodeGeneral is for indicating a generatl response indication command success
	successCodeGeneral successCode = "000"
	// successCodeReversal is for indicating a successful transaction reversal
	successCodeReversal successCode = "400"
	// successCodeDayClosed is for indicating a successful closing of the merchant day
	successCodeDayClosed successCode = "500"
)

func successCodeByCommand(cmd command) successCode {
	switch cmd {
	case commandCloseDay:
		return successCodeDayClosed
	case commandTxReverse:
		return successCodeReversal
	default:
		return successCodeGeneral
	}
}
