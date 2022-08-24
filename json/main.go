package json

import (
	"encoding/json"
	"fmt"
	"github.com/luci/go-render/render"
)

func Main() {
	testUnmarshalOmitempty()
}

type TagTest struct {
	Tag string `json:"tag"`
	//UsePPE int    `json:"usePPE,omitempty"`
	UsePPE int `json:"usePPE"`
}

func testUnmarshalOmitempty() {
	var tag = TagTest{
		Tag: "hc_boe",
	}
	var get, _ = json.Marshal(tag)
	fmt.Printf("get  >> \n%+v", string(get))

	var str = string(get)
	var re_get = TagTest{}
	_ = json.Unmarshal([]byte(str), &re_get)
	fmt.Printf("\nre get >>\n%+v", render.Render(re_get))
}
