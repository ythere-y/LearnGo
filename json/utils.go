package json

import (
	"encoding/json"
	"fmt"
)

func ConvertToJson(target interface{}) string {

	bytes, err := json.Marshal(target)
	if err != nil {
		return "marshal err"
	}
	Str := string(bytes)
	return Str
}

func UtilsTest() {
	{
		fmt.Printf("test1 = %v\n", ConvertToJson("hello"))
	}
	{
		testConfig := &config{
			Name:  "jack",
			Value: "build",
		}
		fmt.Printf("test2 = %v\n", ConvertToJson(testConfig))
	}
	{
		testPtrConfig := &PtrConfig{
			Cfg: &config{
				Name:  "queen",
				Value: "assemble",
			},
			Score: 101,
			Flag:  false,
		}
		fmt.Printf("test3 = %v\n", ConvertToJson(testPtrConfig))
	}
}
