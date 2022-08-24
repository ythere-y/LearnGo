package marshal

import (
	"encoding/json"
	"fmt"
)

type Combine struct {
	Type   string
	Region string
}

type Msg struct {
	Name        string            `json:"name,omitempty"`
	Value       int               `json:"value,omitempty"`
	CombineInfo Combine           `json:"combineInfo,omitempty"`
	ScoreList   []int             `json:"scoreList"`
	MapSample   map[string]string `json:"mapSample"`
}

func Main() {
	msg := Msg{
		//Name:  "nicky",
		//Value: 1234,
		CombineInfo: struct {
			Type   string
			Region string
		}{Type: "hello", Region: "build"},
		//ScoreList: []int{1, 2, 3, 4},
		//MapSample: map[string]string{"go": "gogo", "jojo": "adventure"},
	}

	richText, _ := json.Marshal(msg)
	fmt.Println(string(richText))
	//for _, line := range richText {
	//	fmt.Println(line)
	//}
}
