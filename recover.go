package gorecover

import (
	"runtime/debug"
	"strings"
)

func Recover(f func()) (err error) {
	defer func() {
		err = ConvErr(recover())
	}()
	f()
	return nil
}

func RecoverWithStackTrace(f func()) (err error, stackTrace string) {
	defer func() {
		if e := recover(); e != nil {
			s := string(debug.Stack())
			idx := strings.Index(s, "\npanic(")
			if idx > 0 {
				stackTrace = s[idx+1:]
			}
			err = ConvErr(e)
		}
	}()
	f()
	return nil, ""
}
