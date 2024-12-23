package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MUstHandleError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
