package _map

import (
	"encoding/json"
	"fmt"
	"strings"
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
	test_get := judgeDisplayTag(tagRule, params)
	fmt.Printf("test_get = %+v\n", test_get)
}

func judgeDisplayTag(displayTagRule string, params map[string]interface{}) string {
	if len(displayTagRule) == 0 {
		return ""
	}
	var filterMap = make(map[string]string)
	err := json.Unmarshal([]byte(displayTagRule), &filterMap)
	if err != nil {
		return "error"
	}
	priorityKey := ElementDefaultPriority

	if pk := filterMap[consts.ElementPriority]; len(pk) != 0 {
		priorityKey = strings.Split(pk, consts.ElementSeparator)
	}

	var defaultKey string
	if k, ok := filterMap[consts.ElementDefault]; ok {
		defaultKey = k
		delete(filterMap, consts.ElementDefault)
	}

	for i, _ := range priorityKey { //遍历优先级列表
		key := priorityKey[i]
		value := filterMap[key]
		if value == "*" {
			return key
		}
		if len(value) <= 0 { //为空的话，表示不匹配，直接过
			continue
		}
		isMatch, err := yql.Match(value, params)
		if err != nil {
			continue
		}
		if isMatch {
			return key
		}
	}
	return defaultKey
}
