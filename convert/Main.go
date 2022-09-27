package convert

import (
	"code.byted.org/gopkg/thrift"
	"fmt"
	"strconv"
)

func Main() {
	//var ReleaseId = "7143488578604450054"
	var ReleaseId *string
	ReleaseId = thrift.StringPtr("123")
	ReleaseId = nil
	var releaseIdParam int64
	releaseIdParam, parseErr := strconv.ParseInt(*ReleaseId, 10, 64) // todo fix here
	if parseErr != nil {
		return
	}
	fmt.Printf("get release id = %v\n", releaseIdParam)
}
