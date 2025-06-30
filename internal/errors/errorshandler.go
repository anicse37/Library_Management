package errors_package

var err error

func GetError() error {

	return err
}
func SetError(seterr error) {
	err = seterr
}
