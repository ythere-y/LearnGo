package _map

import (
	"fmt"
	"time"
)

func assembleRuleParams() map[string]interface{} {
	//write hard code for now, later can move to tcc
	//convert related param
	var result = make(map[string]interface{})
	result["amount_value"] = 10
	result["country_code"] = "US"
	result["currency"] = "USD"
	result["fx_token"] = ""
	result["merchant_id"] = "sdk_test"
	result["quote_amount_value"] = ""
	result["quote_currency"] = ""
	result["sdk_type"] = "go"
	result["sdk_version"] = nil
	result["show_billing_address"] = nil
	result["transaction_time"] = time.Now()
	result["transaction_trade_type"] = "merchant_instant_payment"
	result["user_id"] = "7099330234814678022"
	result["version"] = 5010

	return result
}

func ProblemInWork() {
	params := assembleRuleParams()
	tagRule := ""
	fmt.Printf("params = %+v, tagRule = %+v\n", params, tagRule)
}
