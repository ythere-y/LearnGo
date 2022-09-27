package _map

import (
	"code.byted.org/hc_test/interview"
	"encoding/json"
	"fmt"

	"github.com/luci/go-render/render"
)

func Main() {
	testUnmarshalStringIntoMap()
	//IfMapIsReference()
	//testMapQuery()
	//syncMapUsage()
}

func testUnmarshalStringIntoMap() {
	{
		rawStr := "{" +
			"\"sdk\":\"true\"," +
			"\"hc\":\"false\"" +
			"}"
		var testMap = make(map[string]string)
		err := json.Unmarshal([]byte(rawStr), &testMap)
		if err != nil {
			println("unmarshal failed")
			return
		}
		var flag bool
		flag = false
		testName := "bad"
		for key, val := range testMap {
			if key == testName {
				if val == "true" {
					flag = true
				}
			}
		}
		fmt.Printf("[testName = %v], [result flag = %v]\n", testName, flag)
	}
	{
		rawStr := "{" +
			"\"sdk\":true," +
			"\"hc\":false" +
			"}"
		var testMap = make(map[string]bool)
		err := json.Unmarshal([]byte(rawStr), &testMap)
		if err != nil {
			println("unmarshal failed")
			return
		}
		{
			testName := "bad"
			var flag = false
			if val, ok := testMap[testName]; ok {
				flag = val
			}
			fmt.Printf("[testName = %v], [result flag = %v]\n", testName, flag)
		}
		{
			testName := "hc"
			var flag = false
			if val, ok := testMap[testName]; ok {
				flag = val
			}
			fmt.Printf("[testName = %v], [result flag = %v]\n", testName, flag)
		}
		{
			testName := "sdk"
			var flag = false
			if val, ok := testMap[testName]; ok {
				flag = val
			}
			fmt.Printf("[testName = %v], [result flag = %v]\n", testName, flag)
		}
	}
	return
}

func IfMapIsReference() {
	testMap := make(map[string]string)
	testMap = map[string]string{
		"payment_code": "payment_code_val",
		"iban":         "iban_val",
	}
	addVals := [2]string{"hello", "world"}

	fmt.Printf("befor the adjust, map = %+v\n", render.Render(testMap))
	testAdjustMapByfunc(testMap, addVals)
	fmt.Printf("after the adjust, map = %+v\n", render.Render(testMap))

	{
		// indicate that, can't adjust map's value directly by modify val in for
		fmt.Printf("\n[___ test adjust map in for loop ___]\n")
		for key, val := range testMap {
			if key == "payment_code" && val == "payment_code_val" {
				val = "deleted"
			}
		}
		fmt.Printf("after the adjust, map = %+v\n", render.Render(testMap))
	}

}

func testAdjustMapByfunc(testMap map[string]string, addVals [2]string) {
	interview.Title(testAdjustMapByfunc)
	testMap[addVals[0]] = addVals[0] + "_val"
	testMap[addVals[1]] = addVals[1] + "_val"
}
func testMapQuery() {
	var tMap = make(map[string]bool)
	tMap["hell"] = false
	tMap["one"] = true
	if ok := tMap["query"]; !ok {
		println("no query")
	} else {
		println("get query = %v", tMap["query"])
	}
}
