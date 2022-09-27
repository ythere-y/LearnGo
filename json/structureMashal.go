package json

import (
	"encoding/json"
	"fmt"
)

type PtrConfig struct {
	Cfg   *config `json:"cfg"`
	Score int     `json:"score"`
	Flag  bool    `json:"flag"`
}

// notes
// 结构体使用marshal需要使用属性大写
// 而后面不一定要用json解释
type config struct {
	Name string `json:"name"`
	//Name string `json:"name"`
	//name string
	Value string `json:"value"`
	//Value string `json:"value"`
	//value string
}

func configMarshal() {

	lastVersionConfigs := make(map[string]config, 0)
	lastVersionConfigs["hc"] = config{
		Name:  "hello",
		Value: "world",
	}
	//lastVersionConfigs["hc"] = config{
	//	name:  "hello",
	//	value: "world",
	//}
	bytes, err := json.Marshal(lastVersionConfigs)
	if err != nil {
		return
	}
	Str := string(bytes)
	fmt.Printf("get str > = \n%+v\n", Str)
}
