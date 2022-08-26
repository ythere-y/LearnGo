package json

import (
	"encoding/json"
	"fmt"
	"github.com/luci/go-render/render"
)

func Main() {
	testUnmarshalOmitempty()
}

type Student struct {
	Name  *string `json:"name"`
	Score int     `json:"score"`
}

type Order struct {
	OutTradeId *string `json:"outTradeId"`
}
type TagTest struct {
	Input Order `json:"Input"`
	//Student interface{}
	//Tag     string `json:"tag"`
	//UsePPE int    `json:"usePPE,omitempty"`
	//UsePPE int `json:"usePPE"`
}

func testUnmarshalOmitempty() {
	//var name = "Tom"
	//var tag = TagTest{
	//	Student: Student{
	//		Name:  &name,
	//		Score: 123,
	//	},
	//	Tag: "hc_boe",
	//}
	var get = "{\n  \"Input\": {\n    \"OutTradeId\": \"7036923020384699141\",\n    \"TradeId\":    \"ARMAAMCmObi9PZE1Dv6WIu3eUnMoD\",\n    \"MerchantId\": \"11202107MqrYt9\"\n  }\n}"
	//var get = "{     \"Student\": {\"name\":\"Tom\",\"score\":123},\"tag\":\"hc_boe\",\"usePPE\":0}\n"
	//var get = "{\n  \"input\": {\n    \"OutTradeId\": \"7036923020384699141\",\n    \"TradeId\":    \"ARMAAMCmObi9PZE1Dv6WIu3eUnMoD\",\n    \"MerchantId\": \"11202107MqrYt9\"\n  }\n}"
	//var get, _ = json.Marshal(tag)
	fmt.Printf("get  >> \n%v", string(get))

	var str = string(get)
	var re_get = TagTest{}
	_ = json.Unmarshal([]byte(str), &re_get)
	fmt.Printf("\nre get >>\n%+v", render.Render(re_get))
}
