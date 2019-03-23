package gorecover

import "fmt"

func ConvErr(err interface{}) error {
	if err == nil {
		return nil
	}
	e, ok := err.(error)
	if !ok {
		e = fmt.Errorf("%+v", err)
	}
	return e
}
