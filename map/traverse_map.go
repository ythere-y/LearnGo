package _map

import (
	"encoding/json"
	"fmt"
)

func mapTraverse() {
	{
		// find key by value

		find := "7126821491647283204"
		//find := "7124225739431444483"
		var keyList []string
		rawStr := "{\"auto_failover_msg\": [\"abcd\",\"7124225739431444483\"],\n" +
			"    \"save_merchant_rules_msg\":[\"7126821491647283204\"],\n" +
			"    \"one_more\":[\"ALL\"]" +
			"}"

		var testMap = make(map[string][]string)
		err := json.Unmarshal([]byte(rawStr), &testMap)
		if err != nil {
			panic(err)
		}
		for key, val := range testMap {
			for _, id := range val {
				if id == "ALL" {
					keyList = append(keyList, key)
					break
				}
				if id == find {
					keyList = append(keyList, key)
					break
				}
			}
		}
		fmt.Printf("get keyList = %+v\n", keyList)

	}
}
