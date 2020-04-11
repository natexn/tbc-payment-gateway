package gateway

import (
	"testing"
)

func TestBuildQueryStr(t *testing.T) {
	testCases := []struct {
		Input  map[string]interface{}
		Result string
	}{
		{
			Input:  map[string]interface{}{requestParamTransID: "B2R+qPonHFveOtgPMYs+J0Zb8kw="},
			Result: "trans_id=B2R%2BqPonHFveOtgPMYs%2BJ0Zb8kw%3D",
		},
		{
			Input: map[string]interface{}{
				requestParamCommand:         commandSMSTx,
				requestParamAmount:          500,
				requestParamCurrency:        CurrencyGEL,
				requestParamClientIPAddress: "127.0.0.1",
				requestParamLanguage:        TemplateLanguageEn,
				requestParamDescription:     "UFCTEST",
				requestParamMessageType:     messageTypeSMS,
			},
			Result: "command=v&amount=500&currency=981&client_ip_addr=127.0.0.1&language=EN&description=UFCTEST&msg_type=SMS",
		},
	}
	for _, tc := range testCases {
		if tc.Result != buildQueryStr(tc.Input) {
			t.Fatalf("unexpected result for input: %s. \nExpected: %v \nGot: %v", tc.Input, tc.Result, buildQueryStr(tc.Input))
		}
	}
}
