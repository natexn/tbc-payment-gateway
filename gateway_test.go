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
			Result: "amount=500&client_ip_addr=127.0.0.1&command=v&currency=981&description=UFCTEST&language=EN&msg_type=SMS",
		},
	}
	for _, tc := range testCases {
		if tc.Result != buildQueryStr(tc.Input) {
			t.Fatalf("unexpected result for input: %v. \nExpected: %s \nGot: %s", tc.Input, tc.Result, buildQueryStr(tc.Input))
		}
	}
}
