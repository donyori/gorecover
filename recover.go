package gorecover

func Recover(f func()) (err error) {
	defer func() {
		err = ConvErr(recover())
	}()
	f()
	return nil
}
