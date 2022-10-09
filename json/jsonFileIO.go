package json

import (
	"code.byted.org/pipo/toolbox/errs"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type STUT struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func jsonFileIOTest() {
	var (
		ctx          = context.Background()
		get          = &STUT{}
		caseFile     = "./json/data/test-1.json"
		getErr       errs.Error
		testCaseName = "testB"
	)
	Perr := Parse(ctx, caseFile, testCaseName, get)
	if Perr != nil {
		fmt.Printf("Parse err : %v\n", Perr)
	}
	fmt.Printf("Parse success.\nget =>\n%+v\ngetErr => \n%+v\n", get, getErr)
}

func Parse(ctx context.Context, caseFile string, CaseName string, mockResp interface{}) errs.Error {
	content, err := os.ReadFile(caseFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(content, &m)
	for k, v := range m[CaseName].(map[string]interface{}) {
		switch k {
		case "mockResp":
			b, _ := json.Marshal(v)
			if err := json.Unmarshal(b, mockResp); err != nil {
				fmt.Printf("err: %v\n", err)
				return errs.ParamError
			}
		default:
			return errs.ParamError
		}
	}
	return nil
}
